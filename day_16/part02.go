/* advent of code 2019: day 16, part 02 */
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
    s := strings.Trim(string(content), " \t\r\n\v\f")
    offset, _ := strconv.Atoi(s[:7])
    data := make([]int, len(s)*10000)
    for i := 0; i < len(s)*10000; i++ {
        data[i] = int(s[i%len(s)] - '0')
    }
    data = data[offset:]
    for phase := 0; phase < 100; phase++ {
        for i := len(data) - 1; i > 0; i-- {
            data[i-1] = (data[i-1] + data[i]) % 10
        }
    }
    fmt.Printf("part 2: ")
    for i := 0; i < 8; i++ {
        fmt.Printf("%d", data[i])
    }
    fmt.Printf("\n")
}
