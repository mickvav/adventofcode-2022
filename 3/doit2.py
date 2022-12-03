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
    group = []
    for line in f:
        chars = list(line.strip())
        group.append(set(chars))
        if len(group) == 3:
            s = group[0].intersection(group[1]).intersection(group[2])
            print(list(s)[0])
            score += priority(ord(list(s)[0]))
            group = []
            print(score)

