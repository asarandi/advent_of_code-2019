/* advent of code 2019: day 15, part 02 */
package main

import (
    "container/list"
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

func getParams(array []int64, index, base int64) (int64, int64, int64, int64) {
    var size, i, j, k int64
    instructionLengths := map[int64]int64{
        1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2, 99: 1}

    size = instructionLengths[array[index]%100]
    if size >= 2 {
        if array[index]/100%10 == 1 {
            i = index + 1
        } else if array[index]/100%10 == 2 {
            i = array[index+1] + base
        } else {
            i = array[index+1]
        }
    }
    if size >= 3 {
        if array[index]/1000%10 == 1 {
            j = index + 2
        } else if array[index]/1000%10 == 2 {
            j = array[index+2] + base
        } else {
            j = array[index+2]
        }
    }
    if size >= 4 {
        if array[index]/10000%10 == 1 {
            log.Fatal("error")
        } else if array[index]/10000%10 == 2 {
            k = array[index+3] + base
        } else {
            k = array[index+3]
        }
    }
    return size, i, j, k
}

var isFinished = false

func exec(array []int64, index int64) (bool, int64, int64) {
    var size, i, j, k, base int64
    for ; !isFinished; {
        size, i, j, k = getParams(array, index, base)
        switch array[index] % 100 {
        case 1:
            array[k] = array[i] + array[j]
        case 2:
            array[k] = array[i] * array[j]
        case 3:
            return true, i, index + size
        case 4:
            return false, i, index + size
        case 5:
            if array[i] != 0 {
                index = array[j] - size
            }
        case 6:
            if array[i] == 0 {
                index = array[j] - size
            }
        case 7:
            if array[i] < array[j] {
                array[k] = 1
            } else {
                array[k] = 0
            }
        case 8:
            if array[i] == array[j] {
                array[k] = 1
            } else {
                array[k] = 0
            }
        case 9:
            base += array[i]
        case 99:
            isFinished = true
        default:
            log.Fatal("error")
        }
        index += size
    }
    return false, -1, -1
}

type point struct {
    y, x int
}

func (p1 point) add(p2 point) point {
    return point{p1.y + p2.y, p1.x + p2.x}
}

var area map[point]int64
var directions = [4]point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var stack *list.List

func nextStep(currentPos point) (bool, int64) {
    for k := 0; k < 4; k++ {
        candidatePos := currentPos.add(directions[k])
        if _, ok := area[candidatePos]; ok {
            continue
        }
        return false, int64(k)
    }
    if stack.Len() == 0 {
        isFinished = true
        return false, 0
    }
    k := stack.Remove(stack.Front())
    return true, k.(int64) ^ 1
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\n\r\v\f")
    split := strings.Split(s, ",")
    array := make([]int64, len(split)*16)
    for idx, v := range split {
        i, err := strconv.ParseInt(strings.Trim(v, " "), 10, 64)
        if err != nil {
            log.Fatal(err)
        }
        array[idx] = i
    }

    currentPos, oxygenPos := point{0, 0}, point{0, 0}
    candidatePos := currentPos.add(directions[0])
    wantInput, backtracking := false, false
    area = make(map[point]int64)
    stack = list.New()
    var index, counter, dir, numSteps int64
    for ; !isFinished; {
        wantInput, index, counter = exec(array, counter)
        if wantInput {
            backtracking, dir = nextStep(currentPos)
            array[index] = dir + 1
            candidatePos = currentPos.add(directions[dir])
        } else {
            if array[index] == 1 || array[index] == 2 {
                currentPos = candidatePos
                if !backtracking {
                    stack.PushFront(dir)
                }
                if array[index] == 2 {
                    oxygenPos = candidatePos
                    numSteps = int64(stack.Len())
                }
            }
            area[candidatePos] = array[index]
        }
    }
    fmt.Println("part 1:", numSteps)

    oxygen := make(map[point]int)
    queue := list.New()
    oxygen[oxygenPos] = 0
    queue.PushFront(oxygenPos)
    maxG := 0
    for ; queue.Len() > 0; {
        parent := queue.Remove(queue.Front()).(point)
        if oxygen[parent] > maxG {
            maxG = oxygen[parent]
        }
        for k := 0; k < 4; k++ {
            child := parent.add(directions[k])
            if area[child] != 1 {
                continue
            }
            if _, ok := oxygen[child]; ok {
                continue
            }
            oxygen[child] = oxygen[parent] + 1
            queue.PushFront(child)
        }
    }
    fmt.Println("part 2:", maxG)
}
