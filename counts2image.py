import sys
import numpy as np
import matplotlib.pyplot as plt
import io
from PIL import Image

def kmer_to_coordinates(kmer):
    mapping = {'A': (0, 0), 'G': (1, 0), 'C': (0, 1), 'T': (1, 1)}
    x_bin = ''.join(str(mapping[base][0]) for base in kmer)
    y_bin = ''.join(str(mapping[base][1]) for base in kmer)
    x_coord = int(x_bin, 2)
    y_coord = int(y_bin, 2)
    return x_coord, y_coord

def count2rgb(count):
    return tuple(np.uint8(np.array([int(count) >> 16 & 0xff, int(count) >> 8 & 0xff, int(count) & 0xff])))

if len(sys.argv) == 1:
    print("Please provide DNA length.")
    sys.exit(1)
else:
    len_seq = int(sys.argv[1])


size = 2**len_seq
#img = Image.new('RGB', (size, size))
my_img = np.zeros((size,size,3), dtype='uint8')
img_stdout=Image.new('RGB',(size,size))
for line in sys.stdin:
    seq, count = line.strip().split("\t")
    
    # Map sequence to matrix location
    X,Y=kmer_to_coordinates(seq)
    rgb= count2rgb(count)
    my_img[X,Y]=rgb

# Save the image
## IMPORTANT: We take transpose of the numpy array to match ACGT bit mapping
transposed_img=np.transpose(my_img, axes=(1, 0, 2))
#plt.imshow(transposed_img)
#plt.imsave("my_image.png",transposed_img)
buf = io.BytesIO()
#plt.imsave(io.BytesIO(), transposed_img, format='PNG')
Image.fromarray(transposed_img).save(buf, format='PNG')
sys.stdout.buffer.write(buf.getvalue())
#img.save(sys.stdout.buffer, format='PNG')
