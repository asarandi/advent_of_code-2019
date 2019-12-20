/* advent of code 2019: day 19, part 01 */
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
    ic.program_copy = make([]int64, len(ic.program))
    copy(ic.program_copy, ic.program)

    res, i, input, size := 0, 0, 0, 50
    haveY := true
    for ; ic.state != finished; {
        ic.run()
        if ic.state == wantInput {
            if haveY {
                input = i / size
                haveY = false
            } else {
                input = i % size
                haveY = true
            }
            ic.program[ic.i] = int64(input)
        }
        if ic.state == haveOutput {
            i++
            if ic.program[ic.i] == 1 {
                res++
            }
        }
        if ic.state == finished && i < size*size {
            ic.reset()
        }
    }
    fmt.Println("part 1:", res)
}
