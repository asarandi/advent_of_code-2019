/* advent of code 2019: day 07, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "math"
    "strconv"
    "strings"
)

func getParams(array []int, index int) (int, int, int, int) {
    var size, i, j, k int
    instructionLengths := map[int]int{
        1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 99: 1}

    size = instructionLengths[array[index]%100]
    if size >= 2 {
        if array[index]/100%10 == 1 {
            i = index + 1
        } else {
            i = array[index+1]
        }
    }
    if size >= 3 {
        if array[index]/1000%10 == 1 {
            j = index + 2
        } else {
            j = array[index+2]
        }
    }
    if size >= 4 {
        k = array[index+3]
    }
    return size, i, j, k
}

var programs [][]int
var programsCopy [][]int
var inputs [][]int
var pcs []int

//returns -1 when program halts or next pc position
func execute(amp int) int {
    program := programs[amp]
    pc := pcs[amp]
    for ; pc != -1; {
        size, i, j, k := getParams(program, pc)
        switch program[pc] % 100 {
        case 1:
            program[k] = program[i] + program[j]
        case 2:
            program[k] = program[i] * program[j]
        case 3:
            program[i] = inputs[amp][0]
            inputs[amp] = inputs[amp][1:] //dequeue first
        case 4:
            inputs[(amp+1)%5] = append(inputs[(amp+1)%5], program[i]) //enqueue last
            return pc + size
        case 5:
            if program[i] != 0 {
                pc = program[j] - size
            }
        case 6:
            if program[i] == 0 {
                pc = program[j] - size
            }
        case 7:
            if program[i] < program[j] {
                program[k] = 1
            } else {
                program[k] = 0
            }
        case 8:
            if program[i] == program[j] {
                program[k] = 1
            } else {
                program[k] = 0
            }
        case 99:
            return -1
        default:
            log.Fatal("error")
        }
        pc += size
    }
    return -1
}

func getPhases(x int) ([]int, bool) {
    phases := make([]int, 5)
    for i := 4; i >= 0; i-- {
        k := x % 10
        if k < 5 { /* XXX */
            return nil, false
        }
        phases[i] = k
        x /= 10
    }
    for i := 0; i < 4; i++ {
        for j := i + 1; j < 5; j++ {
            if phases[i] == phases[j] {
                return nil, false
            }
        }
    }
    return phases, true
}

func isFinished() bool {
    res := true
    for i := 0; i < 5; i++ {
        res = res && pcs[i] == -1
    }
    return res
}

func prepare() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\n\r\v\f")
    split := strings.Split(s, ",")

    programs = make([][]int, 5)
    programsCopy = make([][]int, 5)
    inputs = make([][]int, 5)
    pcs = make([]int, 5)

    for k := 0; k < 5; k++ {
        programs[k] = make([]int, len(split))
        programsCopy[k] = make([]int, len(split))
        for idx, v := range split {
            i, err := strconv.Atoi(v)
            if err != nil {
                log.Fatal(err)
            }
            programs[k][idx] = i
            programsCopy[k][idx] = i
        }
    }
}

func main() {

    prepare()

    res := math.MinInt32

    for j := 55555; j < 99999; j++ {
        phases, valid := getPhases(j)
        if !valid {
            continue
        }

        for i := 0; i < 5; i++ {
            copy(programs[i], programsCopy[i])
            inputs[i] = make([]int, 0)
            inputs[i] = append(inputs[i], phases[i])
            pcs[i] = 0
        }

        inputs[0] = append(inputs[0], 0) //additional input for first amp only

        for i := 0; !isFinished(); {
            pcs[i] = execute(i)
            i = (i + 1) % 5
        }
        if inputs[0][0] > res {
            res = inputs[0][0]
        }
    }

    fmt.Println(res)
}
