/* advent of code 2019: day 13, part 02 */
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
	"os"
	"bufio"
	"time"
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
var paddle_y = 0
var ball_y = 0

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

func getchar() byte {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    return input[0]
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

	rows, columns := 22 + 1, 43 + 1
    result := make([][]byte, rows)
    for i := 0; i < rows; i++ {
        result[i] = make([]byte, columns)
        for j := 0; j < columns; j++ {
            result[i][j] = '@'
        }
    }

    i, j := 0, 0
    finished = false
	array[0] = 2
	score := 0
    in := make(chan int64)
    out := make(chan int64)
    go exec(array, in, out)
    for ; !finished; {

		fmt.Printf("\033c")	// clear screen
		for k:=0; k<rows; k++ {
			fmt.Println(string(result[k]))
		}
		fmt.Println("score", score)
		time.Sleep(1 * time.Millisecond)

		if want {
			want = false
			if paddle_y > ball_y {
				in <- -1
			}
			if paddle_y < ball_y {
				in <- 1
			}
			if paddle_y == ball_y {
				in <- 0
			}
		}

        select {
        case y := <-out:
            x := <-out
			tile := <-out
			if y == -1 && x == 0 {
				score = int(tile)
				break
			}
			i = int(x)
			j = int(y)
			var c byte
			switch tile {
				case 0: c = ' '
				case 1: c = '#'
				case 2: c = 'x'
				case 3: c = '-'
						paddle_y = int(y)
				case 4: c = 'o'
						ball_y = int(y)
			default:
				fmt.Println("error")
				break
			}
			result[i][j] = c

        default:
            break
        }

    }
    close(in)
    close(out)
}
