#!/usr/bin/env python3

from typing import List


class Monkey:
    operand: int
    dividor: int
    items: List[int]
    m1: int
    m2: int

    def _op_sum(self, v) ->  int:
        return v+self.operand

    def _op_mult(self, v) -> int:
        return v*self.operand

    def _op_sq(self, v) -> int:
        return v*v

    def test(self, v) -> bool:
        return ( v % self.dividor == 0)
           
    def __init__(self, file):
        line = file.readline().strip()
        assert line.startswith("Starting items: "), f"line: {line}"
        self.items = [int(i) for i in line[16:].split(", ")]
        line = file.readline().strip()
        assert(line.startswith("Operation: new = ")), f"line: {line}"
        if line[17:] == "old * old":
            self.op = self._op_sq
        elif line[17:23] == "old + ":
            self.operand = int(line[23:])
            self.op = self._op_sum
        elif line[17:23] == "old * ":
            self.operand = int(line[23:])
            self.op = self._op_mult
        else:
            assert(False), f"line: {line}"

        line = file.readline().strip()
        assert(line.startswith("Test: divisible by ")), f"line: {line}"

        self.dividor = int(line[19:])
        line = file.readline().strip()
        assert(line.startswith("If true: throw to monkey ")), f"line: {line}"

        self.m1 = int(line[25:])
        print(f"m1: {self.m1}") 
        line = file.readline().strip()
        assert(line.startswith("If false: throw to monkey ")), f"line: {line}"

        self.m2 = int(line[26:])
        print(f"m2: {self.m2}") 
        self.ops = 0

    def round(self, monkeys):
        while len(self.items) > 0:
            w = self.items.pop(0)
            w = self.op(w)
            w = int(w/3)
            if self.test(w):
                monkeys[self.m1].items.append(w)
            else:
                monkeys[self.m2].items.append(w)
            self.ops += 1
        

def do_round(m):
    for i in m:
        i.round(m)

def process_file(filename):
    m = []
    with open(filename) as f:
        line="\n"
        while line != "":
            line = f.readline()
            m.append(Monkey(f))
            line = f.readline()
    for rn in range(0,20):
        do_round(m)
    o = []
    for i in m:
        print(i.ops)
        o.append(i.ops)
    o.sort()
    print("res: ")
    print(o[-1]*o[-2])
import pdb
pdb.set_trace()
process_file("test.txt")
process_file("input.txt")
