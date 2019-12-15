#!/usr/bin/env python3

# advent of code 2019: day 14, part 1

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

def calc(quantity, name, res=0):
    global extras, products, produced_quantities
    if name not in products:
        return quantity
    quantity -= extras[name]
    while quantity > 0:
        quantity -= produced_quantities[name]
        for child_n, child_q in products[name].items():
            res += calc(child_q, child_n)
    extras[name] = abs(quantity)
    return res

print(calc(1, 'FUEL'))
