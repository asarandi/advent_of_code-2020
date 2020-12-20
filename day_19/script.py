#!/usr/bin/env python3

import sys
from copy import deepcopy

with open("input.txt") as fp:
#with open("sample2.txt") as fp:
    data = fp.read().split('\n\n')
    fp.close()

rules = [x.strip() for x in data[0].splitlines()]
messages = [x.strip() for x in data[1].splitlines()]
obj = {}
for i in rules:
    n = int(i.split(': ')[0])
    s = i.split(': ')[1]
    if '|' in s:
        obj[n] = {}
        for zu in s.split('|'):
            t = tuple([int(x) for x in zu.strip().split(' ')])
            obj[n][t] = True
    else:
        if s[0] == '"':
            obj[n] = s[1]
        else:
            nums = tuple([int(x) for x in s.split(' ')])
            obj[n] = nums

def search(needle):
    global obj, res
    curr = [[0]]
    pref = ['']
    while len(curr) > 0:
        nodeN = curr.pop(0)
        s = pref.pop(0)
        lens = len(s)
        if not needle[:lens] == s:
            continue
        while len(nodeN) > 0:
            fork = False
            n = nodeN.pop(0)
            if isinstance(obj[n], str):
                s += obj[n]
                continue
            elif isinstance(obj[n], tuple):
                nodeN = list(obj[n]) + nodeN
                continue                    
            elif isinstance(obj[n], dict):
                fork = True
                for k in obj[n]:
                    curr.append(list(k)+deepcopy(nodeN))
                    pref.append(deepcopy(s))
                break
        if fork:
            continue
        if s == needle:
            return True
    return False            

res = 0
for m in messages:
    res = res + 1 if search(m) else res
print("part 1:", res)


obj[8] = {tuple([42]): True, tuple([42, 8]): True}
obj[11] = {tuple([42, 31]): True, tuple([42,11,31]): True}
res = 0
for m in messages:
    res = res + 1 if search(m) else res
print("part 2:", res)
