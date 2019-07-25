import asyncio
import os
import signal
from nats.aio.client import Client as NATS
import datastream_pb2


natsURL = "nats://nats:4222"
topic = "datapipeline-demo"

def run(loop):
    nc = NATS()
    @asyncio.coroutine
    def closed_cb():
        print("Connection to NATS is closed.")
        yield from asyncio.sleep(0.1, loop=loop)
        loop.stop()

    options = {
        "servers": [natsURL],
        "io_loop": loop,
        "closed_cb": closed_cb
    }

    yield from nc.connect(**options)
    print("Connected to NATS at {}...".format(nc.connected_url.netloc))

    @asyncio.coroutine
    def subscribe_handler(msg):
        datastreamMsg = datastream_pb2.DataStreamMessage()
        datastreamMsg.ParseFromString(msg.data)
        print("Received a message on {data}".format(data=datastreamMsg.payload))

    yield from nc.subscribe(topic, cb=subscribe_handler)
    print("Subscribed to topic: {}".format(topic))

    def signal_handler():
        if nc.is_closed:
            return
        print("Disconnecting...")
        loop.create_task(nc.close())

    for sig in ('SIGINT', 'SIGTERM'):
        loop.add_signal_handler(getattr(signal, sig), signal_handler)


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(run(loop))
    try:
        loop.run_forever()
    finally:
        loop.close()