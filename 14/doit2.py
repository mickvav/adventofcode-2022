#!/usr/bin/env python3


class Map:
    def __init__(self, min_x, max_x, min_y, max_y):
        self.min_x = min_x
        self.max_x = max_x
        self.min_y = min_y
        self.max_y = max_y
        self.v = [[0 for i in range(min_x,max_x+1)] for j in range(min_y, max_y+1)]
        self.sand_volume = 0

    def get(self, x, y):
        if x < self.min_x or x > self.max_x or y < self.min_y or y > self.max_y:
            raise ValueError()
        try:
            return self.v[y - self.min_y][x-self.min_x]
        except IndexError:
            import pdb
            pdb.set_trace()
            print(x,y)
    
    def put(self, x, y, value):
        if x < self.min_x or x > self.max_x or y < self.min_y or y > self.max_y:
            raise ValueError()
        self.v[y - self.min_y][x-self.min_x] = value

    def put_wall(self, w):
        for i in range(len(w)-1):
            if w[i+1][0] == w[i][0]:
                ymin = min(w[i][1], w[i+1][1])
                ymax = max(w[i][1], w[i+1][1])
                for y in range(ymin, ymax+1):
                    self.put(w[i][0],y, 1)
            if w[i+1][1] == w[i][1]:
                xmin = min(w[i][0], w[i+1][0])
                xmax = max(w[i][0], w[i+1][0])
                for x in range(xmin, xmax+1):
                    self.put(x,w[i][1], 1)

    def trace_sand_piece(self) -> bool:
        x = 500
        y = 0
        if self.get(x,y) != 0:
            return False
        try:
            while y <= self.max_y:
                if self.get(x,y+1) == 0:
                    y=y+1
                    continue
                if self.get(x-1,y+1) == 0:
                    x=x-1
                    y=y+1
                    continue
                if self.get(x+1,y+1) == 0:
                    x=x+1
                    y=y+1
                    continue
                self.put(x,y,2)
                self.sand_volume += 1
                return True
            return False
        except ValueError:
            return False

    def count_sand(self) -> int:
        while self.trace_sand_piece():
            pass
        return self.sand_volume
        



def parse_file(filename):
    min_x = 500
    max_x = 500
    min_y = 0
    max_y = 0
    walls = []
    with open(filename) as f:
        for line in f:
            points = line.strip().split(" -> ")
            w = []
            for p in points:
                x,y = p.split(",")
                x= int(x)
                y= int(y)
                min_x = min(x,min_x)
                max_x = max(x,max_x)
                min_y = min(y,min_y)
                max_y = max(y,max_y)
                w.append((x,y))
            walls.append(w)
    floor_y = max_y + 2
    floor_xmin = 500 - floor_y - 2
    floor_xmax = 500 + floor_y + 2
    walls.append([(floor_xmin,floor_y),(floor_xmax,floor_y)])
    min_x = min(min_x,floor_xmin)
    max_x = max(max_x,floor_xmax)
    max_y = max(max_y,floor_y)
    
    m = Map(min_x-1,max_x+1, min_y-1, max_y+1)
    for w in walls:
        m.put_wall(w)
    print("====", m.count_sand())

parse_file("test.txt")
parse_file("input.txt")
