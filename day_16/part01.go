/* advent of code 2019: day 16, part 01 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func pattern(i, j int) int {
    base := []int{0, 1, 0, -1}
    return base[(j+1)/(i+1)&3]
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\r\n\v\f")
    data := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        data[i] = int(s[i] - '0')
    }
    for phase := 0; phase < 100; phase++ {
        tmp := make([]int, len(data))
        for i := 0; i < len(data); i++ {
            for j := 0; j < len(data); j++ {
                tmp[i] += data[j] * pattern(i, j)
            }
            tmp[i] = abs(tmp[i]) % 10
        }
        data = tmp
    }
    fmt.Printf("part 1: ")
    for i := 0; i < 8; i++ {
        fmt.Printf("%d", data[i])
    }
    fmt.Printf("\n")
}
