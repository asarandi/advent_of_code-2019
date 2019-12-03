#!/usr/bin/env python

# advent of code 2019: day 03, part 01

import sys

def get_input(f):
    data = None
    with open(f) as fp:
        data = fp.read().splitlines()
        fp.close()
    return data        

def get_xy(start,s):
    d = s[0]
    i = int(s[1:])
    y0,x0 = start
    if d == 'R':
        return (y0,x0+i)        
    if d == 'U':
        return (y0-i,x0)
    if d == 'D':
        return (y0+i,x0)
    if d == 'L':
        return (y0,x0-i)
    return None        

def trace(wire):
    res = []
    src = (0,0)
    for s in wire.split(','):
        dst = get_xy(src, s)
        res.append((src,dst))
        src = dst
    return res        

def is_vertical(line):
    src, dst = line
    return True if src[1] == dst[1] else False      # same x coordinate for both

#  
#    0123456789a
#  0 ...........
#  1 .+-----+...
#  2 .|.....|...
#  3 .|..+--X-+.
#  4 .|..|..|.|.
#  5 .|.-X--+.|.  5,3 - 5,7 horizontal
#  6 .|..|....|.  3,4 - 6,4 vertical
#  7 .|.......|.
#  8 .o-------+.
#  9 ...........
#  

def fix(line):
    src, dst = line
    sy, sx = src
    dy, dx = dst
    if is_vertical(line):
        return (dst, src) if sy > dy else (src, dst)
    else:
        return (dst, src) if sx > dx else (src, dst)

def intersections(wire1,wire2):
    res = []
    for line1 in wire1:
        for line2 in wire2:
            if is_vertical(line1) == is_vertical(line2):
                continue
            line1 = fix(line1)
            line2 = fix(line2)
            src0, dst0 = line1
            src1, dst1 = line2
            if is_vertical(line1):
                if not ((src0[0] <= src1[0]) and (dst0[0] >= src1[0])):
                    continue
                if not ((src1[1] <= src0[1]) and (dst1[1] >= src0[1])):
                    continue
                res.append((src1[0], src0[1]))
            else:
                if not ((src1[0] <= src0[0]) and (dst1[0] >= src0[0])):
                    continue
                if not ((src0[1] <= src1[1]) and (dst0[1] >= src1[1])):
                    continue
                res.append((src0[0], src1[1]))
    if (0,0) in res:                
        res.remove((0,0))                
    return res         

def distances(intersections):
    res = []
    for i in intersections:
        y,x = i
        res.append(abs(y) + abs(x))
    return res        


if __name__ == '__main__':
    if len(sys.argv) < 2:
        sys.exit("please provide file name")
    data = get_input(sys.argv[1])
    wire1 = trace(data[0])
    wire2 = trace(data[1])
    i = intersections(wire1, wire2)
    dist = distances(i)
    print(min(dist))


