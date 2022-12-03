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

def decide(opp, outcome) -> str:
    if opp == "A": # Rock
        if outcome == "X": # Lose
            return "Z" # Scisors
        if outcome == "Y": # Draw
            return "X" # Rock
        if outcome == "Z": # Win
            return "Y" # Paper
    if opp == "B": # Paper
        if outcome == "X": # Lose
            return "X" # Rock
        if outcome == "Y": # Draw
            return "Y" # Paper
        if outcome == "Z": # Win
            return "Z" # Scisors
    if opp == "C": # Scisors
        if outcome == "X": # Lose
            return "Y" # Paper
        if outcome == "Y": # Draw
            return "Z" # Scisors
        if outcome == "Z": # Win
            return "X" # Rock

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
        opp, outcome = line.strip().split(" ")
        me = decide(opp, outcome)
        score += winnerscore(opp, me)
        score += rscore(me)
        print(score)
