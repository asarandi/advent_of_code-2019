#!/usr/bin/env python3

# advent of code 2019: day 14, part 2

import math

with open("input.txt") as fp:
    data = fp.read().splitlines()
    fp.close

def get_quantity_name(s):
    quantity, name = s.split(" ")
    return int(quantity), name

extras = {}
products = {}
produced_quantities = {}
for line in data:
    children, parent = line.split(" => ")
    parent_q, parent_n = get_quantity_name(parent)
    extras[parent_n] = 0
    produced_quantities[parent_n] = parent_q
    ingridients = {}
    for child in children.split(", "):
        child_q, child_n = get_quantity_name(child)
        extras[child_n] = 0
        ingridients[child_n] = child_q
    products[parent_n] = ingridients

def clear_extras():
    global extras
    for k in extras.keys():
        extras[k] = 0

def bfs(quantity, name, res=0):
    clear_extras()
    queue = []
    queue.append(tuple([quantity, name]))
    while queue:
        quantity, name = queue.pop(0)
        if name not in products:
            res += quantity
            continue
        quantity -= extras[name]
        if quantity <= 0:
            extras[name] = abs(quantity)
            continue
        k = math.ceil(quantity / produced_quantities[name])
        extras[name] = k * produced_quantities[name] - quantity
        for child_n, child_q in products[name].items():
            queue.append(tuple([k * child_q, child_n]))
    return res


cost = bfs(1, 'FUEL')
goal = 1000000000000
lower = 1
upper = 1<<32
middle = -1
while lower <= upper:
    mid = lower + upper >> 1
    ore = bfs(mid, 'FUEL')
    if ore + cost < goal:
        lower = mid + 1
    elif ore > goal:
        upper = mid - 1
    else:
        break

print("part 1:", cost)
print("part 2:", mid)
