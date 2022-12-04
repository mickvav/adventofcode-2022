#!/usr/bin/env python3

def overlaps(r1,r2):
    s1,e1 = [int(i) for i in r1.split('-')] 
    s2,e2 = [int(i) for i in r2.split('-')] 
    if s1 >= s2 and e1 <= e2:
        return True
    if s2 >= s1 and e2 <= e1:
        return True
    if s1 >= s2 and s1 <= e2:
        return True
    if e1 >= s2 and e1 <= e2:
        return True
    if s2 >= s1 and s2 <= e1:
        return True
    if e2 >= s1 and e2 <= e1:
        return True
    return False

with open("input.txt") as f:
    n=0
    for line in f:
        ranges = line.strip().split(',')
        if overlaps(ranges[0], ranges[1]):
            n+=1
    print(n)

