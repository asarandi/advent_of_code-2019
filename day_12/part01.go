/* advent of code 2019: day 12, part 01 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "regexp"
    "strconv"
    "strings"
)

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\r\n\v\f")
    split := strings.Split(s, "\n")
    re := regexp.MustCompile("\\<x=(\\-?\\d+),\\s?y=(\\-?\\d+),\\s?z=(\\-?\\d+)\\>")
    pos := make([][]int, 4)
    vel := make([][]int, 4)
    for i, line := range split {
        pos[i] = make([]int, 3)
        vel[i] = make([]int, 3)
        rs := re.FindStringSubmatch(line)
        for j := 0; j < 3; j++ {
            pos[i][j], _ = strconv.Atoi(rs[1+j])
        }
    }
    res := 0
    for steps := 0; steps < 1000; steps++ {
        for i := 0; i < 4; i++ {
            for j := 0; j < 4; j++ {
                for k := 0; k < 3; k++ {
                    if pos[i][k] > pos[j][k] {
                        vel[i][k] -= 1
                        vel[j][k] += 1
                    }
                }
            }
        }
        res = 0
        for i := 0; i < 4; i++ {
            k, p := 0, 0
            for j := 0; j < 3; j++ {
                pos[i][j] += vel[i][j]
                k += abs(pos[i][j])
                p += abs(vel[i][j])
            }
            res += k * p
        }
    }
    fmt.Println("part 1:", res)
}
