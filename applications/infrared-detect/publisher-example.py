import asyncio
import os
import time
import xi_iot_pb2
from nats.aio.client import Client as NATS
from nats.aio.errors import ErrConnectionClosed, ErrTimeout, ErrNoServers

# NOTE: This file is purely for diriving load through the nats server/broker passed in demo-app.py


dst_nats_topic = os.environ.get('DST_NATS_TOPIC')
if dst_nats_topic is None:
    print("please set DST_NATS_TOPIC in environment variables")

nats_broker_url = os.environ.get('NATS_BROKER_URL')
if nats_broker_url  is None:
    print("please set NATS_BROKER_URL in environment variables")

async def publish(loop):
    nc = NATS()
    await nc.connect(nats_broker_url, loop=loop)
    i = 0
    base_msg = "this is jpeg {i}"
    while True:
        _msg = xi_iot_pb2.DataStreamMessage()
        _msg.payload = str.encode(base_msg.format(i=i))
        await nc.publish(dst_nats_topic, _msg.SerializeToString())
        await asyncio.sleep(2)
        i += 1

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(publish(loop))
    loop.run_forever()
    loop.close()
