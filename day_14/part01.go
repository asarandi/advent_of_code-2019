package main

/* advent of code 2019: day 14, part 1 & part 2 */

import (
    "fmt"
    "io/ioutil"
    "math"
    "strconv"
    "strings"
)

var productIngredients map[string]map[string]int
var producedQuantities map[string]int
var extras map[string]int

func getQuantityName(s string) (int, string) {
    split := strings.Split(s, " ")
    quantity, _ := strconv.Atoi(split[0])
    return quantity, split[1]
}

func clearExtras() {
    for k := range extras {
        extras[k] = 0
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func dfs(quantity int, name string) int {
    if _, ok := productIngredients[name]; !ok {
        return quantity
    }
    quantity -= extras[name]
    res, k := 0, 0
    if quantity > 0 {
        k = int(math.Ceil(float64(quantity) / float64(producedQuantities[name])))
        quantity -= k * producedQuantities[name]
    }
    extras[name] = abs(quantity)
    for childN, childQ := range productIngredients[name] {
        res += dfs(childQ*k, childN)
    }
    return res
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    trim := strings.Trim(string(content), " \t\r\n\f\v")
    split := strings.Split(trim, "\n")
    productIngredients = make(map[string]map[string]int)
    producedQuantities = make(map[string]int)
    extras = make(map[string]int)
    for _, line := range split {
        reaction := strings.Split(line, " => ")
        parentQ, parentN := getQuantityName(reaction[1])
        producedQuantities[parentN] = parentQ
        extras[parentN] = 0
        ingredients := make(map[string]int)
        for _, child := range strings.Split(reaction[0], ", ") {
            childQ, childN := getQuantityName(child)
            ingredients[childN] = childQ
            extras[childN] = 0
        }
        productIngredients[parentN] = ingredients
    }
    cost := dfs(1, "FUEL")
    fmt.Println("part 1:", cost)
    lo, hi, mid, ore, goal := 1, 1<<32, 0, 0, 1000000000000
    for ; lo <= hi; {
        mid = (lo + hi) >> 1
        clearExtras()
        ore = dfs(mid, "FUEL")
        if ore+cost < goal {
            lo = mid + 1
        } else if ore > goal {
            hi = mid - 1
        } else {
            break
        }
    }
    fmt.Println("part 2:", mid)
}
