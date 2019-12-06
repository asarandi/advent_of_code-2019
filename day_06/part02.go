/* advent of code 2019: day 06, part 02 */

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
	dist := make(map[string]int)
	o := "YOU"
	d := 0
	for {
		if childParent[o] == "" {
			break
		}
		dist[o] = d
		d++
		o = childParent[o]
	}
	o = "SAN"
	d = 0
	for {
		if dist[o] != 0 {
			break
		}
		d++
		o = childParent[o]
	}
	fmt.Println(d + dist[o] - 2)
}
