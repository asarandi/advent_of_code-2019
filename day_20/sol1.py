#!/usr/bin/env python3

with open("input.txt") as fp:
    data = fp.read().splitlines()
    fp.close()

outer, inner, pairs, grid = {}, {}, {} ,{}
start, finish = None, None
for i in range(len(data)):
    for j in range(len(data[i])):
        grid[(i,j)] = data[i][j]
        if 'A'<=data[i][j]<='Z':
            if i < 1 or i > len(data) - 2: continue
            if j < 1 or j > len(data[i]) - 2: continue
            pt, name = None, None
            if data[i+1][j] == '.': pt = i+1, j; name = data[i-1][j] + data[i][j]
            if data[i-1][j] == '.': pt = i-1, j; name = data[i][j] + data[i+1][j]
            if data[i][j+1] == '.': pt = i, j+1; name = data[i][j-1] + data[i][j]
            if data[i][j-1] == '.': pt = i, j-1; name = data[i][j] + data[i][j+1]
            if not pt: continue
            if i == 1 or i == len(data) - 2 or j == 1 or j == len(data[i]) - 2:
                outer[pt] = name
                if name == 'AA': start = pt
                if name == 'ZZ': finish = pt
            else:
                inner[pt] = name

for ok,ov in outer.items():
    for ik,iv in inner.items():
        if ov == iv:
            pairs[ok] = ik
            pairs[ik] = ok

queue = [(start, 0)]
visited = set()
while queue:
    pos, g = queue.pop(0)
    if pos in visited:
        continue
    if pos == finish:
        print("part 1:", g)
        break
    visited.add(pos)
    if pos in pairs:
        queue.append((pairs[pos], g + 1))
    i,j = pos
    moves = [(i+1,j), (i-1,j), (i, j+1), (i, j-1)]
    for m in moves:
        if m not in grid:
            continue
        if grid[m] != '.':
            continue
        queue.append((m, g + 1))
