/* advent of code 2019: day 08, part 01 */
package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"math"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(content), " \t\n\r\v\f")

	minZeros := math.MaxInt32
	res := math.MaxInt32

	for i:=0; i<len(s);i+=150 {
		zeros, ones, twos := 0, 0, 0
		for j:=0; j<150; j++ {
			if s[i+j] == '0' {
				zeros++
			}
			if s[i+j] == '1' {
				ones++
			}
			if s[i+j] == '2' {
				twos++
			}
		}
		if zeros < minZeros {
			minZeros = zeros
			res = ones * twos
		}
	}
	fmt.Println(res)
}
