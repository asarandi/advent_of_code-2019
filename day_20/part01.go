/* advent of code 2019: day 20, part 1 */
package main

import (
    "container/list"
    "fmt"
    "io/ioutil"
    "strings"
)

var grid map[point]byte
var inner map[point]string
var outer map[point]string
var pairs map[point]point
var start point
var finish point

type point struct {
    y, x int
}

func isUpper(b byte) bool {
    return b >= 'A' && b <= 'Z'
}

func (p point) up() point    { return point{p.y - 1, p.x} }
func (p point) right() point { return point{p.y, p.x + 1} }
func (p point) down() point  { return point{p.y + 1, p.x} }
func (p point) left() point  { return point{p.y, p.x - 1} }

func parse(f string) {
    content, err := ioutil.ReadFile(f)
    if err != nil {
        panic(err)
    }
    split := strings.Split(strings.Trim(string(content), "\n"), "\n")
    grid = make(map[point]byte)
    inner = make(map[point]string)
    outer = make(map[point]string)
    pairs = make(map[point]point)
    for i := 0; i < len(split); i++ {
        for j := 0; j < len(split[i]); j++ {
            p := point{i, j}
            c := split[i][j]
            grid[p] = c
        }
    }
    for k, v := range grid {
        if !isUpper(v) {
            continue
        }
        p, name := point{0, 0}, ""
        u, r, d, l := grid[k.up()], grid[k.right()], grid[k.down()], grid[k.left()]
        if isUpper(u) && d == '.' {
            p = k.down()
            name = string(u) + string(v)
        }
        if isUpper(d) && u == '.' {
            p = k.up()
            name = string(v) + string(d)
        }
        if isUpper(l) && r == '.' {
            p = k.right()
            name = string(l) + string(v)
        }
        if isUpper(r) && l == '.' {
            p = k.left()
            name = string(v) + string(r)
        }
        if p.y == 0 && p.x == 0 {
            continue
        }
        if p.y == 2 || p.y == len(split)-3 || p.x == 2 || p.x == len(split[0])-3 {
            outer[p] = name
        } else {
            inner[p] = name
        }
        if name == "AA" {
            start = p
        }
        if name == "ZZ" {
            finish = p
        }
    }
    for ok, ov := range outer {
        for ik, iv := range inner {
            if ov == iv {
                pairs[ik] = ok
                pairs[ok] = ik
            }
        }
    }
}

type node struct {
    pos  point
    dist int
}

func main() {
    parse("input.txt")
    queue := list.New()
    queue.PushBack(node{start, 0})
    seen := make(map[point]bool)
    for ; queue.Len() > 0; {
        n := queue.Remove(queue.Front()).(node)
        pos, dist := n.pos, n.dist
        if _, ok := seen[pos]; ok {
            continue
        }
        seen[pos] = true
        if pos == finish {
            fmt.Println("part 1:", dist)
            break
        }
        if portal, ok := pairs[pos]; ok {
            queue.PushBack(node{portal, dist + 1})
        }
        moves := []point{pos.up(), pos.down(), pos.left(), pos.right()}
        for _, move := range moves {
            if grid[move] == '.' {
                queue.PushBack(node{move, dist + 1})
            }
        }
    }
}
