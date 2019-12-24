/* advent of code 2019: day 24, part 1 */
package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

const SIZE = 5
const BUG = '#'
const SPACE = '.'

type board [][]byte

func (b board) toString() string {
    var res string
    for i := 0; i < SIZE; i++ {
        res = res + string(b[i]) + "\n"
    }
    return res
}

func (b board) countAdjacent(i, j int) int {
    res := 0
    if i > 0 && b[i-1][j] == BUG {
        res++
    }
    if j > 0 && b[i][j-1] == BUG {
        res++
    }
    if i+1 < len(b) && b[i+1][j] == BUG {
        res++
    }
    if j+1 < len(b) && b[i][j+1] == BUG {
        res++
    }
    return res
}

func (b board) emptyCopy() board {
    res := make([][]byte, len(b))
    for i := 0; i < len(b); i++ {
        res[i] = make([]byte, len(b[i]))
        for j := 0; j < len(b[i]); j++ {
            res[i][j] = SPACE
        }
    }
    return res
}

func (b board) age() board {
    res := b.emptyCopy()
    for i := 0; i < len(b); i++ {
        for j := 0; j < len(b[i]); j++ {
            count := b.countAdjacent(i, j)
            if b[i][j] == BUG && count == 1 {
                res[i][j] = BUG
            }
            if b[i][j] == SPACE && (count == 1 || count == 2) {
                res[i][j] = BUG
            }
        }
    }
    return res
}

func (b board) rating() int {
    res := 0
    for i := 0; i < len(b); i++ {
        for j := 0; j < len(b[i]); j++ {
            if b[i][j] == BUG {
                res += 1 << uint(i*len(b[i])+j)
            }
        }
    }
    return res
}

func fromString(s string) board {
    trim := strings.Trim(s, " \t\r\n\v\f")
    split := strings.Split(trim, "\n")
    data := make([][]byte, len(split))
    for i := range split {
        data[i] = make([]byte, len(split[i]))
        for j := range split[i] {
            data[i][j] = split[i][j]
        }
    }
    return data
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    b := fromString(string(content))
    seen := make(map[string]int)
    for {
        s := b.toString()
        if val, ok := seen[s]; ok {
            fmt.Println("part 1:", val)
            break
        }
        seen[s] = b.rating()
        b = b.age()
    }
}
