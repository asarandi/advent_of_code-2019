/* advent of code 2019: day 05, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

func get_params(array []int, index int) (int, int, int, int) {
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

func exec(array []int, input int) {
    done := false
    for index := 0; !done; {
        size, i, j, k := get_params(array, index)
        switch (array[index] % 100) {
        case 1:
            array[k] = array[i] + array[j]
        case 2:
            array[k] = array[i] * array[j]
        case 3:
            array[i] = input
        case 4:
            fmt.Println("output", array[i])
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
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\n\r\v\f");
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

    exec(array, 1)
    exec(array_copy, 5)
}
