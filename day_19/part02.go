/* advent of code 2019: day 19, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    _ "math/rand"
    "strconv"
    "strings"
)

func isValid(ic *intCode, y, x int) bool {
    haveY, res := true, true
    for ; ic.state != finished; {
        ic.run()
        if ic.state == wantInput {
            if haveY {
                ic.program[ic.i] = int64(y)
                haveY = false
            } else {
                ic.program[ic.i] = int64(x)
                haveY = true
            }
        }
        if ic.state == haveOutput {
            res = ic.program[ic.i] == 1
            break
        }
    }
    ic.reset()
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
    ic.program_copy = make([]int64, len(ic.program))
    copy(ic.program_copy, ic.program)
    i, j, size := 100, 100, 100
    for {
        if !isValid(&ic, i, j) {
            j++
            continue
        }
        if !isValid(&ic, i, j+size-1) {
            i++
            continue
        }
        if !isValid(&ic, i+size-1, j) {
            j++
            continue
        }
        break
    }
    fmt.Println("part 2:", i*10000+j)
}
