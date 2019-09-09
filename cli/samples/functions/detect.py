#!/usr/bin/python


import cStringIO
import json
import io
import numpy as np
from PIL import Image
from PIL import ImageDraw
from PIL import ImageFont
import tensorflow as tf
import logging
import threading
from Queue import Queue

BASE_PATH = "/mllib/objectdetection"

# ssd_inception_v2_coco	latency - 42ms
PATH_TO_CKPT = BASE_PATH + '/ssd_inception_v2_coco_2017_11_17/frozen_inference_graph.pb'
PATH_TO_LABELS = BASE_PATH + '/mscoco_label_map.pbtxt'

boxes = None
scores = None
classes = None
num_detections = None
frame_count = 0

labels = ["dummy",
          "person",
          "bicycle",
          "car",
          "motorcycle",
          "airplane",
          "bus",
          "train",
          "truck",
          "boat",
          "traffic light",
          "fire hydrant",
          "street sign",
          "stop sign",
          "parking meter",
          "bench",
          "bird",
          "cat",
          "dog",
          "horse",
          "sheep",
          "cow",
          "elephant",
          "bear",
          "zebra",
          "giraffe",
          "hat",
          "backpack",
          "umbrella",
          "shoe",
          "eye glasses",
          "handbag",
          "tie",
          "suitcase",
          "frisbee",
          "skis",
          "snowboard",
          "sports ball",
          "kite",
          "baseball bat",
          "baseball glove",
          "skateboard",
          "surfboard",
          "tennis racket",
          "bottle",
          "plate",
          "wine glass",
          "cup",
          "fork",
          "knife",
          "spoon",
          "bowl",
          "banana",
          "apple",
          "sandwich",
          "orange",
          "broccoli",
          "carrot",
          "hot dog",
          "pizza",
          "donut",
          "cake",
          "chair",
          "couch",
          "potted plant",
          "bed",
          "mirror",
          "dining table",
          "window",
          "desk",
          "toilet",
          "door",
          "tv",
          "laptop",
          "mouse",
          "remote",
          "keyboard",
          "cell phone",
          "microwave",
          "oven",
          "toaster",
          "sink",
          "refrigerator",
          "blender",
          "book",
          "clock",
          "vase",
          "scissors",
          "teddy bear",
          "hair drier",
          "toothbrush",
          "hair brush"]


class ObjectDetector(object):

    def __init__(self):
        self.detection_graph = self._build_graph()
        self.sess = tf.Session(graph=self.detection_graph)
        logging.info("Loaded tensorflow gpu")

    def _build_graph(self):
        detection_graph = tf.Graph()
        with detection_graph.as_default():
            od_graph_def = tf.GraphDef()
            with tf.gfile.GFile(PATH_TO_CKPT, 'rb') as fid:
                serialized_graph = fid.read()
                od_graph_def.ParseFromString(serialized_graph)
                tf.import_graph_def(od_graph_def, name='')

        return detection_graph

    def _load_image_into_numpy_array(self, image):
        (im_width, im_height) = image.size
        return np.array(image.getdata()).reshape(
            (im_height, im_width, 3)).astype(np.uint8)

    def detect(self, image):
        image_np = self._load_image_into_numpy_array(image)
        image_np_expanded = np.expand_dims(image_np, axis=0)

        graph = self.detection_graph
        image_tensor = graph.get_tensor_by_name('image_tensor:0')
        boxes = graph.get_tensor_by_name('detection_boxes:0')
        scores = graph.get_tensor_by_name('detection_scores:0')
        classes = graph.get_tensor_by_name('detection_classes:0')
        num_detections = graph.get_tensor_by_name('num_detections:0')

        (boxes, scores, classes, num_detections) = self.sess.run(
            [boxes, scores, classes, num_detections],
            feed_dict={image_tensor: image_np_expanded})

        boxes, scores, classes, num_detections = map(
            np.squeeze, [boxes, scores, classes, num_detections])

        return boxes, scores, classes.astype(int), num_detections

# one of the few fonts available inside tensorflow-python image
font = ImageFont.truetype("DejaVuSansMono.ttf", 18)

def draw_bounding_box_on_image(image, box, color='green', thickness=2, display_str_list=()):
    draw = ImageDraw.Draw(image)
    im_width, im_height = image.size
    ymin, xmin, ymax, xmax = box
    (left, right, top, bottom) = (xmin * im_width, xmax * im_width,
                                  ymin * im_height, ymax * im_height)
    draw.line([(left, top), (left, bottom), (right, bottom),
               (right, top), (left, top)], width=thickness, fill=color)

    # If the total height of the display strings added to the top of the bounding
    # box exceeds the top of the image, stack the strings below the bounding box
    # instead of above.
    display_str_heights = [font.getsize(ds)[1] for ds in display_str_list]
    # Each display_str has a top and bottom margin of 0.05x.
    total_display_str_height = (1 + 2 * 0.05) * sum(display_str_heights)

    if top > total_display_str_height:
        text_bottom = top
    else:
        text_bottom = bottom + total_display_str_height
    # Reverse list and print from bottom to top.
    for display_str in display_str_list[::-1]:
        text_width, text_height = font.getsize(display_str)
        margin = np.ceil(0.05 * text_height)
        draw.rectangle([(left, text_bottom - text_height - 2 * margin),
                        (left + text_width, text_bottom)],
                       fill='#216df3')
        draw.text((left + margin, text_bottom - text_height - margin), display_str, fill="white",
                  font=font)
        text_bottom -= text_height - 2 * margin


def detect_objects(data):
    global boxes
    global scores
    global classes
    global num_detections
    global frame_count
    image = Image.open(io.BytesIO(data)).convert('RGB')

    basewidth = 300
    wpercent = (basewidth/float(image.size[0]))
    hsize = int((float(image.size[1])*float(wpercent)))
    image = image.resize((basewidth,hsize), Image.ANTIALIAS)

    if frame_count % 24 == 0:
        boxes, scores, classes, num_detections = client.detect(image)
    num_boxes = 0
    if frame_count % 1 == 0:
        for i in range(num_detections):
            if num_boxes > 3:
                break
            cls = classes[i]
            if scores[i] < 0.2:
                continue
            draw_bounding_box_on_image(
                image, boxes[i], display_str_list=[labels[cls]])
            num_boxes = num_boxes + 1
    buffer = cStringIO.StringIO()
    image.save(buffer, format="JPEG")
    img_str = buffer.getvalue()
    frame_count = frame_count + 1
    
    if frame_count % 100 == 0:
        logging.info(frame_count)
    return img_str

def read_frame():
    while True:
        frame = inQ.get()
        output_img = detect_objects(frame)
        outQ.put(output_img)

client = ObjectDetector()
inQ = Queue(200)
outQ = Queue(200)

read_thr = threading.Thread(target=read_frame)
read_thr.name = "inferenceRead"
read_thr.daemon = True
read_thr.start()

def main(ctx, msg):
    if inQ.full():
        inQ.get()
    inQ.put(msg)

    while not outQ.empty():
        output_img = outQ.get()
        ctx.send(output_img)
    return
