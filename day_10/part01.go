/* advent of code 2019: day 10, part 01 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	_"strconv"
	"strings"
	"math"
)

type point struct {
	x, y int
}

var grid map[point]int

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(content), " \t\r\n\v\f")
	split := strings.Split(s, "\n")
	grid = make(map[point]int)
	for x, _ := range split {
		for y, val := range split[x] {
			if val == '#' {
				grid[point{x,y}] = 0
			}
		}
	}
	res := point{0,0}
	for src, _ := range grid {
		angles := make(map[float64]bool)
		for dst, _ := range grid {
			if src == dst {
				continue
			}
			a := math.Atan2(float64(src.x-dst.x),float64(src.y-dst.y));
			angles[a] = true
		}
		grid[src] = len(angles)
		if grid[src] > grid[res] {
			res = src
		}
	}
//	for k,v := range grid {
//		fmt.Println(k,v)
//	}
	fmt.Println(res, grid[res])
}
