/* advent of code 2019: day 08, part 01 */
package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	_"strconv"
	"math"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(content), " \t\n\r\v\f")
	data := make([]int, len(s))
	for i:=0; i<len(s); i++ {
		data[i] = int(s[i] - '0')
	}

	var i, j, zeros, ones, twos int

	mi := math.MaxInt32
	res := math.MaxInt32

	for j=0; j<len(data); {
		zeros, ones, twos = 0, 0, 0
		for i=0; i<25*6; i++ {
			if data[i+j] == 0 {
				zeros++
			}
			if data[i+j] == 1 {
				ones++
			}
			if data[i+j] == 2 {
				twos++
			}
		}
		if zeros < mi {
			mi = zeros
			res = ones * twos
		}
		j+=25*6
	}
	fmt.Println(res)
}
