/* advent of code 2019: day 09, part 01 */
package main

import (
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

func exec(array, input []int64) []int64 {
    var index, size, i, j, k, base int64
    done := false
    output := make([]int64,0)
    for index = 0; !done; {
        size, i, j, k = getParams(array, index, base)
        switch array[index] % 100 {
        case 1:
            array[k] = array[i] + array[j]
        case 2:
            array[k] = array[i] * array[j]
        case 3:
            array[i] = input[0]
            input = input[1:]
        case 4:
            output = append(output, array[i])
        case 5:
            if array[i] != 0 { index = array[j] - size }
        case 6:
            if array[i] == 0 { index = array[j] - size }
        case 7:
            if array[i] < array[j] { array[k] = 1 } else { array[k] = 0 }
        case 8:
            if array[i] == array[j] { array[k] = 1 } else { array[k] = 0 }
        case 9:
            base += array[i]
        case 99:
            done = true
        default:
            log.Fatal("error")
        }
        index += size
    }
    return output
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\n\r\v\f");
    split := strings.Split(s, ",")
    array := make([]int64, len(split)*16)
    for idx, v := range split {
        i, err := strconv.ParseInt(v, 10, 64)
        if err != nil {
            log.Fatal(err)
        }
        array[idx] = i
    }
    fmt.Println("part 1:", exec(array, []int64{1})[0])
    fmt.Println("part 2:", exec(array, []int64{2})[0])
}
