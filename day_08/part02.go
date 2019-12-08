/* advent of code 2019: day 08, part 02 */
package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	_"strconv"
	_"math"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(content), " \t\n\r\v\f")
	data := make([]int, len(s))
	for i:=0; i<len(s); i++ {
		data[i] = int(s[i] - '0')
	}

	var i, j int

	layers := make([][]int,0)

	for j=0; j<len(data); {
		layer := make([]int,25*6)
		for i=0; i<25*6; i++ {
			layer[i] = data[i+j]
		}
		layers = append(layers, layer)
		j+=25*6
	}
	num := j/(25*6)

	final := make([]int,25*6)

	for i=0; i<25*6; i++ {
		done := false
		pixel := 2
		for j=0; j<num && !done; j++ {
			if layers[j][i] == 2 {
				continue
			}
			pixel = layers[j][i]
			break
		}
		final[i] = pixel
	}

	for i=0; i<6; i++ {
		for j=0; j<25; j++ {
			if final[i*25+j] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("o")
			}
//			fmt.Printf("%d", final[i*25+j]);
		}
		fmt.Printf("\n")
	}

//	fmt.Println(layers)
}
