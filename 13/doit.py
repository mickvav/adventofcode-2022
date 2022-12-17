#!/usr/bin/env python3

from enum import Enum

class decision(Enum):
    Right = 1
    Wrong = 2
    NoIdea = 3


class value:

    def __repr__(self):
        if self.is_int:
            return str(self.int_value)
        else:
            return str(self.elems)

    def __init__(self, o):
        if type(o) is int:
            self.is_list = False
            self.is_int = True
            self.int_value = o
        if type(o) is list:
            self.elems = []
            self.is_list = True
            self.is_int = False
            for e in o:
                self.elems.append(value(e))

    def compare_to(self, right) -> decision:
        print(f"- Compare {self} vs {right}")
        if self.is_int and right.is_int:
            if self.int_value > right.int_value:
                return decision.Wrong
            elif self.int_value < right.int_value:
                return decision.Right
            return decision.NoIdea
        if self.is_list and right.is_list:
            for i, v in enumerate(self.elems):
                if i >= len(right.elems):
                    return decision.Wrong
                cmpres = v.compare_to(right.elems[i])
                if cmpres == decision.Right or cmpres == decision.Wrong:
                    return cmpres
            if len(right.elems) == len(self.elems):
                return decision.NoIdea
            return decision.Right
        if self.is_int and right.is_list:
            if len(right.elems) < 1:
                return decision.Wrong
            r = right.elems[0]
            return value(self.int_value).compare_to(r)
        if self.is_list and right.is_int:
            return self.compare_to(value([right.int_value]))
            if len(self.elems) > 1:
                return False
            if len(self.elems) == 0:
                return True
            if self.elems[0].is_int and self.elems[0].int_value > right.int_value:
                return False
            return self.elems[0].compare_to(value(right.int_value))

def parse_file(filename):
    with open(filename) as f:
        ln = 0
        idx = 1
        idxsum = 0
        for line in f:
            line=line.strip()
            if ln % 3 == 0:
                o1 = eval(line)
            if ln % 3 == 1:
                o2 = eval(line)
            if ln % 3 == 2:
                v1 = value(o1)
                v2 = value(o2)
                if v1.compare_to(v2) == decision.Right:
                    idxsum += idx
                    print(f"Right: {idx}")
                idx += 1
            ln += 1

        v1 = value(o1)
        v2 = value(o2)
        if v1.compare_to(v2) == decision.Right:
            idxsum += idx
        print(idxsum)

parse_file("test.txt")
print("====")
parse_file("input.txt")



