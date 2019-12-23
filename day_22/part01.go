/* advent of code 2019: day 22, day 1 */
package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    split := strings.Split(strings.Trim(string(content), " \t\r\n\f\v"), "\n")
    size := 10007
    deck := make([]int, size)
    for i := 0; i < size; i++ {
        deck[i] = i
    }
    for _, line := range split {
        words := strings.Split(strings.Trim(line, " \t\n"), " ")
        idx, _ := strconv.Atoi(words[len(words)-1])
        switch words[1] {
        case "with": //increment
            newDeck := make([]int, size)
            for j, k := 0, 0; k < size; k++ {
                newDeck[j] = deck[k]
                j = (j + idx) % size
            }
            deck = newDeck
        case "into": //new stack
            for k := 0; k < size>>1; k++ {
                deck[k], deck[size-1-k] = deck[size-1-k], deck[k]
            }
        default: //cut
            if idx < 0 {
                idx += size
            }
            newDeck := append(deck[idx:], deck[:idx]...)
            deck = newDeck
        }
    }
    for i := 0; i < size; i++ {
        if deck[i] == 2019 {
            fmt.Println("part 1:", i)
            break
        }
    }
}
