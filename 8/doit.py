#!/usr/bin/env python3

from typing import List

def readmap(filename) -> List[List[int]]:
    res = []
    with open(filename) as f:
        for line in f:
            row = [int(c) for c in list(line.strip())]
            res.append(row)
    return res

def mark_visible(res) -> List[List[int]]:
    marks = []
    for row in res:
        m = -1
        newrow = []
        for v in row:
            if m < v:
               m = v
               newrow.append(1)
            else:
               newrow.append(0)
        m = -1
        for idx in reversed(range(len(row))):
            if m < row[idx]:
                m = row[idx]
                newrow[idx] = 1
        marks.append(newrow)
    for col in range(len(newrow)):
        m = -1
        for row in range(len(marks)):
            if m < res[row][col]:
                m = res[row][col]
                marks[row][col] = 1
        m = -1
        for row in reversed(range(len(marks))):
            if m < res[row][col]:
                m = res[row][col]
                marks[row][col] = 1
    return marks

def count_visible(marks) -> int:
    s = 0
    for row in marks:
        s += sum(row)
    return s

def scenic_score(res, r, c):
    s1 = r
    h = res[r][c]
    for r1 in reversed(range(r)):
        if res[r1][c] >= h:
            s1 = r - r1
            break
    s2 = len(res) - r - 1
    for r1 in range(r+1, len(res)):
        if res[r1][c] >= h:
            s2 = r1 - r
            break
    s3 = c
    for c1 in reversed(range(c)):
        if res[r][c1] >= h:
            s3 = c - c1
            break
    s4 = len(res[r]) - c - 1
    for c1 in range(c+1, len(res[r])):
        if res[r][c1] >= h:
            s4 = c1 - c
            break
    return s1 * s2 * s3 * s4

def max_score(res):
    m = 0
    for r in range(1,len(res)-1):
        for c in range(1, len(res[r]) - 1):
            mp = scenic_score(res, r,c)
            if mp > m:
                m = mp
    return(m)


h = readmap("test.txt")
print(h)

m = mark_visible(h)
print(count_visible(m))
print(max_score(h))
print('----')
        
h = readmap("input.txt")
m = mark_visible(h)
print(count_visible(m))
print(max_score(h))
                 
