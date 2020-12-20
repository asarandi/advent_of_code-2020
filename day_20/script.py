#!/usr/bin/env python3

import math
import sys

img = {}    # tileNum => tileData
config = {} # tileNum => allBorders

def borders(tile: int) -> tuple:
    global img
    data = img[tile]
    res = [0,0,0,0]
    for i in range(10):
        coords = [[0,i],[9,i],[i,0],[i,9]] # tblr
        for j, co in enumerate(coords):
            res[j] <<= 1
            y, x = co
            if data[y][x] == '#':
                res[j] |= 1
    return tuple(res)

def mirrorInt(n: int) -> int:
    x = 0
    for i in range(10):
        x <<= 1
        x |= (n >> i) & 1
    return x

def allBorders(tile: int) -> tuple:
    b = borders(tile)
    m = []
    for n in b:
        m.append(mirrorInt(n))
    return b + tuple(m)

def isVertical(i: int, j: int) -> bool:
    R = borders(i)[3]
    L = borders(j)[2]
    return R == L

def isHorizontal(i: int, j: int) -> bool:
    B = borders(i)[1]
    T = borders(j)[0]
    return B == T

def vflip(tile: int):
    global img
    data = img[tile]
    res = []
    for i in range(10):
        row = []
        for j in range(9,-1,-1):
            row.append(data[i][j])
        res.append(''.join(row))
    img[tile] = res

def hflip(tile: int):
    global img
    data = img[tile]
    res = []
    for i in range(9,-1,-1):
        row = []
        for j in range(10):
            row.append(data[i][j])
        res.append(''.join(row))
    img[tile] = res

def rotate(tile: int):
    global img
    data = img[tile]
    res = []
    for i in range(10):
        row = []
        for j in range(9,-1,-1):
            row.append(data[j][i])
        res.append(''.join(row))
    img[tile] = res

def matches(tile: int) -> tuple:
    global config
    res = set()
    for key, value in config.items():
        if key == tile:
            continue
        for v in value:
            if v in config[tile]:
                res.add(key)
                break
    return tuple(res)

def printTile(tile: int):
    for line in img[tile]:
        print(line)
    print()

#
#
#

with open("input.txt") as fp:
#with open("sample.txt") as fp:
    data = fp.read().split('\n\n')
    fp.close()

data = [x.strip() for x in data]
data = list(filter(lambda x: len(x) > 0, data))
for tile in data:
    n = int(tile[5:9])
    img[n] = tile.splitlines()[1:]
    config[n] = allBorders(n)
#    print(n, config[n])
#    printTile(n)

### part 1, xref
counts = {}
for i, iv in config.items():
    counts[i] = 0 if i not in counts else counts[i]
    for j, jv in config.items():
        if i != j and len(set(iv+jv)) != 16:
            counts[i] += 1

corners = tuple(filter(lambda x: counts[x] == 2, counts))
edges   = tuple(filter(lambda x: counts[x] == 3, counts))
inner   = tuple(filter(lambda x: counts[x] == 4, counts))

print("part 1:", math.prod(corners))

size = int(math.sqrt(len(config)))
res = [[-1] * size for _ in range(size)]

for topLeft in corners:
    h,v = False,False
    for i in range(8):
        if i % 4 == 0: hflip(topLeft)
        edge = matches(topLeft)[0]
        for j in range(8):
            if j % 4 == 0: hflip(edge)
            h = isHorizontal(topLeft,edge)
            if h: break
            rotate(edge)
        edge = matches(topLeft)[1]
        for j in range(8):
            if j % 4 == 0: vflip(edge)
            v = isVertical(topLeft,edge)
            if v: break
            rotate(edge)
        if h and v: break
        rotate(topLeft)
    if h and v:
        res[0][0] = topLeft
        break

found = {topLeft: True}
for i in range(size):
    for j in range(size):
        if res[i][j] != -1: continue
        ti, fn = [], []
        if i > 0:
            ti.append(res[i-1][j])
            fn.append(isHorizontal)
        if j > 0:
            ti.append(res[i][j-1])          
            fn.append(isVertical)
        for t in config:
            if t in found: continue
            m, ok = matches(t), True
            for prev in ti: ok &= prev in m
            if not ok: continue
            match = False
            for y in range(8):
                if y % 4 == 0: hflip(t)
                ok = True
                for x,f in enumerate(fn):
                    ok &= f(ti[x], t)
                if ok:
                    match = True
                    break
                rotate(t)
            if match:
                res[i][j] = t
                found[t] = True
                break

hashes = 0
data = ["" for _ in range(size*8)]
for i in range(size):
    for j in range(size):
        tile = img[res[i][j]]
        for y in range(8):
            for x in range(8):
                c = tile[1+y][1+x]
                hashes = hashes + 1 if c == '#' else hashes
                data[i*8+y] += c

monster = [
    "                  # ",
    "#    ##    ##    ###",
    " #  #  #  #  #  #   ",
    ]
mh, mw, mc = len(monster), len(monster[0]), ''.join(monster).count('#')

def iflip(data: []) -> []:
    res = []
    for i in range(len(data)-1,-1,-1):
        row = []
        for j in range(len(data[0])):
            row.append(data[i][j])
        res.append(''.join(row))
    return res

def irotate(data: []) -> []:
    res = []
    for i in range(len(data[0])):
        row = []
        for j in range(len(data)-1,-1,-1):
            row.append(data[j][i])
        res.append(''.join(row))
    return res

def is_monster(i,j) -> bool:
    global data, monster
    for y in range(mh):
        for x in range(mw):
            c = monster[y][x]
            if c != '#':
                continue
            if data[i+y][j+x] != '#':
                return False
    return True

count = 0
for x in range(8):
    if x % 4 == 0: data = iflip(data)
    for i in range(size*8-mh):
        for j in range(size*8-mw):
            if is_monster(i,j):
                count += 1
    if count != 0: break                
    data = irotate(data)               

print("part 2:", hashes - mc*count)

