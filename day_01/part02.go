/* advent of code 2019: day 01, part 02 */

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

func f(i int) int {
    i = i / 3 - 2
    if (i < 1) {
        return 0
    }
    return i + f(i)
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    split := strings.Split(string(content), "\n")
    sum := 0
    for _, v := range split {
        if len(v) < 1 {
            continue ;
        }
        i, err := strconv.Atoi(v)
        if err != nil {
            log.Fatal(err)
        }
        sum += f(i)
    }
    fmt.Println(sum)
}
