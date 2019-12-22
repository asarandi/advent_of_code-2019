/* advent of code 2019: day 21, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

//  .................
//  .................
//  .................
//  #####.###########

//  .................
//  .................
//  .................
//  #####...#########

//  .................
//  .................
//  .................
//  #####.#..########

//  .................
//  .................
//  .................
//  #####...####..###

//  .................
//  .................
//  .................
//  #####.##.##..####

//  .................
//  .................
//  .................
//  #####.#.#...#####
//       ABCDEFGHIJ

const solution = "NOT B J\nNOT C T\nOR T J\nAND H J\nNOT A T\nOR T J\nAND D J\nRUN\n"

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
            c := ic.program[ic.i]
            if c >= 0 && c < 128 {
                fmt.Printf(string(c))
            } else {
                fmt.Println("output", c)
            }
        }
    }
}
