#!/usr/bin/env python3

with open("input.txt") as f:
    line = f.read()
    for i in range(0, len(line)-13):
        s = len(set(list(line[i:i+14])))
        if s == 14:
            print(i+14)
            exit(0)

