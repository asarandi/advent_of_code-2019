/* advent of code 2019: day 18, part 02 */
package main

import (
    "container/heap"
    "fmt"
    "io/ioutil"
    "strings"
)

type point struct {
    y, x int
}

var allKeys [26]bool
var numKeys int
var area map[point]byte

type state struct {
    pos  [4]point
    keys [26]bool
}

type Item struct {
    node     *state
    parent   *state
    distance int
    priority int
    index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Item)
    item.index = len(*pq)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    item := (*pq)[len(*pq)-1]
    *pq = (*pq)[0 : len(*pq)-1]
    return item
}

func (s state) heuristic() int {
    count := 0
    for i := 0; i < len(s.keys); i++ {
        if s.keys[i] {
            count++
        }
    }
    return numKeys - count
}

func (s state) children() []state {
    res := make([]state, 0)
    moves := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for r := 0; r < 4; r++ {
        for i := 0; i < 4; i++ {
            childPos := point{s.pos[r].y + moves[i].y, s.pos[r].x + moves[i].x}
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
            child.pos[r] = childPos
            if c >= 'a' && c <= 'z' {
                child.keys[int(c-'a')] = true
            }
            res = append(res, child)
        }
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
    content, err := ioutil.ReadFile("sample_24.txt")
    if err != nil {
        panic(err)
    }
    split := strings.Split(strings.Trim(string(content), " \n\t\r\f\v"), "\n")
    area = make(map[point]byte)
    r := 0
    start := state{}
    for i := 0; i < len(split); i++ {
        for j := 0; j < len(split[0]); j++ {
            c := split[i][j]
            p := point{i, j}
            area[p] = c
            if c >= 'a' && c <= 'z' {
                allKeys[int(c-'a')] = true
                numKeys++
            }
            if c == '@' {
                start.pos[r] = p
                r++
            }
        }
    }
    closed := make(map[state]bool)
    distances := make(map[state]int)
    distances[start] = 0
    pq := make(PriorityQueue, 0)
    item := Item{&start, nil, 0, start.heuristic(), 0}
    heap.Push(&pq, &item)
    for ; pq.Len() > 0; {
        item := heap.Pop(&pq).(*Item)
        node := *item.node
        //fmt.Println(node, "heuristic", node.heuristic(), "dist", distances[node])
        if node.isGoal() {
            fmt.Println("part 2:", distances[node])
            break
        }
        if _, ok := closed[node]; ok {
            continue
        }
        closed[node] = true
        for _, child := range node.children() {
            if _, ok := closed[child]; ok {
                //fmt.Println("closed", child)
                continue
            }
            distances[child] = distances[node] + 1
            f := distances[child] + child.heuristic()
            chi := Item{&child, &node, distances[child], f, 0}
            pq.Push(&chi)
            fmt.Println("parent", node)
            fmt.Println(" child", child)
        }
    }
}
