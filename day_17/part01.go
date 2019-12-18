/* advent of code 2019: day 17, part 01 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

type point struct {
    y, x int
}

type grid map[point]byte

func (s grid) isIntersection(p point) bool {
    res := s[p] == '#'
    res = res && s[point{p.y + 1, p.x}] == '#'
    res = res && s[point{p.y - 1, p.x}] == '#'
    res = res && s[point{p.y, p.x + 1}] == '#'
    res = res && s[point{p.y, p.x - 1}] == '#'
    return res
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    trim := strings.Trim(string(content), " \t\n\r\v\f")
    split := strings.Split(trim, ",")
    ic := intCode{}
    ic.program = make([]int64, len(split)*16)
    for idx, v := range split {
        i, err := strconv.ParseInt(strings.Trim(v, " "), 10, 64)
        if err != nil {
            log.Fatal(err)
        }
        ic.program[idx] = i
    }
    i, j := 0, 0
    screen := make(grid)
    for ; ic.state != finished; {
        ic.run()
        if ic.state == haveOutput {
            c := byte(ic.program[ic.i])
            if c == '\n' {
                i++
                j = 0
            } else {
                screen[point{i, j}] = c
                j++
            }
        }
    }
    res := 0
    for p := range screen {
        if screen.isIntersection(p) {
            res += p.y * p.x
        }
    }
    fmt.Println("part 1:", res)
}
