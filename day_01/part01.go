/* advent of code 2019: day 01, part 01 */

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

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
        sum += i / 3 - 2
    }
    fmt.Println(sum)
}
