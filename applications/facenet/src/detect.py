# Copyright (c) 2019 Nutanix, Inc.
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

from align.detect_face import create_mtcnn, detect_face

import cv2
import numpy as np
import tensorflow as tf

from collections import namedtuple
import logging
import time

INPUT_IMAGE_SIZE = 160

class FaceEmbedder:
    def __init__(self, sess, facenet_model_pb):
        self.sess = sess
        with tf.gfile.GFile(facenet_model_pb, 'rb') as f:
            graph_def = tf.GraphDef()
            graph_def.ParseFromString(f.read())
            tf.import_graph_def(graph_def, name='')
        g = tf.get_default_graph()
        self.images_placeholder = g.get_tensor_by_name('input:0')
        self.embeddings = g.get_tensor_by_name('embeddings:0')
        self.phase_train_placeholder = g.get_tensor_by_name('phase_train:0')
        self._input_image_size = INPUT_IMAGE_SIZE

    def input_image_size(self):
        return self._input_image_size

    def _standardize(self, x):
        return (x - 127.5) * 0.0078125

    def __call__(self, image):
        reshaped = image.reshape(-1, self._input_image_size, self._input_image_size, 3)
        reshaped = self._standardize(reshaped)
        feed_dict = {
            self.images_placeholder: reshaped,
            self.phase_train_placeholder: False
        }
        return self.sess.run(self.embeddings, feed_dict=feed_dict)

class FaceDetector:
    def __init__(self, sess, mtcnn_model_dir=None):
        self.pnet, self.rnet, self.onet = create_mtcnn(sess, mtcnn_model_dir)

    def __call__(self, img, minsize=20, threshold=[0.6, 0.7, 0.7], factor=0.709):
        """ generates bounding boxes containing detected faces """
        bounding_boxes, _ = detect_face(img, minsize, self.pnet, self.rnet, self.onet, threshold, factor)
        return (np.squeeze(bbox[0:4]) for bbox in bounding_boxes)

def _add_margin(bb0, margin, img_shape):
    """ expands the bounding box by margin/2 in all directions """
    bb = np.zeros(4, dtype=np.int32)
    bb[0] = np.maximum(bb0[0] - margin // 2, 0) # left
    bb[1] = np.maximum(bb0[1] - margin // 2, 0) # top
    bb[2] = np.minimum(bb0[2] + margin // 2, img_shape[1]) # right
    bb[3] = np.minimum(bb0[3] + margin // 2, img_shape[0]) # bottom
    return bb

def crop(img, bb, margin):
    l, t, r, b = _add_margin(bb, margin, img.shape)
    return img[t:b, l:r, :]

def _detect_faces(img, face_detector, margin=44, **kwargs):
    for bb in face_detector(img, **kwargs):
        yield bb, crop(img, bb, margin)

Face = namedtuple('Face', ['id', 'v', 'last_seen', 'times_seen'])

class FaceRecognizer:
    def __init__(self, sess, embedding_model_path, max_identities=50, min_cos=0.5):
        self.reset()

        self.min_cos = min_cos
        self.max_identities = max_identities

        self.embedder = FaceEmbedder(sess, embedding_model_path)
        self.detector = FaceDetector(sess)

    def reset(self):
        self.cache = []
        self.new_face_id = 0

    def _register_new(self, v):
        f = Face(self.new_face_id, v, time.time(), 1)
        self.new_face_id += 1
        return f

    def _register_repeated(self, face, v):
        v = (face.times_seen * face.v + v) / (face.times_seen + 1)
        v /= np.linalg.norm(v)
        return Face(face.id, v, time.time(), face.times_seen + 1)

    def _find_previous(self, v):
        """ returns index of the closest match in self.cache, or None if there's no good match """
        most_similar = None
        max_cos = 0
        for i, f in enumerate(self.cache):
            dp = np.dot(f.v, v)
            if dp > max_cos:
                max_cos = dp
                most_similar = i
        if max_cos < self.min_cos:
            return None
        return most_similar

    def _drop_old_entries(self):
        if len(self.cache) > self.max_identities:
            self.cache.sort(key=lambda f: -f.last_seen)
            self.cache = self.cache[:self.max_identities]

    def _color(self, id):
        import random
        random.seed(id * id * id)
        r = random.randint(0, 255)
        g = random.randint(0, 255)
        b = random.randint(0, 255)
        return (b, g, r)

    def bboxes(self, img, bb_threshold=40, **kwargs):
        result = []
        img_size = self.embedder.input_image_size()
        for bb, face in _detect_faces(img, self.detector, **kwargs):
            if min(face.shape[0], face.shape[1]) < img_size - bb_threshold:
                logging.debug("Skipping face because the detected box is small {}".format(face.shape))
                continue

            face = cv2.resize(face, (img_size, img_size), interpolation=cv2.INTER_CUBIC)

            v = np.squeeze(self.embedder(face))
            idx = self._find_previous(v)
            if idx is None:
                self.cache.append(self._register_new(v))
                self._drop_old_entries() # TODO check that it works
                idx = len(self.cache) - 1
            else:
                self.cache[idx] = self._register_repeated(self.cache[idx], v)
            result.append((idx, bb))
        return result

    def draw_overlay(self, img, bboxes, id2color=None):
        output_img = img.copy()
        for idx, bb in bboxes:
            f = self.cache[idx]
            color = (id2color or self._color)(f.id)
            bb = bb.astype(int)
            cv2.rectangle(output_img, (bb[0], bb[1]), (bb[2], bb[3]), color, 5)
            cv2.putText(output_img, "ID={}".format(idx), (bb[0], bb[1]),
                        cv2.FONT_HERSHEY_DUPLEX, 0.7, (255, 192, 192), 1, cv2.LINE_AA)
        return output_img

    def __call__(self, img, id2color=None):
        return self.draw_overlay(img, self.bboxes(img), id2color)
