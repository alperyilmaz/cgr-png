#!/usr/bin/env python3

from math import log
from PIL import Image
import sys

# Define DNA sequence mapping
seq = {
    0: {0: "A", 1: "C"},
    1: {0: "G", 1: "T"}
}

# Define function to convert coordinates to DNA kmer
def coord_to_seq(x, y):
    xarr = list(format(x, '0{}b'.format(base)))
    yarr = list(format(y, '0{}b'.format(base)))
    return ''.join(seq[int(xarr[i])][int(yarr[i])] for i in range(base))

# Specify image file
if len(sys.argv) > 1:
    file = sys.argv[1]
else:
    print("No filename provided. Please provide image filename.")

# Open image file
img = Image.open(file)
width, height = img.size

# Calculate base for DNA kmer conversion
base = int(log(width, 2))

# Iterate over pixels in the image
for y in range(height):
    for x in range(width):
        r, g, b = img.getpixel((x, y))
        rgb = (r << 16) + (g << 8) + b
        print(f'{coord_to_seq(x, y)}\t{rgb}')
