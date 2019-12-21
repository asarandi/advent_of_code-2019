/* advent of code 2019: day 21, part 01 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

//const solution = "NOT A J\nNOT B T\nAND T J\nNOT C T\nAND T J\nAND D J\nWALK\n"

//.....0100..
//.....010...
//.....000...     // OR A T, OR B T, OR C T, NOT T J,
//.....0.....     // NOT A J


const solution = "NOT A J\nNOT C T\nWALK\n"

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
    for i:=0; ic.state != finished; {
        ic.run()
        if ic.state == wantInput {
            ic.program[ic.i] = int64(solution[i])
            i++
        }
        if ic.state == haveOutput {
            fmt.Printf(string(ic.program[ic.i]))
        }
    }
}
