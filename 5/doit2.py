#!/usr/bin/env python3

from typing import List
def parse_drawing(lines: List[str]) -> List[List[str]]:
    cols = int(len(lines[-1])/4)
    res = [[] for i in range(0, cols)]
    for c in range(0, cols):
        for layer in range(len(lines)-2,-1,-1):
            char = lines[layer][1+4*c]
            if char != ' ':
                res[c].append(char)
    return res

def test_parse():
    lines = []
    with open("test.txt") as f:
        for line in f:
            if line != "\n":
                lines.append(line)
            else:
                res = parse_drawing(lines)
                assert( res == [['Z','N'],['M','C','D'],['P']])

test_parse()


def process_moves(lines : List[str], state: List[List[str]]) -> List[List[str]]:
    for line in lines:
        _m, n, _f, s , _t, d = line.strip().split(' ')
        n=int(n)
        s = int(s) - 1
        d = int(d) - 1
        stack = []
        for op in range(0,n):
            if len(state[s]) > 0:
                w = state[s].pop()
                stack.append(w)
        while len(stack) > 0:
            w = stack.pop()
            state[d].append(w)
    return state

def run(file):
    with open(file) as f:
        lines = []
        for line in f:
            if line != "\n":
                lines.append(line)
            else:
                state = parse_drawing(lines)
                break
        lines = []
        for line in f:
            lines.append(line)
        res_state = process_moves(lines, state)

        tops = [res_state[i][-1] for i in range(0, len(res_state))]
        print("".join(tops))

run("test.txt")
run("input.txt")
 

