/* advent of code 2019: day 04, part 01 */
package main

import "fmt"

func is_valid(x int) bool {
    res := false
    prev := -1
    for ; x > 0; x/=10 {
        z := x % 10
        if prev == -1 {
            prev = z
            continue
        }
        if prev < z {
            return false
        }
        res = res || (prev == z)
        prev = z
    }
    return res
}

func main() {
    res := 0
    for i := 235741; i < 706948; i++ {
        if is_valid(i) {
            res++
        }
    }
    fmt.Println(res)
}
