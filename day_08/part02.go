/* advent of code 2019: day 08, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\n\r\v\f")
    for i := 0; i < 150; i++ {
        for j := 0; j < len(s); j += 150 {
            if s[i+j] == '2' {
                continue
            }
            if s[i+j] == '1' {
                fmt.Printf("#")
            } else {
                fmt.Printf(".")
            }
            break
        }
        if (i+1)%25 == 0 {
            fmt.Println()
        }
    }
}
