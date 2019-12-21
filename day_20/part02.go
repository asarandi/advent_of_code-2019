/* advent of code 2019: day 20, part 2 */
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

func isupper(b byte) bool {
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
        if !isupper(v) {
            continue
        }
        p, name := point{0, 0}, ""
        u, r, d, l := grid[k.up()], grid[k.right()], grid[k.down()], grid[k.left()]
        if isupper(u) && d == '.' {
            p = k.down()
            name = string(u) + string(v)
        }
        if isupper(d) && u == '.' {
            p = k.up()
            name = string(v) + string(d)
        }
        if isupper(l) && r == '.' {
            p = k.right()
            name = string(l) + string(v)
        }
        if isupper(r) && l == '.' {
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

type state struct {
    pos   point
    level int
}



type dist struct {
    s state
    d int
}

func main() {
    parse("input.txt")
    queue := list.New()
    queue.PushBack(state{start, 0})
    make
    dist[state{start, 0}] = 0
    seen := make(map[state]bool)
    for ; queue.Len() > 0; {
        node := queue.Remove(queue.Front()).(state)
        if _, ok := seen[node]; ok {
            continue
        }
        seen[node] = true
        if _, ok := inner[node.pos]; ok {
            queue.PushBack(state{pairs[node.pos], node.level + 1, node.d + 1})
        }
        if _, ok := outer[node.pos]; ok {
            if node.level == 0 {
                if node.pos == finish {
                    fmt.Println("part 2:", node.d)
                    break
                }
            } else {
                if node.pos != start && node.pos != finish {
                    queue.PushBack(state{pairs[node.pos], node.level - 1, node.d + 1})
                }
            }
        }
        moves := []point{node.pos.up(), node.pos.down(), node.pos.left(), node.pos.right()}
        for _, move := range moves {
            if grid[move] == '.' {
                queue.PushBack(state{move, node.level, node.d + 1})
            }
        }
    }
}
