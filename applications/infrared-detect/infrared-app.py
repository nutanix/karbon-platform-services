import numpy as np
import asyncio
import os
import xi_iot_pb2
import cv2
import io
import boto3
import time
import psycopg2
from nats.aio.client import Client as NATS
from nats.aio.errors import ErrConnectionClosed, ErrTimeout, ErrNoServers
from PIL import Image
from configparser import ConfigParser

nc = None
nats_broker_url,src_nats_topic,dst_nats_topic = "","",""
msgs_received, msgs_sent = 0, 0
net = None
boxes = []
confidences = []
params = None
# classIDs = []

def get_nats_meta():
    global nats_broker_url,src_nats_topic,dst_nats_topic
    nats_broker_url = os.environ.get('NATS_ENDPOINT')
    if nats_broker_url is None:
        print('nats broker not provided in environment var NATS_ENDPOINT')
        exit(1)

    src_nats_topic = os.environ.get('NATS_SRC_TOPIC')
    if src_nats_topic is None:
        print('src nats topic not provided in environment var NATS_SRC_TOPIC')
        exit(1)
    dst_nats_topic = os.environ.get('NATS_DST_TOPIC')

    if dst_nats_topic is None:
        print('dst nats broker not provided in environment var NATS_DST_TOPIC')
        exit(1)
    return nats_broker_url, src_nats_topic, dst_nats_topic

def config(filename='database.ini', section='postgresql'):
    # create a parser
    parser = ConfigParser()
    # read config file
    parser.read(filename)
 
    # get section, default to postgresql
    db = {}
    if parser.has_section(section):
        params = parser.items(section)
        for param in params:
            db[param[0]] = param[1]
    else:
        raise Exception('Section {0} not found in the {1} file'.format(section, filename))
 
    return db

def load_net():
    # weights and configs for YOLO neural network
    weightsPath = 'yolov3.weights'
    configPath = 'yolov3.cfg'

    # load our YOLO object detector trained on COCO dataset (80 classes)
    print("[INFO] loading YOLO from disk...")
    net = cv2.dnn.readNetFromDarknet(configPath,weightsPath)
    return net

def detect(image):
    global net,boxes,confidences
    
    # load our input image and grab its spatial dimensions
    (H, W) = image.shape[:2]

    # determine only the *output* layer names that we need from YOLO
    ln = net.getLayerNames()
    ln = [ln[i[0] - 1] for i in net.getUnconnectedOutLayers()]

    # construct a blob from the input image and then perform a forward
    # pass of the YOLO object detector, giving us our bounding boxes and
    # associated probabilities
    blob = cv2.dnn.blobFromImage(image, 1 / 255.0, (416, 416),
        swapRB=True, crop=False)
    net.setInput(blob)
    layerOutputs = net.forward(ln)

    boxes.clear()
    confidences.clear()

    detections = 0
    # loop over each of the layer outputs
    for output in layerOutputs:
        # loop over each of the detections
        for detection in output:
            # extract the class ID and confidence (i.e., probability) of
            # the current object detection
            scores = detection[5:]
            classID = np.argmax(scores)
            confidence = scores[classID]

            # filter out weak predictions by ensuring the detected
            # probability is greater than the minimum probability
            if confidence > 0.5:
                # scale the bounding box coordinates back relative to the
                # size of the image, keeping in mind that YOLO actually
                # returns the center (x, y)-coordinates of the bounding
                # box followed by the boxes' width and height
                box = detection[0:4] * np.array([W, H, W, H])
                (centerX, centerY, width, height) = box.astype("int")

                # use the center (x, y)-coordinates to derive the top and
                # and left corner of the bounding box
                x = int(centerX - (width / 2))
                y = int(centerY - (height / 2))

                # update our list of bounding box coordinates, confidences,
                # and class IDs
                if classID == 0: # to make sure model only detects person
                    detections += 1
                    boxes.append([x, y, int(width), int(height)])
                    confidences.append(float(confidence))
                    # classIDs.append(classID)
    return detections

def draw_boxes(image):
    global boxes,confidences
    # apply non-maxima suppression to suppress weak, overlapping bounding
    # boxes
    idxs = cv2.dnn.NMSBoxes(boxes, confidences, 0.5, 0.3)

    # ensure at least one detection exists
    if len(idxs) > 0:
        # loop over the indexes we are keeping
        for i in idxs.flatten():
            # extract the bounding box coordinates
            (x, y) = (boxes[i][0], boxes[i][1])
            (w, h) = (boxes[i][2], boxes[i][3])

            # draw a bounding box rectangle and label on the image
            # color = [int(c) for c in COLORS[classIDs[i]]]
            color = (0,255,0)
            cv2.rectangle(image, (x, y), (x + w, y + h), color, 2)
            text = "{}: {:.4f}".format("person", confidences[i])
            cv2.putText(image, text, (x, y - 5), cv2.FONT_HERSHEY_SIMPLEX,
                0.5, color, 2)
    pil_img = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)
    im_pil = Image.fromarray(pil_img)
    buffer = io.BytesIO()
    im_pil.save(buffer, format="JPEG")
    img_str = buffer.getvalue()
    return img_str

async def message_handler(msg):
    global msgs_received, msgs_sent, params
    subject = msg.subject
    reply = msg.reply
    # data = msg.data.decode()
    # print("Received a message on '{subject} {reply}': {data}".format(
    #    subject=subject, reply=reply))
    msgs_received += 1
    if msgs_received % 100 == 0:
      print('msgs_received: ', msgs_received)
    # print("processed {data}".format(data=_msg.SerializeToString()))
    # ***************** your app's business logic here ********************
    # RFC: We could leverage `reply` topic as the destination topic which would not require DST_NATS_TOPIC to be provided
    # await nc.publish(reply, data)
    # ---------Busines logic-------------
    _msg = xi_iot_pb2.DataStreamMessage()
    _msg.ParseFromString(msg.data)
    image = Image.open(io.BytesIO(_msg.payload)).convert('RGB')
    image_np = np.array(image)
    if msgs_sent % 20 == 0:
        sql = """INSERT INTO detections(room,num)
             VALUES(%s,%s);"""
        d_num = detect(image_np)
        conn = None
        try:
            # connect to the PostgreSQL database
            conn = psycopg2.connect(**params)
            # create a new cursor
            cur = conn.cursor()
            # execute the INSERT statement
            cur.execute(sql, ('HQ1-511-Fuchsia (Cisco Room 14)',d_num,))
            # commit the changes to the database
            conn.commit()
            # close communication with the database
            cur.close()
        except (Exception, psycopg2.DatabaseError) as error:
            print(error)
        finally:
            if conn is not None:
                conn.close()
        print("Detect msg_sent: %d" % msgs_sent)
    image_str = draw_boxes(image_np)
    _msg = xi_iot_pb2.DataStreamMessage()
    _msg.payload = image_str
    print("Msgs_sent: %d" % msgs_sent)
    await nc.publish(dst_nats_topic, _msg.SerializeToString())
    # -------------------------------------
    # print("message published")
    msgs_sent += 1
    if msgs_sent % 100 == 0:
      print('msgs sent: ', msgs_sent)

async def run(loop):
    nats_broker_url, src_nats_topic, dst_nats_topic = get_nats_meta()
    print ("broker: {b}, src topic: {s}, dst_topic: {d}".format(b=nats_broker_url, s=src_nats_topic, d=dst_nats_topic))

    global net
    net = load_net()
    print("net loaded")
    global params
    print("params loaded")
    params = config()
    global nc
    nc = NATS()

    # This will return immediately if the server is not listening on the given URL
    await nc.connect(loop=loop, servers=[str(nats_broker_url)])
    print('connected')
    await nc.subscribe(str(src_nats_topic), cb=message_handler)

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    try:
        loop.run_until_complete(run(loop))
        loop.run_forever()
    finally:
        nc.drain()
        loop.close()
