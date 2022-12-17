#!/usr/bin/env python3
import re

def dist(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])

def parse_file(filename, yscan):
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
    print(sensors)
    print(beacons)
    for pos in range(min_x, max_x + 1):
        can_be = True
        for i in range(len(sensors)):
            d = dist((pos,yscan), sensors[i])
            if d <= distances[i]:
                can_be = False
                break
        if can_be:
            s+=1
    print(max_x - min_x - s )

parse_file("test.txt", 10)
parse_file("input.txt", 2000000)


