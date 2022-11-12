from sys import stdin, stdout
from lt import encode
import os
import sys

# Stream a fountain of 1024B blocks to stdout
block_size = 1024

filepth = sys.argv[1]
print(filepth)
outputpth = "/home/mininet/LToutput/output.mp4"

with open(filepth, 'rb') as f:
    for block in encode.encoder(f, block_size):
        print(block)
        sys.stdout.flush()
        if os.path.exists(outputpth):
            print("Finish Decoding! Quit")
            break
f1.close()