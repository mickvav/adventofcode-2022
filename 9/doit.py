#!/usr/bin/env python3

from typing import List
H=[0,0]
T=[0,0]

def get_next_T(H,T) -> List[int]:
    T1 = [T[0], T[1]]
    if abs(H[0]-T[0]) == 1 and abs(H[1]-T[1]) == 1:
        return T1
    if H[0] == T[0]:
        if H[1] - T[1] > 1:
            T1[1] = H[1]-1
            return T1
        if T[1] - H[1] > 1:
            T1[1] = H[1]+1
            return T1
        return T1
    if H[1] == T[1]:
        if H[0] - T[0] > 1:
            T1[0] = H[0]-1
            return T1
        if T[0] - H[0] > 1:
            T1[0] = H[0]+1
            return T1
        return T1
    if H[0] > T[0]:
        T1[0] += 1
        if  H[1] > T[1]:
            T1[1] += 1
        elif H[1] < T[1]:
            T1[1] -= 1
        return T1
    if H[0] < T[0]:
        T1[0] -= 1
        if H[1] > T[1]:
            T1[1] += 1
        elif H[1] < T[1]:
            T1[1] -= 1
        return T1

def parse_file(filename):
    with open(filename) as f:
        H=[0,0]
        T=[0,0]
        Log=[]
        for line in f:
            C, L = line.strip().split(" ")
            L = int(L)
            while L>0:
                L-=1
                if C == 'R':
                    H[0] += 1
                if C == 'L':
                    H[0] -= 1
                if C == 'U':
                    H[1] += 1
                if C == 'D':
                    H[1] -= 1
                T1 = get_next_T(H,T)
                Log.append(tuple(T1))
                T = T1
        print(len(set(Log)))

parse_file("test.txt")
parse_file("input.txt")




