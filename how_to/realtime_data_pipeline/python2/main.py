#!/usr/bin/env python

import tornado.ioloop
import tornado.options
import logging
import json
import os
from tornado.options import define, options
import datastream_pb2
import sys
from nats.io import Client as NATS

logging.basicConfig(stream=sys.stdout, level=logging.DEBUG)

natsIPAddress = "nats"
natsPort = 4222
natsSubTopic = "datapipeline-demo"

def on_message(message=None):
    logging.info("Received message from NATS")
    datastreamMsg = datastream_pb2.DataStreamMessage()
    datastreamMsg.ParseFromString(str(message.data))
    payload = json.loads(datastreamMsg.payload)
    logging.info("nats got payload: {}".format(json.dumps(payload)))

@tornado.gen.coroutine
def main():
    logging.info("Connecting to NATS")
    server = "nats://%s:%d" % (natsIPAddress, natsPort)
    servers = []
    servers.append(server)
    opts = { "verbose": True, "servers": servers }
    nc = NATS()
    yield nc.connect(**opts)
    yield nc.subscribe(natsSubTopic, "", on_message)

if __name__ == "__main__":
    main()
    tornado.options.parse_command_line()
    tornado.ioloop.IOLoop.current().start()
