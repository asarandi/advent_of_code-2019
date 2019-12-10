/* advent of code 2019: day 10, part 02 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"math"
)

type asteroid struct {
	x, y int
	d int
	a float64
}

var asteroids []asteroid

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (a asteroid) distance (b asteroid) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func (a asteroid) angle (b asteroid) float64 {
	return math.Atan2(float64(a.x-b.x), float64(a.y-b.y))
}

type ByDistance []asteroid
type ByAngle    []asteroid

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool {
	if a[i].d == a[j].d {
		return a[i].a < a[j].a
	}
	return a[i].d < a[j].d
}

func (a ByAngle)    Len() int           { return len(a) }
func (a ByAngle)    Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAngle)    Less(i, j int) bool {
	if a[i].a == a[j].a {
		return a[i].d < a[j].d
	}
	return a[i].a < a[j].a
}

var base = asteroid{22,17,0,0}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(content), " \t\r\n\v\f")
	split := strings.Split(s, "\n")
	asteroids = make([]asteroid,0)
	for x, _ := range split {
		for y, val := range split[x] {
            if x == base.x && y == base.y {
                continue
            }
			if val == '#' {
				ast := asteroid{x,y,0,0}
				ast.d = base.distance(ast)
				ast.a = base.angle(ast)
				asteroids = append(asteroids,ast)
			}
		}
	}

    count:=0
	sort.Sort(ByDistance(asteroids))
	sort.Sort(ByAngle(asteroids))

    angles := make(map[float64]bool)

	for i:=0; i<len(asteroids); i++ {
		if asteroids[i].a < math.Pi/2 || angles[asteroids[i].a] {
			continue
		}
        angles[asteroids[i].a] = true
        asteroids = append(asteroids[:i], asteroids[i+1:]...)
        i--
        count += 1
    }
//    fmt.Println(count)

	for i:=0; i<len(asteroids); i++ {
		if asteroids[i].a >= math.Pi/2 || angles[asteroids[i].a] {
			continue
		}
        angles[asteroids[i].a] = true
        if count + 1 == 200 {
            fmt.Println("result", asteroids[i].y*100+asteroids[i].x)
            break
        }
        asteroids = append(asteroids[:i], asteroids[i+1:]...)
        i--
        count += 1
    }

//    fmt.Println(count)

}
