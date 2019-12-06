/* advent of code 2019: day 05, part 02 */
package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

/*

opcode 1 - add parameter 1 to parameter 2, store result at parameter 3
opcode 2 - multiply parameter 1 to parameter 2, store result at parameter 3
opcode 3 - save input to parameter
opcode 4 - output parameter
opcode 5 - jump if true
opcode 6 - jump if false
opcode 7 - less than
opcode 8 - equals
opcode 99 - halt

*/

func p1(x int) int {
	return (x / 100) % 10
}

func p2(x int) int {
	return (x / 1000) % 10
}

func p3(x int) int {
	return (x / 10000) % 10
}


func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(content), " \t\n\r\v\f");
	split := strings.Split(s, ",")
	a := make([]int, len(split))
	for idx, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		a[idx] = i
	}
	done := false
	for i:=0; !done; {
		var val1, val2, size int
		switch(a[i] % 100) {
		case 1:
			size = 4
			if p1(a[i])==1 { val1 = a[i+1] } else { val1 = a[a[i+1]] }
			if p2(a[i])==1 { val2 = a[i+2] } else { val2 = a[a[i+2]] }
			a[a[i+3]] = val1 + val2
		case 2:
			size = 4
			if p1(a[i])==1 { val1 = a[i+1] } else { val1 = a[a[i+1]] }
			if p2(a[i])==1 { val2 = a[i+2] } else { val2 = a[a[i+2]] }
			a[a[i+3]] = val1 * val2
		case 3:
			size = 2
			a[a[i+1]] = 5	/* XXX */
		case 4:
			size = 2
			if p1(a[i])==1 { val1 = a[i+1] } else { val1 = a[a[i+1]] }
			fmt.Println("output", val1)

		case 5:
			size = 3
			if p1(a[i])==1 { val1 = a[i+1] } else { val1 = a[a[i+1]] }
			if p2(a[i])==1 { val2 = a[i+2] } else { val2 = a[a[i+2]] }
			if val1 != 0 { i = val2 - size }

		case 6:
			size = 3
			if p1(a[i])==1 { val1 = a[i+1] } else { val1 = a[a[i+1]] }
			if p2(a[i])==1 { val2 = a[i+2] } else { val2 = a[a[i+2]] }
			if val1 == 0 { i = val2 - size }

		case 7:
			size = 4
			if p1(a[i])==1 { val1 = a[i+1] } else { val1 = a[a[i+1]] }
			if p2(a[i])==1 { val2 = a[i+2] } else { val2 = a[a[i+2]] }
			if val1 < val2 { a[a[i+3]] = 1 } else { a[a[i+3]] = 0 }

		case 8:
			size = 4
			if p1(a[i])==1 { val1 = a[i+1] } else { val1 = a[a[i+1]] }
			if p2(a[i])==1 { val2 = a[i+2] } else { val2 = a[a[i+2]] }
			if val1 == val2 { a[a[i+3]] = 1 } else { a[a[i+3]] = 0 }

		case 99:
			size = 1
			done = true
		default:
			log.Fatal("error")
		}
		i += size
	}
}
