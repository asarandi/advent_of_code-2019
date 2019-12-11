/* advent of code 2019: day 11, part 02 */
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

var finished = false

func exec(array []int64, in, out chan int64) {
    var index, size, i, j, k, base int64
    for index = 0; !finished; {
        size, i, j, k = getParams(array, index, base)
        switch array[index] % 100 {
        case 1:
            array[k] = array[i] + array[j]
        case 2:
            array[k] = array[i] * array[j]
        case 3:
            array[i] = <-in
        case 4:
            out <- array[i]
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
            finished = true
        default:
            log.Fatal("error")
        }
        index += size
    }
}

type point struct {
    y, x int
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
        i, err := strconv.ParseInt(strings.Trim(v, " "), 10, 64)
        if err != nil {
            log.Fatal(err)
        }
        array[idx] = i
    }

    rows, columns := 6, 40
    result := make([][]byte, rows)
    for i := 0; i < rows; i++ {
        result[i] = make([]byte, columns)
        for j := 0; j < columns; j++ {
            result[i][j] = '.'
        }
    }
    i, j, direction := 0, 0, 0
    moves := []point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //URDL
    turns := []int{3, 1}
    visited := make(map[point]int64)
    finished = false

    in := make(chan int64)
    out := make(chan int64)
    go exec(array, in, out)
    in <- 1
    for ; !finished; {
        select {
        case color := <-out:
            turn := <-out
            visited[point{i, j}] = color
            if color == 1 {
                result[i][j] = '#'
            }
            direction = (direction + turns[turn]) & 3
            i += moves[direction].y
            j += moves[direction].x
            if !finished {
                in <- visited[point{i, j}]
            }
        default:
            break
        }
    }
    close(in)
    close(out)

    for i = 0; i < rows; i++ {
        fmt.Println(string(result[i]))
    }
}
