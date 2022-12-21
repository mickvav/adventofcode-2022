#!/usr/bin/env python3

from collections import defaultdict

class InitialShape:
    def __init__(self, c):
        if c == "-":
            self.pattern=[16 + 8 + 4 + 2    ]
        if c == "+":
            self.pattern=[     8            ,
                          16 + 8 + 4        ,
                               8            
                          ]
        if c == "L":
            self.pattern=[16 + 8 + 4,
                                   4,
                                   4,
                          ]
        if c == "I":
            self.pattern=[16, 16, 16, 16]
        if c == "o":
            self.pattern=[16 + 8, 16+8]
    def shape(self, h):
        return Shape(self.pattern.copy(), h)


class Shape:
    def __init__(self, pattern, h):
        self.pattern = pattern
        self.h = h

    def Left(self, g):
        new_pattern = [p << 1 for p in self.pattern]
        if any(p > 127 for p in new_pattern):
            return
        if self.check_collision(new_pattern, self.h, g):
            return
        self.pattern = new_pattern

    def __repr__(self):
        res = ''
        for i,l in enumerate(reversed(self.pattern)):
            res += '|' + ''.join('@' if (l >> x) % 2 else ' ' for x in range(6,-1,-1)) + f"| {self.h + len(self.pattern) - i}\n"
        return res


    def check_collision(self, pattern, h, g):
        collision = False
        for dh in range(0,len(pattern)):
            h_in_pool = h + dh 
            if h_in_pool < len(g.pool):
                intersection = g.pool[h_in_pool] & pattern[dh]
                if intersection > 0:
                    collision = True
                    break
            else:
                break
        return collision


    def Right(self, g):
        if any(p % 2 == 1 for p in self.pattern):
            return
        new_pattern = [p >> 1 for p in self.pattern]
        if self.check_collision(new_pattern, self.h, g):
            return
        self.pattern = new_pattern

    def Down(self, g) -> bool:
        if self.h == 0:
            for h in range(0, len(self.pattern)):
                if h < len(g.pool):
                    g.pool[h] |= self.pattern[h]
                else:
                    g.pool.append(self.pattern[h])
            return False
        collision = self.check_collision(self.pattern, self.h - 1, g)
        if collision:
            for dh in range(0,len(self.pattern)):
                h_in_pool = self.h + dh
                if h_in_pool < len(g.pool):
                    g.pool[h_in_pool] |= self.pattern[dh]
                else:
                    g.pool.append(self.pattern[dh])
            return False
        self.h -= 1
        return True
            

class Glass:
    def __init__(self):
        self.pool = []

    def __repr__(self):
        res = ''
        for i,l in enumerate(reversed(self.pool)):
            res += '|' + ''.join('#' if (l >> x) % 2 else ' ' for x in range(6,-1,-1)) + f"| {len(self.pool) - i}\n"
        res += "+-------+"
        return res

inits = [InitialShape(i) for i in ["-", "+", "L","I","o"]]

class predictor:
    def __init__(self, states):
        self.states = states
        self.repeating_tuples = [tpl for tpl in states if len(states[tpl]) > 1]
        self.ri_period = max(states[tpl][1][2] - states[tpl][0][2] for tpl in self.repeating_tuples)


    def predict(self, ri):
        ri0 = ri % self.ri_period
        for tpl in self.repeating_tuples:
#            if self.states[tpl][0][2] % self.ri_period == ri0:
#                h0 = self.states[tpl][0][1]
#                dh = self.states[tpl][1][1] - h0
#                nr = (ri - self.states[tpl][0][2]) // self.ri_period
#                return h0 + nr * dh
            if self.states[tpl][-1][2] % self.ri_period == ri0:
                h0 = self.states[tpl][-2][1]
                dh = self.states[tpl][-1][1] - h0
                nr = (ri - self.states[tpl][-2][2]) // self.ri_period
                return h0 + nr * dh
        return -1
        


def process_file(filename, nops):
    with open(filename) as f:
        line = f.readline()
    line = line.strip()
    r = None
    ri = 0
    g = Glass()
    i = 0
    print(f"line: {line} {len(line)}")
    states = defaultdict(list)
    pred = None
    while ri <= nops:
        op = line[i % len(line)]
        if r is None:
            h=len(g.pool) + 3
            r = inits[ri % len(inits)].shape(h)
            ri += 1
            if (len(g.pool)>40):
                tpl = tuple(g.pool[-40:] + [ri % len(inits), i % len(line)])
                if len(states[tpl]) > 1:
                    pred = predictor(states)

                    exp = pred.predict(ri)
                    if exp == -1:
                        exp = pred.predict(ri)
                    print(f"actual: {len(g.pool)} predicted: {exp}")
                    print(f"period: {i} {ri} {states[tpl][1][0] - states[tpl][0][0]}, {states[tpl][1][1] - states[tpl][0][1]} {states[tpl][1][2] - states[tpl][0][2]}")
                    e1 =  pred.predict(1000000000001)
                                       
                    if e1 != -1:
                        print(e1)
                        return
                states[tpl].append([i,len(g.pool), ri])


        if op == ">":
            r.Right(g)
        if op == "<":
            r.Left(g)
        if not(r.Down(g)):
            r = None
        i += 1
    print(len(g.pool))
            

#process_file("test.txt", 20000)
process_file("input.txt", 202200)


