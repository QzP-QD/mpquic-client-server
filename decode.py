from sys import stdin, stdout
from lt import decode

import fcntl
import os

fcntl.fcntl(stdin, fcntl.F_SETFL, os.O_NONBLOCK)

filepth = "/home/mininet/LToutput/tmpfile"


f2 = open('/home/mininet/LToutput/x.txt', 'w')

while True:
    if not os.path.exists(filepth):
        continue
    elif os.stat(filepth).st_size > 0:
        break

with open(filepth, 'rb') as f1:
    decoder = decode.LtDecoder()
    for block in decode.read_blocks(f1):
        print("Decoding...")
        f2.write("123\n")
        f2.flush()
        decoder.consume_block(block)
        if decoder.is_done():
            print("Decode Success!")
            break

# while True :
#     try:
#         cc = sys.stdin.read()
#     except TypeError as e:
#             print('no std input readed')
#     else:
#         print(cc)
#         break

# for block in decode.read_blocks(stdin.read()):
#     print("Decoding...")
#     f1.write(block)
#     f1.flush()
#
#     decoder.consume_block(block)
#     if decoder.is_done():
#         print("Decode Success!")
#         break

# You can collect the decoded transmission as bytes
data = decoder.bytes_dump()

f = open('/home/mininet/LToutput/output.mp4', 'wb')
f.write(data)
