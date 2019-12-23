/* advent of code 2019: day 23, part 1 and 2 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

const numMachines = 50
const NAT = 255

var machines [numMachines]intCode
var inputs [numMachines][]int64

func (ic *intCode) getPackets() (int64, int64, int64) {
    addr := ic.getOutput()
    ic.run()
    packetX := ic.getOutput()
    ic.run()
    packetY := ic.getOutput()
    return addr, packetX, packetY
}

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
    for i := 0; i < numMachines; i++ {
        machines[i] = intCode{make([]int64, len(program)), nil, 0, 0, 0, 0, 0, 0, 0}
        copy(machines[i].program, program)
        inputs[i] = make([]int64, 1)
        inputs[i][0] = int64(i)
    }
    natSeenPackets, natPacketX, natPacketY, partOne := make(map[int64]bool), int64(0), int64(0), false
    for {
        numIdleMachines := 0
        for i := 0; i < numMachines; i++ {
            if machines[i].state == haveOutput {
                addr, packetX, packetY := machines[i].getPackets()
                if addr == NAT {
                    natPacketX, natPacketY = packetX, packetY
                    if !partOne {
                        partOne = true
                        fmt.Println("part 1:", natPacketY)
                    }
                } else {
                    inputs[addr] = append(inputs[addr], packetX, packetY)
                }
            } else if machines[i].state == wantInput {
                if len(inputs[i]) > 0 {
                    machines[i].setInput(inputs[i][0])
                    inputs[i] = inputs[i][1:]
                } else {
                    machines[i].setInput(-1)
                    numIdleMachines++
                }
            }
            machines[i].run()
        }
        if numIdleMachines == numMachines {
            if _, ok := natSeenPackets[natPacketY]; ok {
                fmt.Println("part 2:", natPacketY)
                break
            }
            natSeenPackets[natPacketY] = true
            inputs[0] = append(inputs[0], natPacketX, natPacketY)
        }
    }
}
