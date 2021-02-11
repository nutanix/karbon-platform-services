import os
import re
import json
from collections import Counter
from confluent_kafka import Consumer, TopicPartition
import logging
import numpy as np
import pandas as pd
import matplotlib.cm as cm
import mpld3
import matplotlib.pyplot as plt
from matplotlib import rcParams

kafkaBroker = os.getenv('KAFKA_BROKER')
if kafkaBroker is None:
  kafkaBroker = "localhost:9092"

mountPath = os.getenv('MOUNT_PATH')
if mountPath is None:
  mountPath = "/www"

kafkaTopic = os.getenv('KAFKA_TOPIC')
if kafkaTopic is None:
  kafkaTopic = 'wordpress_db.wordpress_db.wp_woocommerce_order_items'

versionString = os.getenv('VERSION_STRING')
if versionString is None:
  versionString = "Version 0"

color = os.getenv('COLOR')
if color is None:
  color = "b"

debug = os.getenv('DEBUG')
if debug is None:
  debug = False
else:
  debug = True 

group = os.getenv('GROUP')
if group is None:
  group = "groupid1"

productDict = {}

kConsumer = Consumer({
    'bootstrap.servers': kafkaBroker,
    'enable.auto.commit': 'False',
    'group.id': group,
    'auto.offset.reset': 'earliest'
})

kConsumer.subscribe([kafkaTopic])

def readMsg():
  logging.debug("readMsg from kafkaTopic: %s", kafkaTopic)
  msg = kConsumer.poll(5.0)

  if msg is None:
    logging.debug('Received message: None')
    return None;
  if msg.error():
    logging.warning("Consumer error: {}".format(msg.error()))
    return None

  logging.debug('Received message: {}'.format(msg.value().decode('utf-8')))
  msgJson= json.loads(msg.value())

  logging.debug("got msg from kafkaTopic: %s", msgJson)
  return msgJson

def processOrder(msgJson):
  if msgJson is None:
    return None
  order = msgJson['payload']['after']
  logging.info("got order payload from kafkaTopic: %s", order)
  if order['order_item_type'] != 'line_item':
    return None
  logging.info("received order for item: %s", order['order_item_name'])
  return order['order_item_name']

def createHtml():
  htmlStr = f"""<html>
<head>
<script>
function updateImage() {{
  document.getElementById("img").src = "/recommendation-service/recommendation-service.png?ts=" + encodeURIComponent(new Date().toString());
  setTimeout(updateImage, 10000);
}};
</script>
<style>
.center {{
  display: block;
  margin-left: auto;
  margin-right: auto;
}}
* {{
  font-family: Arial, Helvetica, sans-serif
}}
</style>
</head>

<body onload="updateImage();">
<h1><center>{versionString}</center></h1>
<h2><center>Recommendation based on most ordered product</center></h2>
<img id="img" src="/recommendation-service/recommendation-service.png" class="center"/>
</body>
</html>
  """

  f = open(mountPath + '/index.html', 'w')
  f.write(htmlStr)

def createGraph():
  most_common = dict(Counter(productDict).most_common(5))

  plt.barh(list(most_common.keys()), most_common.values(), color=color)
  plt.yticks(rotation=20)
  plt.tight_layout()
  plt.tick_params(top='off', bottom='off', left='off', right='off', labelleft='off', labelbottom='on')
  plt.savefig(mountPath + '/recommendation-service.png')


def createGraph2():

  rcParams['font.family'] = 'sans-serif'

  most_common = dict(Counter(productDict).most_common(5))
  #plt.rcdefaults()
  fig, ax = plt.subplots()
  y_pos = np.arange(len(most_common))

  ax.barh(y_pos, most_common.values(), align='center', color=color)
  ax.set_yticks(y_pos)
  ax.set_yticklabels(list(most_common.keys()))
  ax.set_xlabel('Number of units sold')
  ax.set_title('Recommendation Service ' + versionString)
  plt.tight_layout()
  mpld3.save_html(fig, mountPath + '/graph.html')


def main():
  global productDict

  logging.info("Kafka Broker: %s", kafkaBroker)
  logging.info("Kafka Topic: %s", kafkaTopic)
  logging.info("Mount Path: %s", mountPath)

  createHtml()

  try:
    while True:
      # Read from kafkaTopic
      msg = readMsg()
      if msg is None:
        continue

      order = processOrder(msg)
      if order is None:
        continue
      if order in productDict:
        productDict[order] += 1
      else:
        productDict[order] = 1

      logging.info("product Dict: %s", productDict)
      createGraph2()
      createGraph()

  finally:
    logging.debug("Closing consumer")
    kConsumer.close()

if __name__ == "__main__":
  if debug:
      logging.basicConfig(format='%(asctime)s %(levelname)s:%(message)s', level=logging.DEBUG)
  else:
      logging.basicConfig(format='%(asctime)s %(levelname)s:%(message)s', level=logging.INFO)
  main()
