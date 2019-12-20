/* advent of code 2019: day 07, part 01 */
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

func exec(array []int, phase, input int) int {
    done := false
    flag := false
    res := -1
    for index := 0; !done; {
        size, i, j, k := getParams(array, index)
        switch array[index] % 100 {
        case 1:
            array[k] = array[i] + array[j]
        case 2:
            array[k] = array[i] * array[j]
        case 3:
            if !flag {
                array[i] = phase
                flag = true
            } else {
                array[i] = input
            }
        case 4:
            res = array[i]
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
        case 99:
            done = true
        default:
            log.Fatal("error")
        }
        index += size
    }
    return res
}

func get_phases(x int) ([]int, bool) {
    phases := make([]int, 5)
    for i := 4; i >= 0; i-- {
        k := x % 10
        if k > 4 { /* XXX */
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

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\n\r\v\f")
    split := strings.Split(s, ",")
    array := make([]int, len(split))
    for idx, v := range split {
        i, err := strconv.Atoi(v)
        if err != nil {
            log.Fatal(err)
        }
        array[idx] = i
    }
    array_copy := make([]int, len(array))
    copy(array_copy, array)

    res := math.MinInt32

    for j := 0; j < 44444; j++ {
        phases, valid := get_phases(j)
        if !valid {
            continue
        }

        input := 0
        for i := 0; i < 5; i++ {
            copy(array, array_copy)
            input = exec(array, phases[i], input)
        }

        if input > res {
            res = input
        }
    }
    fmt.Println(res)
}
