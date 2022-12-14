#!/usr/bin/env python3

def get_sig_strength(X: int,C: int) -> int:
    if (C + 20) % 40 == 0:
        return X*C
    return 0

def process_file(filename):
    X = 1
    C = 1
    S = 0
    with open(filename) as f:
        for line in f:
            line = line.strip()
            if line == "noop":
                S += get_sig_strength(X,C)
                C += 1
                continue
            cmd, op = line.split(' ')
            op = int(op)
            if cmd == 'addx':
                S += get_sig_strength(X,C)
                C += 1
                S += get_sig_strength(X,C)
                C += 1
                X += op
    print(S)

process_file("test.txt")
process_file("input.txt")

