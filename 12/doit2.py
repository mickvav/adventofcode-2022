#!/usr/bin/env python3
from typing import List

class map:
    def __init__(self, filename):
        self.res = []
        self.distances = []
        m = {}
        m['S'] = 0
        m['E'] = ord('z') - ord('a')
        for i in range(ord('a'), ord('z') + 1 ):
            m[chr(i)] = i - ord('a')
        row_n = 0
        with open(filename) as f:
            for line in f:
                line = line.strip()
                self.res.append([m[c] for c in line])
                self.distances.append([-1 for c in line])
                if 'S' in line:
                    self.S_row = row_n
                    self.S_col = line.find('S')
                if 'E' in line:
                    self.E_row = row_n
                    self.E_col = line.find('E')
                    self.distances[row_n][self.E_col] = 0
                row_n += 1
        while(self.dijkstra_step() > 0):
            print(".", end='')
        print(self.distances[self.S_row][self.S_col])

    def shortest_a(self) -> int:
        known = self.distances[self.S_row][self.S_col]

        for row in range(len(self.distances)):
            for col in range(len(self.distances[row])):
                if self.res[row][col] == 0 and self.distances[row][col] < known and self.distances[row][col] >= 0:
                    known = self.distances[row][col]
        return known

        
    def dijkstra_step(self) -> int:
        ops = 0
        for row in range(len(self.distances)):
            for col in range(len(self.distances[row])):
                if self.distances[row][col] >= 0:
                    cdp1 = self.distances[row][col] + 1
                    if row > 0: # ^
                        if (self.distances[row-1][col] < 0 or self.distances[row-1][col] > cdp1 ) and self.res[row][col] - self.res[row-1][col] <= 1:
                            self.distances[row-1][col] = self.distances[row][col] + 1
                            ops += 1
                    if row < len(self.distances) - 1: # v
                        if (self.distances[row+1][col] < 0 or self.distances[row+1][col] > cdp1 ) and self.res[row][col] - self.res[row+1][col] <= 1:
                            self.distances[row+1][col] = self.distances[row][col] + 1
                            ops += 1
                    if col > 0: # <
                        if (self.distances[row][col-1] < 0 or self.distances[row][col-1] > cdp1 ) and self.res[row][col] - self.res[row][col-1] <= 1:
                            self.distances[row][col-1] = self.distances[row][col] + 1
                            ops += 1
                    if col < len(self.distances[row]) - 1:
                        if (self.distances[row][col+1] < 0 or self.distances[row][col+1] > cdp1 )and self.res[row][col] - self.res[row][col+1] <= 1:
                            self.distances[row][col+1] = self.distances[row][col] + 1
                            ops += 1
        return ops






m = map("test.txt")
m1 = map("input.txt")
print(m.shortest_a())
print(m1.shortest_a())
print(m.res)
print(m.distances)
