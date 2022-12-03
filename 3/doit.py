#!/usr/bin/env python3

a = ord('a')
z = ord('z')
A = ord('A')
Z = ord('Z')

def priority(c) -> int:
    if c >= a and c <= z:
        return 1 + c - a
    return 27 + c - A

score = 0
with open("input.txt") as f:
    for line in f:
        chars = list(line.strip())
        print(len(chars))
        s1 = set(chars[0:int(len(chars)/2)])
        s2 = set(chars[int(len(chars)/2):])
        c = list(s1.intersection(s2))[0]
        p = priority(ord(c))
        score+=p
        print(c, p, score)
