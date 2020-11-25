# Copyright (c) 2018 Nutanix, Inc.
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

from detect import FaceRecognizer

import xi_iot_pb2

from nats.aio.client import Client as NATS
from nats.aio.errors import ErrConnectionClosed, ErrTimeout, ErrNoServers
import cv2
import numpy as np
import tensorflow as tf

import asyncio
import logging
import os
import sys
import time
from typing import Callable

logging.getLogger().setLevel(logging.DEBUG)

def _env_var(name: str, desc: str):
    val = os.environ.get(name)
    if val is None:
        raise Exception('{} is not provided via environment variable {}'.format(desc, name))
    return val

class Transformation:
    def __init__(self, func: Callable[[bytes], bytes]):
        self.nc = None
        self.func = func
        self.msg_idx = 0

        self.nats_endpoint = _env_var('NATS_ENDPOINT', 'NATS broker endpoint')
        self.src_topic = _env_var('NATS_SRC_TOPIC', 'source NATS topic')
        self.dst_topic = _env_var('NATS_DST_TOPIC', 'destination NATS topic')
        logging.info("[{}]: {} -> {}".format(self.nats_endpoint, self.src_topic, self.dst_topic))

    async def run(self, loop):
        self.nc = NATS()
        await self.nc.connect(loop=loop, servers=[str(self.nats_endpoint)])
        logging.info('connected to NATS')
        await self.nc.subscribe(str(self.src_topic), cb=self._message_handler)

    async def _message_handler(self, msg):
        ds_msg = xi_iot_pb2.DataStreamMessage()
        ds_msg.ParseFromString(msg.data)
        try:
            ds_msg.payload = self.func(ds_msg.payload)  # TODO error handling
        except Exception as e:
            logging.error("failed to process message #{}: {}".format(self.msg_idx, e))
        if self.msg_idx % 100 == 0:
            logging.info("processed message #{}".format(self.msg_idx))
        try:
            await self.nc.publish(str(self.dst_topic), ds_msg.SerializeToString())
        except Exception as e:
            logging.error("failed to publish message #{}: {}".format(self.msg_idx, e))
        self.msg_idx += 1

    def __del__(self):
        if self.nc:
            self.nc.drain()

def jpegify(f):
    def g(payload: bytes) -> bytes:
        array = np.frombuffer(payload, np.uint8)
        image = cv2.imdecode(array, cv2.IMREAD_COLOR)
        image = f(image)
        _, buf = cv2.imencode('.jpeg', image)
        return buf.tobytes()
    return g

class inferEveryNthFrame:
    def __init__(self, fr, n, **kwargs):
        self.fr = fr
        self.frame_idx = 0
        self.n = n
        self.bboxes = None
        self.kwargs = kwargs

    def __call__(self, img):
        if self.frame_idx % self.n == 0:
            frame = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)
            t = time.time()
            self.bboxes = self.fr.bboxes(frame, **self.kwargs)
            delta = time.time() - t
            logging.info("{} faces detected in {}s".format(len(self.bboxes), delta))
        self.frame_idx += 1
        return self.fr.draw_overlay(img, self.bboxes)

model_fn = sys.argv[1]

frequency = int(os.environ.get("INFER_EVERY", 10))
bb_threshold = int(os.environ.get("BB_THRESHOLD", 80))
min_cos = float(os.environ.get("MIN_COSINE", 0.5))

sess = tf.Session()
fr = FaceRecognizer(sess, model_fn, min_cos=min_cos)
f = inferEveryNthFrame(fr, frequency, bb_threshold=bb_threshold)
f = jpegify(f)

loop = asyncio.get_event_loop()
try:
    loop.run_until_complete(Transformation(f).run(loop))
    loop.run_forever()
finally:
    loop.close()
