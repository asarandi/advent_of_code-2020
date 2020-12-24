#!/usr/bin/env python3

import sys
from copy import deepcopy

with open("input.txt") as fp:
#with open("sample.txt") as fp:
    data = fp.read().splitlines()
    fp.close()

data = [x.strip() for x in data]
data = list(filter(lambda x: len(x) > 0, data))


steps = {  "e":tuple([0,2]),
        "se":tuple([1,1]),
        "sw":tuple([1,-1]),
         "w":tuple([0,-2]),
        "nw":tuple([-1,-1]),
        "ne":tuple([-1,1]),
        }

tiles = {}
for line in data:
    i,y,x = 0,0,0
    while i < len(line):
        a = line[i:i+2]
        b = line[i:i+1]
        if a in steps:
            p,q = steps[a]
            y,x = y+p, x+q
            i += 2
        elif b in steps:
            p,q = steps[b]
            y,x = y+p, x+q
            i += 1
        else:
            print("wtf")
    tu = tuple([y,x])
    if tu not in tiles:
        tiles[tu] = 0
    tiles[tu] ^= 1

res = 0
for k,v in tiles.items():
    if v == 1:
        res += 1
print("part 1", res)           

for _ in range(100):
    nextgen = {}
    for k,v in tiles.items():
        if v != 1:
            continue
        y,x = k
        for step in steps.values():
            a,b = step
            a,b = a+y,b+x
            tu = tuple([a,b])
            if tu not in tiles:
                nextgen[tu] = 0
    for k,v in nextgen.items():
        tiles[k] = v
    nextgen = {}        

    for k,v in tiles.items():
        y,x = k
        ct = [0,0]
        for step in steps.values():
            a,b = step
            a,b = a+y,b+x
            tu = tuple([a,b])
            val = 0
            if tu in tiles:
                val = tiles[tu]

            ct[val] += 1
        if v == 1: #black
            if (ct[1] == 0) or (ct[1] > 2):
                v ^= 1
        else:
            if (ct[1] == 2):
                v ^= 1
        nextgen[k] = v                
    tiles = nextgen        

res = 0
for k,v in tiles.items():
    if v == 1:
        res += 1
print("part 2", res)           
