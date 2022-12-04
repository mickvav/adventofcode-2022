#!/usr/bin/env python3
m = 0
s = 0
a = []
with open("input.txt") as f:
    for line in f:
        if line.strip() == "":
            if s>m:
                m = s
            a.append(s)
            s=0
            print(m)
        else:
            s+=int(line.strip())
if s>m:
    m = s
    a.append(s)
 
print(sum(sorted(a)[-3:]))
