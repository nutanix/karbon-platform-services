import os
import re
import json
from collections import Counter
from confluent_kafka import Consumer, TopicPartition
import logging
import numpy as np
import pandas as pd
import matplotlib.cm as cm
import matplotlib.pyplot as plt
from matplotlib import rcParams
from wordcloud import WordCloud, STOPWORDS

kafkaBroker = os.getenv('KAFKA_BROKER')
if kafkaBroker is None:
  kafkaBroker = "localhost:9092"

mountPath = os.getenv('MOUNT_PATH')
if mountPath is None:
  mountPath = "/www"

kafkaTopic = os.getenv('KAFKA_TOPIC')
if kafkaTopic is None:
  kafkaTopic = 'wordpress_db.wordpress_db.wp_comments'

msgCnt = 0
wordCounter = Counter()
wordText = ''

kConsumer = Consumer({
    'bootstrap.servers': kafkaBroker,
    'enable.auto.commit': 'False',
    'group.id': 'groupid1',
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

  comment = msgJson['payload']['after']['comment_content']
  logging.info("got comment from kafkaTopic: %s", comment)
  return comment

def wordCloud():
  logging.info("Generating wordCloud from text, cnt = %d", msgCnt)
  wordcloud = WordCloud(stopwords=STOPWORDS, width=800, height=400, background_color="white", max_words=1000).generate(wordText)
  logging.debug("wc freq %s: ", wordcloud.words_)
  wordcloud.to_file(mountPath + "/wordcloud.png")

def wordCloud2():
  logging.info("Generating wordCloud from freq, cnt = %d", msgCnt)
  wordcloud = WordCloud(stopwords=STOPWORDS, width=800, height=400, background_color="white", max_words=1000).generate_from_frequencies(wordCounter)
  wordcloud.to_file(mountPath + "/wordcloud2.png")

def createHtml():
  htmlStr = """<html>
<head>
<script>
function updateImage() {
  document.getElementById("img").src = "wordcloud.png?ts=" + encodeURIComponent(new Date().toString());
  setTimeout(updateImage, 10000);
};
</script>
<style>
.center {
  display: block;
  margin-top: 10em;
  margin-left: auto;
  margin-right: auto;
  width: 50%;
}
</style>
</head>

<body onload="updateImage();">
<img id="img" src="wordcloud.png" class="center"/>
</body>
</html>
"""
  f = open(mountPath + '/index.html', 'w')
  f.write(htmlStr)

def main():
  global msgCnt
  global wordCounter
  global wordText

  logging.basicConfig(format='%(asctime)s %(levelname)s:%(message)s', level=logging.DEBUG)

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

      msgCnt += 1

      # Update the global wordText
      wordText = wordText + ' ' + msg.lower().strip('\n')

      # generate wordCloud from text
      logging.info("wordText: %s", wordText)
      wordCloud()

      #method 2: generate freq here and pass it to wordcloud
      words = re.findall(r'\w+', msg.lower())
      wordCount = [word for word in words if word not in STOPWORDS]

      c1 = Counter(wordCount)
      logging.debug("c1: %s", c1)

      # Update the global wordCounter
      wordCounter.update(c1)

      # generate wordCloud from frequencies
      logging.info("wordCounter: %s", wordCounter.most_common(20))
      wordCloud2()
  finally:
    logging.debug("Closing consumer")
    kConsumer.close()

if __name__ == "__main__":
  main()
