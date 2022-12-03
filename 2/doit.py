#!/usr/bin/env python3


def winnerscore(opp, me) -> int:
    if opp == "A": # Rock
        if me == "X": # Rock
            return 3
        if me == "Y": # Paper
            return 6
        if me == "Z":
            return 0
    if opp == "B": # Paper
        if me == "X": # Rock
            return 0
        if me == "Y": # Paper
            return 3
        if me == "Z": # Scisors
            return 6
    if opp == "C": # Scisors
        if me == "X": # Rock
            return 6
        if me == "Y": # Paper
            return 0
        if me == "Z": # Scisors
            return 3

def rscore(me) -> int:
    if me == "X":
        return 1
    if me == "Y":
        return 2
    if me == "Z":
        return 3
score = 0
with open("input.txt") as f:
    for line in f:
        opp, me = line.strip().split(" ")
        score += winnerscore(opp, me)
        score += rscore(me)
        print(score)
