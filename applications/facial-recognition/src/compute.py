# Computational API for cv calculations
import os
import numpy as np

# -------------------------------------------------
# Compute the areas of rectangles given two corners.
# Args:
#     left_top (N, 2): left top corner.
#     right_bottom (N, 2): right bottom corner.
# Returns:
#     area (N): return the area.
# -------------------------------------------------
def area_of(left_top, right_bottom):
	hw = np.clip(right_bottom - left_top, 0.0, None)
	return hw[..., 0] * hw[..., 1]

# -------------------------------------------------
# Return intersection-over-union (Jaccard index) of boxes.
# Args:
#     boxes0 (N, 4): ground truth boxes.
#     boxes1 (N or 1, 4): predicted boxes.
#     eps: a small number to avoid 0 as denominator.
# Returns:
#     iou (N): IoU values.
# -------------------------------------------------
def jaccard_index(boxes0, boxes1, eps=1e-5):
	overlap_left_top = np.maximum(boxes0[..., :2], boxes1[..., :2])
	overlap_right_bottom = np.minimum(boxes0[..., 2:], boxes1[..., 2:])

	overlap_area = area_of(overlap_left_top, overlap_right_bottom)
	area0 = area_of(boxes0[..., :2], boxes0[..., 2:])
	area1 = area_of(boxes1[..., :2], boxes1[..., 2:])
	return overlap_area / (area0 + area1 - overlap_area + eps)

# -------------------------------------------------
# Perform hard non-maximum-supression to filter out boxes with iou greater
# than threshold
# Args:
#     box_scores (N, 5): boxes in corner-form and probabilities.
#     iou_threshold: intersection over union threshold.
#     top_k: keep top_k results. If k <= 0, keep all the results.
#     candidate_size: only consider the candidates with the highest scores.
# Returns:
#     picked: a list of indexes of the kept boxes
# -------------------------------------------------
def hard_nms(box_scores, iou_threshold, top_k=-1, candidate_size=200):
	scores = box_scores[:, -1]
	boxes = box_scores[:, :-1]
	picked = []
	indexes = np.argsort(scores)
	indexes = indexes[-candidate_size:]
	while len(indexes) > 0:
		current = indexes[-1]
		picked.append(current)
		if 0 < top_k == len(picked) or len(indexes) == 1:
			break
		current_box = boxes[current, :]
		indexes = indexes[:-1]
		rest_boxes = boxes[indexes, :]
		iou = jaccard_index(
			rest_boxes,
			np.expand_dims(current_box, axis=0),
		)
		indexes = indexes[iou <= iou_threshold]
	
	return box_scores[picked, :]