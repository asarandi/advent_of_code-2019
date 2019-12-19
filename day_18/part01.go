package main

import (
    "container/list"
    "fmt"
    "io/ioutil"
    "strings"
)

type point struct {
    y, x int
}

var allKeys [26]bool
var area map[point]byte

type state struct {
    pos  point
    keys [26]bool
}

func (s state) children() []state {
    res := make([]state, 0)
    moves := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for i := 0; i < 4; i++ {
        childPos := point{s.pos.y + moves[i].y, s.pos.x + moves[i].x}
        c, ok := area[childPos]
        if !ok {
            continue
        }
        if c == '#' {
            continue
        }
        if c >= 'A' && c <= 'Z' {
            if !s.keys[int(c-'A')] {
                continue
            }
        }
        child := s
        child.pos = childPos
        if c >= 'a' && c <= 'z' {
            child.keys[int(c-'a')] = true
        }
        res = append(res, child)
    }
    return res
}

func (s state) isGoal() bool {
    res := true
    for i := 0; i < 26; i++ {
        res = res && s.keys[i] == allKeys[i]
    }
    return res
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    split := strings.Split(strings.Trim(string(content), " \n\t\r\f\v"), "\n")
    area = make(map[point]byte)
    start := state{}
    for i := 0; i < len(split); i++ {
        for j := 0; j < len(split[0]); j++ {
            c := split[i][j]
            p := point{i, j}
            area[p] = c
            if c >= 'a' && c <= 'z' {
                allKeys[int(c-'a')] = true
            }
            if c == '@' {
                start.pos = p
            }
        }
    }
    distances := make(map[state]int)
    distances[start] = 0
    queue := list.New()
    queue.PushBack(start)
    for ; queue.Len() > 0; {
        node := queue.Remove(queue.Front()).(state)
        if node.isGoal() {
            fmt.Println("part 1:", distances[node])
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
}
