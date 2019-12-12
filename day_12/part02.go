/* advent of code 2019: day 12, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
	"strconv"
	"regexp"
)

const (
    pos = iota
    pos_copy
    vel
    vel_copy
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}

func lcm(a, b int) int {
    return abs(a * b) / gcd(a, b)
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\r\n\v\f")
    split := strings.Split(s, "\n")
	re := regexp.MustCompile("\\<x=(\\-?\\d+),\\s?y=(\\-?\\d+),\\s?z=(\\-?\\d+)\\>")
    a := make([][][]int, 4)
    for i:=0; i<4; i++ {
        a[i] = make([][]int, 4)
        for j:=0; j<4; j++ {
            a[i][j] = make([]int, 3)
        }
    }
	for i, line := range split {
		rs := re.FindStringSubmatch(line)
		for j:=0; j<3; j++ {
			a[pos][i][j], _ = strconv.Atoi(rs[1+j])
			a[pos_copy][i][j], _ = strconv.Atoi(rs[1+j])
		}
	}

    cycles := make([]int, 3)
    steps := -1
	for {
        steps++
		for i:=0; i<4; i++ {
			for j:=0; j<4; j++ {
				for k:=0; k<3; k++ {
					if a[pos][i][k] > a[pos][j][k] {
						a[vel][i][k] -= 1
						a[vel][j][k] += 1
					}
				}
			}
		}
		for i:=0; i<4; i++ {
	        for j:=0; j<3; j++ {
				a[pos][i][j] += a[vel][i][j]
            }
        }

        cy := 0
	    for j:=0; j<3; j++ {
            if cycles[j] != 0 {
                cy++
                continue
            }
            f := true
		    for i:=0; i<4; i++ {
                f = f && (a[pos][i][j] == a[pos_copy][i][j])
                f = f && (a[vel][i][j] == a[vel_copy][i][j])
			}
            if f {
                cycles[j] = steps + 1
            }
		}
        if cy == 3 {
            break
        }
	}
    res := lcm(lcm(cycles[0], cycles[1]), cycles[2])
    fmt.Println("part 2:", res)
}
