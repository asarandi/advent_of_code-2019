/* advent of code 2019: day 04, part 02 */
package main

import "fmt"

func isValid(x int) bool {
    prev := 10
    var counts [10]int
    for ; x > 0; x /= 10 {
        z := x % 10
        counts[z]++
        if prev < z {
            return false
        }
        prev = z
    }
    for i := 0; i < 10; i++ {
        if counts[i] == 2 {
            return true
        }
    }
    return false
}

func main() {
    res := 0
    for i := 235741; i < 706948; i++ {
        if isValid(i) {
            res++
        }
    }
    fmt.Println(res)
}
