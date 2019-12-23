/* advent of code 2019: day 23, part 01 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    trim := strings.Trim(string(content), " \t\n\r\v\f")
    split := strings.Split(trim, ",")
    program := make([]int64, len(split)*16)
    for idx, v := range split {
        i, err := strconv.ParseInt(strings.Trim(v, " "), 10, 64)
        if err != nil {
            log.Fatal(err)
        }
        program[idx] = i
    }
    var machines []intCode
    var inputs [50][]int64
    for i := 0; i < 50; i++ {
        ic := intCode{}
        ic.program, ic.program_copy = make([]int64, len(program)), make([]int64, len(program))
        copy(ic.program, program)
        copy(ic.program_copy, program)
        machines = append(machines, ic)
        machines[i].run()                             //init
        machines[i].program[machines[i].i] = int64(i) //set machine numbers
        machines[i].run()                             //continue
        inputs[i] = make([]int64, 0)
    }
    done := false
    for ; !done; {
        done = true
        for i := 0; i < 50; i++ {
            done = done && machines[i].state == finished
            if machines[i].state == finished {
                continue
            }
            if machines[i].state == haveOutput {
                address := machines[i].getOutput()
                machines[i].run()
                if machines[i].state != haveOutput {
                    fmt.Println("error: expecting 'haveOutput' state for byte X")
                }
                byteX := machines[i].getOutput()
                machines[i].run()
                if machines[i].state != haveOutput {
                    fmt.Println("error: expecting 'haveOutput' state for byte Y")
                }
                byteY := machines[i].getOutput()
                if address >= int64(len(machines)) {
                    if address == 255 {
                        fmt.Println("part 1:", byteY)
                        done = true
                        break
                    } else {
                        fmt.Println("out of bounds", address, byteX, byteY)
                    }
                } else {
                    inputs[int(address)] = append(inputs[int(address)], byteX)
                    inputs[int(address)] = append(inputs[int(address)], byteY)
                }
            } else if machines[i].state == wantInput {
                if len(inputs[i]) == 0 {
                    machines[i].setInput(-1)
                } else if len(inputs[i]) >= 2 {
                    machines[i].setInput(inputs[i][0])
                    machines[i].run()
                    if machines[i].state != wantInput {
                        fmt.Println("error: expecting 'wantInput' state for byte Y")
                    }
                    machines[i].setInput(inputs[i][1])
                    inputs[i] = inputs[i][2:]
                }
            }
            if machines[i].state != finished {
                machines[i].run()
            }
        }
    }
}
