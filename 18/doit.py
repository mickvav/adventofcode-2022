#!/usr/bin/env python3

from typing import List
def neigh(v: tuple) -> List[tuple]:
    return [
        (v[0]+1,v[1],v[2]),
        (v[0]-1,v[1],v[2]),
        (v[0],v[1]+1,v[2]),
        (v[0],v[1]-1,v[2]),
        (v[0],v[1],v[2]+1),
        (v[0],v[1],v[2]-1)
    ]
def parse_file(filename):
    res = {}
    with open(filename) as f:
        for line in f:
            parts=line.strip().split(',')
            v = tuple(int(i) for i in parts)
            res[v] = 6
        for i,v in res.items():
            for n in neigh(i):
                if n in res:
                    res[i] -= 1
        print( sum(res.values()))

parse_file("test.txt")
parse_file("input.txt")

