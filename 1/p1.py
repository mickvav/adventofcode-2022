#!/usr/bin/env python3
m = 0
s = 0
with open("input.txt") as f:
    for line in f:
        if line.strip() == "":
            if s>m:
                m = s
            s=0
            print(m)
        else:
            s+=int(line.strip())
    print(m)
