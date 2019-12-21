/* advent of code 2019: day 18, part 02: does not work properly */
package main

import (
    "container/list"
    "fmt"
    "io/ioutil"
    "strings"
)

var center point
var robots [4]robot
var grid map[point]byte

type point struct {
    y, x int
}

type robot struct {
    pos           point
    lockedDoors   [26]bool
    neededKeys    [26]bool
    collectedKeys [26]bool
    distance      int
}

func (r robot) isGoal() bool {
    for i := 0; i < 26; i++ {
        if r.neededKeys[i] != r.collectedKeys[i] {
            return false
        }
    }
    return true
}

func (r robot) children() []robot {
    res := make([]robot, 0)
    moves := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for i := 0; i < 4; i++ {
        childPos := r.pos.add(moves[i])
        c, ok := grid[childPos]
        if !ok {
            continue
        }
        if c == '#' {
            continue
        }
        if c >= 'A' && c <= 'Z' {
            if r.lockedDoors[int(c-'A')] && !r.collectedKeys[int(c-'A')] {
                continue
            }
        }
        child := r
        child.pos = childPos
        if c >= 'a' && c <= 'z' {
//          fmt.Println("robot", child.n, "collects key", string(c))
            child.collectedKeys[int(c-'a')] = true
        }
        res = append(res, child)
    }
    return res
}

func (p point) add(q point) point {
    return point{p.y + q.y, p.x + q.x}
}

/* return values 0..3 */
func (p point) getQuadrant() int {
    res := 0
    if p.y > center.y {
        res += 2
    }
    if p.x > center.x {
        res += 1
    }
    return res
}

func prepare(f string) {
    content, err := ioutil.ReadFile(f)
    if err != nil {
        panic(err)
    }
    split := strings.Split(strings.Trim(string(content), " \n\t\r\f\v"), "\n")
    center = point{(len(split) - 1) / 2, (len(split[0]) - 1) / 2}
    grid = make(map[point]byte)
    for i := range split {
        for j := range split[i] {
            p := point{i, j}
            q := p.getQuadrant()
            c := split[i][j]
            grid[p] = c
            if c >= 'a' && c <= 'z' {
                robots[q].neededKeys[int(c-'a')] = true
            }
        }
    }
    grid[center] = '#'
    grid[point{center.y - 1, center.x}] = '#'
    grid[point{center.y + 1, center.x}] = '#'
    grid[point{center.y, center.x - 1}] = '#'
    grid[point{center.y, center.x + 1}] = '#'
    grid[point{center.y - 1, center.x - 1}] = '@'
    grid[point{center.y - 1, center.x + 1}] = '@'
    grid[point{center.y + 1, center.x + 1}] = '@'
    grid[point{center.y + 1, center.x - 1}] = '@'
    for k, v := range grid {
        q := k.getQuadrant()
        if v >= 'A' && v <= 'Z' && robots[q].neededKeys[int(v-'A')] {
            robots[q].lockedDoors[int(v-'A')] = true
        }
        if v == '@' {
            robots[q].pos = k
        }
    }
}

//func printGrid() {
//    for i:=0; i<center.y*2+1; i++ {
//        for j:=0; j<center.x*2+1; j++ {
//            fmt.Printf(string(grid[point{i,j}]))
//        }
//        fmt.Println()
//    }
//}

func main() {
    prepare("input.txt")
//    printGrid()
    for i := 0; i < len(robots); i++ {
        distances := make(map[robot]int)
        distances[robots[i]] = 0
        queue := list.New()
        queue.PushBack(robots[i])
//        flag := false
        for ; queue.Len() > 0; {
            node := queue.Remove(queue.Front()).(robot)
            if node.isGoal() {
//                flag = true
                robots[i] = node
                robots[i].distance = distances[node]
                break
            }
            for _, child := range node.children() {
                if _, ok := distances[child]; ok {
                    continue
                }
                distances[child] = distances[node] + 1
                queue.PushBack(child)
            }
        }
        //fmt.Println("success?", flag)
        //
        //fmt.Println("        robot", i)
        //fmt.Println("   neededKeys", robots[i].neededKeys)
        //fmt.Println("collectedKeys", robots[i].collectedKeys)
        //fmt.Println("  lockedDoors", robots[i].lockedDoors)
        //fmt.Println("     distance", robots[i].distance)

    }
    res := 0
    for i := 0; i < len(robots); i++ {
        res += robots[i].distance
    }
    fmt.Println("part 2:", res)
}
