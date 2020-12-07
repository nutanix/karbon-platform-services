# Facial Recognition object 
import os
import cv2
import dlib
import numpy as np
import tensorflow as tf
import pickle
import logging
import onnx
import onnxruntime as ort
from imutils import face_utils
from compute import hard_nms
from onnx_tf.backend import prepare

logging.getLogger().setLevel(logging.DEBUG)

class FaceRecognizer:

	def __init__(self, sess):
		# Preparing predictor 
		onnx_model = onnx.load('ultra_light_640.onnx')
		self.predictor = prepare(onnx_model)
		self.ort_session = ort.InferenceSession('ultra_light_640.onnx')
		self.input_name = self.ort_session.get_inputs()[0].name
		logging.info("ultralight model loaded")

		self.shape_predictor = dlib.shape_predictor('shape_predictor_5_face_landmarks.dat')

		self.f_a = face_utils.facealigner.FaceAligner(self.shape_predictor, desiredFaceWidth=112, desiredLeftEye=(0.3, 0.3))
		self.threshold = 0.63
		with open("embeddings.pkl", "rb") as f:
			self.saved_embeds, self.names = pickle.load(f)
		logging.info("Loaded saved embeddings")

		self.sess = sess
		self.saver = tf.train.import_meta_graph('mfn.ckpt.meta')
		self.saver.restore(self.sess, 'mfn.ckpt')
		logging.info("tensorflow session restored")

		# Setting up tensors
		g = tf.get_default_graph()
		self.images_placeholder = g.get_tensor_by_name('input:0')
		self.embeddings = g.get_tensor_by_name('embeddings:0')
		self.phase_train_placeholder = g.get_tensor_by_name('phase_train:0')
		self.embedding_size = self.embeddings.get_shape()[1]
	
	# Function to label faces after detection 
	def label_faces(self, frame, boxes, labels):
		if len(labels)>0:
			# Adding label to frame
			for i in range(boxes.shape[0]):
				box = boxes[i, :]
				text = f"{labels[i]}"
				x1, y1, x2, y2 = box
				cv2.rectangle(frame, (x1, y1), (x2, y2), (0,150,0), 2)

				# Draw a label with a name below the face
				cv2.rectangle(frame, (x1, y2 - 20), (x2, y2), (0,150,0), cv2.FILLED)
				font = cv2.FONT_HERSHEY_DUPLEX
				cv2.putText(frame, text, (x1 + 6, y2 - 6), font, 0.3, (255, 255, 255), 1)
		else:
			logging.info("No faces detected")
		
		return frame
	
	# Function to detect faces 
	def predict(self, width, height, confidences, boxes, prob_threshold):
		iou_threshold = 0.5 
		top_k = -1
		boxes = boxes[0]
		confidences = confidences[0]
		picked_box_probs = []
		picked_labels = []

		for class_index in range(1, confidences.shape[1]):
			probs = confidences[:, class_index]
			mask = probs > prob_threshold
			probs = probs[mask]
			if not probs.any() or probs.shape[0] == 0:
				continue
			subset_boxes = boxes[mask, :]
			box_probs = np.concatenate([subset_boxes, probs.reshape(-1, 1)], axis=1)
			box_probs = hard_nms(box_probs,
			iou_threshold=iou_threshold,
			top_k=top_k,
			)
			picked_box_probs.append(box_probs)
			picked_labels.extend([class_index] * box_probs.shape[0])
		if not picked_box_probs:
			return np.array([]), np.array([]), np.array([])
		picked_box_probs = np.concatenate(picked_box_probs)
		picked_box_probs[:, 0] *= width
		picked_box_probs[:, 1] *= height
		picked_box_probs[:, 2] *= width
		picked_box_probs[:, 3] *= height
		return picked_box_probs[:, :4].astype(np.int32), np.array(picked_labels), picked_box_probs[:, 4]
	
	# Function to process frame 
	def process_frame(self, frame):
		# preprocess faces
		h, w, _ = frame.shape
		img = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
		img = cv2.resize(img, (640, 480))
		img_mean = np.array([127, 127, 127])
		img = (img - img_mean) / 128
		img = np.transpose(img, [2, 0, 1])
		img = np.expand_dims(img, axis=0)
		img = img.astype(np.float32)

		# detect faces
		confidences, boxes = self.ort_session.run(None, {self.input_name: img})
		boxes, labels, probs = self.predict(w, h, confidences, boxes, 0.8)

		# Normalize face detections
		faces = []
		boxes[boxes<0] = 0
		for i in range(boxes.shape[0]):
			box = boxes[i, :]
			x1, y1, x2, y2 = box
			gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
			aligned_face = self.f_a.align(frame, gray, dlib.rectangle(left = x1, top=y1, right=x2, bottom=y2))
			aligned_face = cv2.resize(aligned_face, (112, 112))
			aligned_face = aligned_face - 127.5
			aligned_face = aligned_face * 0.0078125
			faces.append(aligned_face)
		
		predictions = []
		faces = np.array(faces)
		feed_dict = {self.images_placeholder: faces, self.phase_train_placeholder: False}
		embeds = self.sess.run(self.embeddings, feed_dict=feed_dict)

		# Determine prediction based on distance
		for embedding in embeds:
			diff = np.subtract(self.saved_embeds, embedding)
			dist = np.sum(np.square(diff), 1)
			idx = np.argmin(dist)
			if dist[idx] < self.threshold:
				predictions.append(self.names[idx])
			else:
				predictions.append("unknown")
		
		return boxes, predictions

	# Call function 
	def __call__(self, img_np):
		return self.process_frame(img_np)