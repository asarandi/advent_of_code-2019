/* advent of code 2019: day 07, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
	"math"
)

func get_params(array []int, index int) (int, int, int, int) {
    var size, i, j, k int
    instructionLengths := map[int]int{
        1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 99: 1}

    size = instructionLengths[array[index]%100]
    if size >= 2 {
        if array[index]/100%10 == 1 {
            i = index + 1
        } else {
            i = array[index+1]
        }
    }
    if size >= 3 {
        if array[index]/1000%10 == 1 {
            j = index + 2
        } else {
            j = array[index+2]
        }
    }
    if size >= 4 {
        k = array[index+3]
    }
    return size, i, j, k
}

var first bool
var flags []bool
var phases []int
var outputs [5]chan int
var finished [5]chan bool

func exec(array []int, idx int) {
    done := false
    for index := 0; !done; {
        size, i, j, k := get_params(array, index)
        switch array[index] % 100 {
        case 1:
            array[k] = array[i] + array[j]
        case 2:
            array[k] = array[i] * array[j]
        case 3:
			if !flags[idx] {
	            array[i] = phases[idx]
				flags[idx] = true
			} else {
				if idx == 0 && !first {
					array[i] = 0
					first = true
				} else {
					array[i] = <-outputs[(idx + 4) % 5]
				}
			}
        case 4:
            outputs[idx] <- array[i]
        case 5:
            if array[i] != 0 { index = array[j] - size }
        case 6:
            if array[i] == 0 { index = array[j] - size }
        case 7:
            if array[i] < array[j] { array[k] = 1 } else { array[k] = 0 }
        case 8:
            if array[i] == array[j] { array[k] = 1 } else { array[k] = 0 }
        case 99:
            done = true
        default:
            log.Fatal("error")
        }
        index += size
    }
	finished[idx] <- true
}

func get_phases(x int) bool {
	for i:=4; i>=0; i-- {
		k := x % 10
		if k < 5 {				/* XXX */
			return false
		}
		phases[i] = k
		x /= 10
	}
	for i:=0; i<4; i++ {
		for j:=i+1; j<5; j++ {
			if phases[i] == phases[j] {
				return false
			}
		}
	}
	return true
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\n\r\v\f");
    split := strings.Split(s, ",")
	programs := make([][]int, 5)
	programs_copy := make([][]int, 5)

	for k:=0;k<5;k++ {
		programs[k] = make([]int, len(split))
		programs_copy[k] = make([]int, len(split))
		for idx, v := range split {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			programs[k][idx] = i
			programs_copy[k][idx] = i
		}
	}

	res := math.MinInt32
	flags = make([]bool, 5)
	phases = make([]int, 5)
	for i := range outputs {
		outputs[i] = make(chan int)
		finished[i] = make(chan bool)
	}

	for j:=55555; j<99999; j++ {
		if !get_phases(j) {
			continue
		}
		first = false
		for i:=0;i<5;i++ {
			flags[i] = false
			copy(programs[i], programs_copy[i])
			go exec(programs[i], i)
		}
		for i:=0;i<4;i++ {
			<-finished[i]
		}

		candidate := <-outputs[4]
		<-finished[4]
		if candidate > res {
			res = candidate
		}
	}
	fmt.Println(res)
}
