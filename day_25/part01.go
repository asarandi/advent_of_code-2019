/* advent of code 2019: day 25, part 1 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

var solution = []string{
    "south",
    "south",
    "south",
    "take fixed point",
    "south",
    "take festive hat",
    "west",
    "west",
    "take jam",
    "south",
    "take easter egg",
    "north",
    "take jam",
    "east",
    "east",
    "east",
    "west",
    "north",
    "west",
    "take asterisk",
    "east",
    "north",
    "west",
    "north",
    "north",
    "take tambourine",
    "south",
    "south",
    "east",
    "north",
    "west",
    "south",
    "take antenna",
    "north",
    "west",
    "south",
    "north",
    "west",
    "take space heater",
    "west",
};

var solution_copy = []string{}

func permuteDropItems() [][]string {
    items := []string{"asterisk", "antenna", "easter egg", "space heater", "jam", "tambourine", "festive hat", "fixed point",}
    res := [][]string{}
    for i:=0; i<(1<<8); i++ {
        dropList := []string{}
        for j:=0; j<8; j++ {
            if (i>>uint(j))&1 == 1 {
                dropList = append(dropList, "drop " + items[j])
            }
        }
        res = append(res, dropList)
    }
    return res
}


func brute(ic intCode, dropList []string) {
    ic.reset()
    solution = make([]string, len(solution_copy))
    copy(solution, solution_copy)
    solution = append(solution, dropList...)
    solution = append(solution, "west", "banana")
    //fmt.Println(solution)
    for step, i, done := 0, 0, false; !done; {
        ic.run()
        if ic.state == haveOutput {
            if ic.getOutput() <= 127 {
                fmt.Printf(string(ic.getOutput()))
            } else {
//                fmt.Println("output", ic.getOutput())
            }
        } else if ic.state == wantInput {
            word := []byte(solution[step])
            var c int64
            if i >= len(word) {
//                fmt.Println("-----------------------> intCode input:", string(word))
                c = int64('\n')
                i = 0
                step++
                if step >= len(solution) {
//                    fmt.Println("done")
                    break
                }
            } else {
                c = int64(word[i])
                i++
            }
            ic.setInput(c)
        } else if ic.state == finished {
//            fmt.Println("game over")
            break
        }
    }
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
    ic.program_copy = make([]int64, len(split)*16)
    for idx, v := range split {
        i, err := strconv.ParseInt(strings.Trim(v, " "), 10, 64)
        if err != nil {
            log.Fatal(err)
        }
        ic.program[idx] = i
        ic.program_copy[idx] = i
    }
    solution_copy = make([]string, len(solution))
    copy(solution_copy, solution)
    for _, v := range permuteDropItems() {
        fmt.Println("trying", v)
//        fmt.Printf("\n\n\n\n")
        brute(ic, v)
    }
}
