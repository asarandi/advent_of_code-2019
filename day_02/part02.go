/* advent of code 2019: day 02, part 02 */

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

func main() {
    var i int

    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    split := strings.Split(strings.Trim(string(content), "\n"), ",")
    array := make([]int, len(split))
    for i = range array {
        array[i], err = strconv.Atoi(split[i])
        if err != nil {
            log.Fatal(err)
        }
    }

    arrayCopy := make([]int, len(array))
    copy(arrayCopy, array)
    done := false

    for j := 0; j <= 99 && !done; j++ {
        for k := 0; k <= 99 && !done; k++ {
            copy(array, arrayCopy)
            array[1] = j
            array[2] = k
            for i = 0; i < len(array); i += 4 {
                var idx0, idx1, idx2, res int

                if array[i] == 99 {
                    break
                }
                idx0, idx1, idx2 = array[i+1], array[i+2], array[i+3]
                if array[i] == 1 {
                    res = array[idx0] + array[idx1]
                } else if array[i] == 2 {
                    res = array[idx0] * array[idx1]
                } else {
                    log.Fatal("error")
                }
                array[idx2] = res
            }
            if array[0] == 19690720 {
                done = true
                fmt.Printf("result: 100 * %d + %d = %d\n", j, k, 100*j+k)
            }
        }
    }
}
