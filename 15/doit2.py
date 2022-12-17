#!/usr/bin/env python3
import re
import numpy as np
import random
from scipy import optimize
from scipy.optimize import Bounds
from time import perf_counter

def dist(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])

def scan_line(sensors, distances, yscan, max_x, l):
    for i,s  in enumerate(sensors):
        s=sensors[i]
        dy = abs(yscan - s[1])
        dx = distances[i] - dy
        if dx >= 0:
            l[max(0,s[0] - dx): s[0] + dx + 1] = 0
    if l.any():
        for i,v in enumerate(l):
            if v == 1:
                print(i, yscan)

def fun(X, args):
    x, y = X
    sensors, distances = args
    res = 0
    for i, s in enumerate(sensors):
        d=dist(s, (x,y)) - 1
        if d <= distances[i]:
            res += (distances[i] - d)
    return res

def search_optimum(sensors, distances, yscan, xscan):
    random.seed()
    b = Bounds([0,0], [xscan-1, yscan-1])
    mf = 100000000
    for i in range(1,1000000):
        first_guess = [random.randint(0,xscan),random.randint(0,yscan)]
        res = optimize.minimize(fun, first_guess, [sensors,distances], bounds=b)
        if res.fun < 1:
            print ("X:", res.x)
            X = [round(res.x[0]), round(res.x[1])]
            print (fun(X, [sensors, distances]))
            print(X[0]*4000000 + X[1])
            return X
        if mf > res.fun:
            mf = res.fun
            print(mf)


def parse_file(filename, yscan, xscan):
    sensors = []
    beacons = []
    distances = []
    r = re.compile("Sensor at x=(.*), y=(.*): closest beacon is at x=(.*), y=(.*)")
    with open(filename) as f:
        for line in f:
            m = r.match(line)
            if m:
                sensors.append((int(m.group(1)), int(m.group(2))))
                beacons.append((int(m.group(3)), int(m.group(4))))

    min_x = sensors[0][0]
    max_x = sensors[0][0]
    for i in range(len(sensors)):
        distances.append(dist(sensors[i], beacons[i]))
        min_x = min(sensors[i][0], min_x)
        min_x = min(beacons[i][0], min_x)
        max_x = max(sensors[i][0], max_x)
        max_x = max(beacons[i][0], max_x)
    min_x -= max(distances)
    max_x += max(distances)
    s = 0
    pc = perf_counter()
    l = np.zeros(xscan + 1)
    search_optimum(sensors, distances, yscan, xscan)
parse_file("test.txt", 20, 20)
parse_file("input.txt", 4000000, 4000000)


