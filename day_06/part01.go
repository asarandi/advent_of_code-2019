/* advent of code 2019: day 06, part 01 */

package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(content), " \n\t\r\f\v")
	orbits := strings.Split(s, "\n")
	childParent := make(map[string]string)
	for _, orbit := range orbits {
		split := strings.Split(orbit, ")")
		childParent[split[1]] = split[0]
	}
	res := 0
	for k, _ := range childParent {
		for ; childParent[k] != ""; {
			res++
			k = childParent[k]
		}
	}
	fmt.Println(res)
}
