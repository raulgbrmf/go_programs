package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)
import "math"

func plusminus(size int, array []int, types []float64, writer io.Writer) {
	for i := 0; i < size; i++ {
		if array[i] == 0 {
			types[0]++
		} else if float64(array[i]) == math.Abs(float64(array[i])) {
			types[1]++
		} else {
			types[2]++
		}
	}
	for i := 0; i < 3; i++ {
		types[i] = types[i] / float64(size)
	}

	s1 := strconv.FormatFloat(types[1], 'f', -1, 64)
	s2 := strconv.FormatFloat(types[2], 'f', -1, 64)
	s0 := strconv.FormatFloat(types[0], 'f', -1, 64)

	// fmt.Print("Types: \n")
	// fmt.Println(types[1])
	// fmt.Println(types[2])
	// fmt.Println(types[0])

	// fmt.Print("Strings: \n")
	_, _ = io.WriteString(writer, s1)
	// fmt.Print("\n")
	_, _ = io.WriteString(writer, s2)
	// fmt.Print("\n")
	_, _ = io.WriteString(writer, s0)
	// fmt.Print("\n")

	//writer.Write(types[1], "\n")
	//writer.WriteString(types[2], "\n")
	//writer.WriteString(types[0], "\n")
}

func main() {
	var size int
	var arrayPosValue int
	fmt.Scanf("%d", &size)
	array := make([]int, size)
	types := make([]float64, 3)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arrayPosValue)
		array[i] = arrayPosValue
	}
	plusminus(size, array, types, os.Stdout)
}
