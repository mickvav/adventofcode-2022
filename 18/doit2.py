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

class vmap:
    def __init__(self, res):
        mins = []
        maxs = []
        for v in res:
            if mins == []:
                mins = list(v)
                maxs = list(v)
                continue
            for i in 0,1,2:
                mins[i] = min(mins[i],v[i])
                maxs[i] = max(maxs[i],v[i])
        self.mins = tuple(i-1 for i in mins)
        self.maxs = tuple(i+1 for i in maxs)
        self.colors = {v: 0 for v in res}
        self.res = res
        self.color()

    def get_color(self, v):
        if v in self.colors:
            return self.colors[v]
        return 0

    def in_box(self,v):
        for i in [0,1,2]:
            if v[i] < self.mins[i]:
                return False
            if v[i] > self.maxs[i]:
                return False
        return True

    def color(self):
        self.colors[self.mins] = 1
        ops = 1
        while ops > 0:
            ops = 0
            new_colors = set()
            for v,c in self.colors.items():
                if c == 1 and v not in self.res:
                    for n in neigh(v):
                        if not(self.in_box(n)):
                            continue
                        if n not in self.res and self.get_color(n) == 0:
                            new_colors.add(n)
                    self.colors[v] = 2
            ops = len(new_colors)
            for n in new_colors:
                self.colors[n]=1




        
        
def parse_file(filename):
    res = {}
    with open(filename) as f:
        for line in f:
            parts=line.strip().split(',')
            v = tuple(int(i) for i in parts)
            res[v] = 6

        print(len(res))
        CM = vmap(res)
        for i,v in res.items():
            for n in neigh(i):
                if n in res or CM.get_color(n) == 0:
                    res[i] -= 1
        print( sum(res.values()))

parse_file("test.txt")
parse_file("input.txt")

