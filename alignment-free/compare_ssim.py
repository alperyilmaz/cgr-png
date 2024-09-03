import sys
import cv2
from skimage.metrics import structural_similarity as ssim
import numpy as np
def compare_images(img1_path, img2_path):
    img1 = cv2.imread(img1_path)
    img2 = cv2.imread(img2_path)

    # Structural Similarity Index (SSIM)
    return ssim(img1, img2, channel_axis=2)
img1_path = sys.argv[1]
img2_path = sys.argv[2]
result = compare_images(img1_path, img2_path)
print(f'{result}')
