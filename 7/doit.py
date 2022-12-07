#!/usr/bin/env python3

from typing import Dict, Any

class Tree:
    dirs : Dict[str,Any]
    files: Dict[str,int]
    size : int
    parent : Any

    def __init__(self):
        self.dirs = {}
        self.files = {}

    def get_size(self) -> int:
        s = 0
        for fsize in self.files.values():
            s += fsize
        for t in self.dirs.values():
            s += t.get_size()
        self.size = s
        return s

    def traverse(self) -> int:
        res = 0
        if self.size <= 100000:
            res += self.size
        for t in self.dirs.values():
            res += t.traverse()
        return res

    def minimise(self, required, known) -> int:
        if self.size < required:
            return known
        if self.size == required:
            return self.size
        if self.size > required:
            candidate = known
            for t in self.dirs.values():
                candidate = t.minimise(required, candidate)
            if self.size < candidate:
                candidate = self.size
            return candidate


def read_tree(filename):
    root = Tree()
    current = root
    with open(filename) as f:
        for line in f:
            line = line.strip()
            if line == "$ cd /":
                current = root
                continue
            if line == "$ cd ..":
                current = current.parent
                continue
            if line == "$ ls":
                continue
            if line[0:5] == "$ cd ":
                targetname=line[5:]
                current = current.dirs[targetname]
                continue
            p1, p2 = line.split(" ")
            if p1 == "dir":
                t = Tree()
                t.parent = current
                current.dirs[p2] = t
                continue
            current.files[p2] = int(p1)
    print(root.get_size())
    print(root.traverse())
    available = 70000000 - root.size
    required = 30000000 - available
    print(root.minimise(required, root.size))


read_tree("test.txt")
read_tree("input.txt")
