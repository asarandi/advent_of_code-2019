/* advent of code 2019: day 10, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "math"
    "sort"
    "strings"
)

type asteroid struct {
    x, y, n, d int
    a          float64
}

var asteroids []asteroid

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func (a asteroid) distance(b asteroid) int {
    return abs(a.x-b.x) + abs(a.y-b.y)
}

func (a asteroid) angle(b asteroid) float64 {
    return math.Atan2(float64(a.x-b.x), float64(a.y-b.y))
}

func (a asteroid) is(b asteroid) bool {
    return a.x == b.x && a.y == b.y
}

func (a asteroid) inQ1() bool {
    if a.a >= math.Pi/2 && a.a <= math.Pi {
        return true
    }
    return false
}

type ByAnglesDistance []asteroid
func (a ByAnglesDistance) Len() int      { return len(a) }
func (a ByAnglesDistance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAnglesDistance) Less(i, j int) bool {
    if a[i].a == a[j].a {
        return a[i].d < a[j].d
    }
    return a[i].a < a[j].a
}

var mapIntAsteroid map[int]asteroid
var asteroid_id = 1

func setAsteroidIds(base asteroid, cond bool) bool {
    res := false
    angles := make(map[float64]bool)
    for i := range asteroids {
        if asteroids[i].inQ1() != cond ||
            asteroids[i].n != 0 ||
            base.is(asteroids[i]) ||
            angles[asteroids[i].a] {
            continue
        }
        angles[asteroids[i].a] = true
        mapIntAsteroid[asteroid_id] = asteroids[i]
        asteroids[i].n = asteroid_id
        asteroid_id += 1
        res = true
    }
    return res
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\r\n\v\f")
    split := strings.Split(s, "\n")
    asteroids = make([]asteroid, 0)
    for x, _ := range split {
        for y, val := range split[x] {
            if val == '#' {
                asteroids = append(asteroids, asteroid{x, y, 0, 0, 0})
            }
        }
    }
    base := asteroids[0]
    base_num_angles := 0
    for _, a := range asteroids {
        angles := make(map[float64]bool)
        for _, b := range asteroids {
            angles[a.angle(b)] = true
        }
        if len(angles) > base_num_angles {
            base = a
            base_num_angles = len(angles)
        }
    }
    fmt.Println("part 1:", base_num_angles)
    for i := range asteroids {
        asteroids[i].d = base.distance(asteroids[i])
        asteroids[i].a = base.angle(asteroids[i])
    }
    sort.Sort(ByAnglesDistance(asteroids))
    mapIntAsteroid = make(map[int]asteroid)
    for cont := true; cont; {
        cont = cont && setAsteroidIds(base, true)
        cont = cont && setAsteroidIds(base, false)
    }
    a := mapIntAsteroid[200]
    fmt.Println("part 2:", a.y*100+a.x)
}
