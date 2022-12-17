#!/usr/bin/env python3

from typing import Dict
from enum import Enum
from functools import cached_property

class decision(Enum):
    Right = 1
    Wrong = 2
    NoIdea = 3


class value:
    debug = False

    def __repr__(self):
        if self.is_int:
            return f"{self.int_value}"
        else:
            return '['+ ",".join(e.__repr__() for e in self.elems) + ']'

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
        if self.debug:
            print(f"compare {self} to {right}")
        if self.is_int and right.is_int:
            if self.int_value > right.int_value:
                if self.debug:
                    print("wrong : ints")
                return decision.Wrong
            elif self.int_value < right.int_value:
                if self.debug:
                    print("right : ints")
                return decision.Right
            if self.debug:
                print("noidea : ints")
            return decision.NoIdea
        if self.is_list and right.is_list:
            for i, v in enumerate(self.elems):
                if i >= len(right.elems):
                    if self.debug:
                        print("wrong : lists - right out of elems")
                    return decision.Wrong
                cmpres = v.compare_to(right.elems[i])
                if cmpres == decision.Right or cmpres == decision.Wrong:
                    if self.debug:
                        print(f"escalating: {cmpres}")
                    return cmpres
            if len(right.elems) == len(self.elems):
                if self.debug:
                    print(f"noidea: lists")
                return decision.NoIdea
            if self.debug:
                print(f"right: lists")
            return decision.Right
        if self.is_int and right.is_list:
            if len(right.elems) < 1:
                if self.debug:
                    print("wrong int-list: empty list")
                return decision.Wrong
            return value([self.int_value]).compare_to(right)
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
        values = [value([[2]]), value([[6]])]
        for line in f:
            line=line.strip()
            if ln % 3 == 0:
                o1 = eval(line)
            if ln % 3 == 1:
                o2 = eval(line)
            if ln % 3 == 2:
                v1 = value(o1)
                v2 = value(o2)
                values.append(v1)
                values.append(v2)
                idx += 1
            ln += 1

        v1 = value(o1)
        v2 = value(o2)
        values.append(v1)
        values.append(v2)
        cmp = get_comparisons(values)
        ordered_values = []
        print("\n".join([f"{i}: " + str(cmp[i]) for i in cmp.keys()]))
        while len(values) > 0:
            for i, v1 in enumerate(values):
                v1_good = True
                for v2 in values:
                    if v2 != v1:
                        c = cmp[str(v1)][str(v2)]
                        if c == decision.Wrong:
                            v1_good = False
                            break
                if v1_good:
                    break
            if v1_good:
                ordered_values.append(values.pop(i))
            print(values)
        print(get_decoder_key(ordered_values))

def get_decoder_key(values):
    i1 = -1
    i2 = -1
    for i,v in enumerate(values):
        if str(v) == '[[2]]':
            i1 = i + 1
        if str(v) == '[[6]]':
            i2 = i + 1
    return i1*i2

def get_comparisons(values) -> Dict[str,Dict[str, decision]]:
    res = {}
    for v1 in values:
        res[str(v1)] = {}
        for v2 in values:
            res[v1.__repr__()][v2.__repr__()] = v1.compare_to(v2)
    return res


parse_file("test.txt")
print("====")
parse_file("input.txt")



