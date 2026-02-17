import os
import sys
from PIL import Image

if len(sys.argv) < 3:
    print('Missing arguments. Usage: python3 jpeg_to_png.py <src_dir> <dest_dir>')
    exit(1)

_, src_dir, dest_dir = sys.argv


if not os.path.exists(src_dir):
    print(f'<src_dir> "{src_dir}" does not exist') 
    exit(1)

if not os.path.exists(dest_dir):
    os.makedirs(dest_dir)

jpeg_files = os.listdir(src_dir)

for jpeg in jpeg_files:
    with Image.open(os.path.join(src_dir, jpeg)) as img:
        img.save(os.path.join(dest_dir, jpeg.replace(".jpg", ".png")))
