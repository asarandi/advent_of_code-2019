/* advent of code 2019: day 03, part 02 */
package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
)

type vertex struct {
    y, x int
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func trace(r *bufio.Reader) map[vertex]int {
    s, err := r.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    split := strings.Split(strings.Trim(s, " \t\n\r\v\f"), ",")
    y0, x0 := 0, 0
    steps := 0
    res := make(map[vertex]int)
    for _, s := range split {
        var y1, x1 int
        switch s[0] {
        case 'R':
            y1, x1 = 0, 1
        case 'D':
            y1, x1 = 1, 0
        case 'L':
            y1, x1 = 0, -1
        case 'U':
            y1, x1 = -1, 0
        default:
            log.Fatal("error")
        }
        dst, _ := strconv.Atoi(s[1:])
        for i := 0; i < dst; i++ {
            y0 += y1
            x0 += x1
            steps++
            res[vertex{y0, x0}] = steps
        }
    }
    return res
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("please provide file name")
    }
    fp, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer fp.Close()

    r := bufio.NewReader(fp)
    red := trace(r)
    blue := trace(r)

    res := math.MaxInt32
    for k, v := range red {
        if blue[k] != 0 {
            t := v + blue[k]
            if t < res {
                res = t
            }
        }
    }
    fmt.Println(res)
}
