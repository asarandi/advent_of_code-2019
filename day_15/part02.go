/* advent of code 2019: day 15, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
	_"os"
	_"bufio"
	_"time"
	_"math/rand"
	"container/list"
)

func getParams(array []int64, index, base int64) (int64, int64, int64, int64) {
    var size, i, j, k int64
    instructionLengths := map[int64]int64{
        1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2, 99: 1}

    size = instructionLengths[array[index]%100]
    if size >= 2 {
        if array[index]/100%10 == 1 {
            i = index + 1
        } else if array[index]/100%10 == 2 {
            i = array[index+1] + base
        } else {
            i = array[index+1]
        }
    }
    if size >= 3 {
        if array[index]/1000%10 == 1 {
            j = index + 2
        } else if array[index]/1000%10 == 2 {
            j = array[index+2] + base
        } else {
            j = array[index+2]
        }
    }
    if size >= 4 {
        if array[index]/10000%10 == 1 {
            log.Fatal("error")
        } else if array[index]/10000%10 == 2 {
            k = array[index+3] + base
        } else {
            k = array[index+3]
        }
    }
    return size, i, j, k
}

var finished = false
var want = false

func exec(array []int64, in, out chan int64) {
    var index, size, i, j, k, base int64
    for index = 0; !finished; {
        size, i, j, k = getParams(array, index, base)
        switch array[index] % 100 {
        case 1:
            array[k] = array[i] + array[j]
        case 2:
            array[k] = array[i] * array[j]
        case 3:
			want = true
            array[i] = <-in
        case 4:
            out <- array[i]
        case 5:
            if array[i] != 0 {
                index = array[j] - size
            }
        case 6:
            if array[i] == 0 {
                index = array[j] - size
            }
        case 7:
            if array[i] < array[j] {
                array[k] = 1
            } else {
                array[k] = 0
            }
        case 8:
            if array[i] == array[j] {
                array[k] = 1
            } else {
                array[k] = 0
            }
        case 9:
            base += array[i]
        case 99:
            finished = true
        default:
            log.Fatal("error")
        }
        index += size
    }
}

type point struct {
    y, x int
}

func (p1 point) add (p2 point) point {
	return point{p1.y + p2.y, p1.x + p2.x}
}

var visited map[point]byte
var current_y, current_x int
var directions = [4]point{{-1,0}, {1,0}, {0,-1}, {0,1}}
var stack *list.List
var backtrack bool

func next_step() int {
	pt := point{current_y, current_x}
	for k:=0;k<4;k++ {
		p := pt.add(directions[k])
		_, ok := visited[p]
		if ok {
			continue
		}
		backtrack = false
		return k
	}
	if stack.Len() == 0 {
		finished = true
		return 0
	}
	k := stack.Remove(stack.Front())
	backtrack = true
	return k.(int) ^ 1
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := strings.Trim(string(content), " \t\n\r\v\f");
    split := strings.Split(s, ",")
    array := make([]int64, len(split)*16)
    for idx, v := range split {
        i, err := strconv.ParseInt(strings.Trim(v, " "), 10, 64)
        if err != nil {
            log.Fatal(err)
        }
        array[idx] = i
    }

    i, j := 0,0
	current_y, current_x = i,j

    finished = false
	visited = make(map[point]byte)
	stack = list.New()
    in := make(chan int64)
    out := make(chan int64)
	dir := -1
	res := -1
	var oxy point
    go exec(array, in, out)
    for ; !finished; {
		if want {
			want = false
			dir = next_step()
			i, j = current_y, current_x
			i += directions[dir].y
			j += directions[dir].x
			in <- int64(dir + 1)
		}

        select {
        case o := <-out:
			switch o {
				case 0: o = '#'
				case 1: o = '.'
					current_y, current_x = i, j
					if !backtrack {
						stack.PushFront(dir)
					}
				case 2: o = 'X'
					current_y, current_x = i, j
					if !backtrack {
						stack.PushFront(dir)
					}
					oxy = point{current_y, current_x}
					res = stack.Len()
			default:
				fmt.Println("error")
				break
			}
			visited[point{i,j}] = byte(o)
        default:
            break
        }

    }
    close(in)
    close(out)
	fmt.Println("part 1:", res)


	oxygen := make(map[point]int)
	queue := list.New()
	oxygen[oxy] = 0
	queue.PushFront(oxy)
	max_g := 0
	for ; queue.Len() > 0 ; {
		parent := queue.Remove(queue.Front()).(point)
		if oxygen[parent] > max_g {
			max_g = oxygen[parent]
		}
		for k:=0; k<4; k++ {
			child := parent.add(directions[k])
			if visited[child] != '.' {
				continue
			}
			if _, ok := oxygen[child]; ok {
				continue
			}
			oxygen[child] = oxygen[parent] + 1
			queue.PushFront(child)
		}
	}
	fmt.Println("part 2:", max_g)
}
