import os
import io
import cv2
import sys
import time
import asyncio
import logging
import xi_iot_pb2
from typing import Callable
import numpy as np
import tensorflow as tf
from detect import FaceRecognizer
from PIL import Image
from nats.aio.client import Client as NATS
from nats.aio.errors import ErrConnectionClosed, ErrTimeout, ErrNoServers


logging.getLogger().setLevel(logging.DEBUG)

def _env_var(name: str, desc: str):
	val = os.environ.get(name)
	if val is None:
		raise Exception('{} is not provided via environment variable {}'.format(desc, name))
	return val

class VideoPipeline:
	
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


class FrameHandler:

	def __init__(self, fr, infer_freq):
		self.fr = fr
		self.infer_freq = infer_freq
		self.num_frames = 0
		self.boxes = np.array([])
		self.labels = []
	
	def __call__(self, img):
		if self.num_frames % self.infer_freq == 0:
			# Running inference every N frames 
			self.boxes, self.labels = self.fr(img)
		self.num_frames += 1
		return self.fr.label_faces(img, self.boxes, self.labels)


def process(fh):
	def fnc(payload: bytes) -> bytes:
		image = Image.open(io.BytesIO(payload)).convert('RGB')
		image_np = np.array(image)
		frame = fh(image_np)
		frame = Image.fromarray(frame)
		buffer = io.BytesIO()
		frame.save(buffer, format="JPEG")
		img_str = buffer.getvalue()
		return img_str
	return fnc

# Starting NATS loop 
sess = tf.Session()
fr = FaceRecognizer(sess)
fh = FrameHandler(fr, 10)
fnc = process(fh)
loop = asyncio.get_event_loop()
try:
	loop.run_until_complete(VideoPipeline(fnc).run(loop))
	loop.run_forever()
finally:
	loop.close()