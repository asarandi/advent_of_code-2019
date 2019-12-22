#!/usr/bin/env python3

# advent of code 2019: day 22, part 1

with open("input.txt") as fp:
    data = fp.read().splitlines()
    fp.close()

def inc(d, k):
    res, i, j = [0] * len(d), 0, 0
    while j < len(d):
        res[i] = d[j]
        i = (i + k) % len(d)
        j += 1
    return res

deck = [i for i in range(10007)]

for line in data:
    if line[:3] == "cut":
        i = int(line[4:])
        deck = deck[i:] + deck[:i]
    if line[:3] == "dea":
        if line[10] == "i":
            deck = inc(deck, int(line[20:]))
        else:
            deck.reverse()
 
print(deck.index(2019))
