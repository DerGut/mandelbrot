package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	from, err := strconv.ParseComplex(os.Args[1], 128)
	if err != nil {
		panic(err)
	}
	to, err := strconv.ParseComplex(os.Args[2], 128)
	if err != nil {
		panic(err)
	}
	stepSize, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		panic(err)
	}

	frame := buildFrame(from, to, stepSize)
	printFrame(frame)
}

func buildFrame(from, to complex128, stepSize float64) [][]bool {
	var frame [][]bool
	for r := real(from); r < real(to); r += stepSize {
		var row []bool
		for i := imag(from); i < imag(to); i += stepSize {
			row = append(row, isDiverging(complex(r, i)))
		}
		frame = append(frame, row)
	}
	return frame
}

func printFrame(frame [][]bool) {
	for _, row := range frame {
		for _, cell := range row {
			if cell {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

const ITERATION_LIMIT = 100

func isDiverging(c complex128) bool {
	n := complex(0, 0)
	for i := 0; i < ITERATION_LIMIT; i++ {
		n = mandelbrot(n, c)
	}

	return math.IsNaN(real(n)) || math.IsNaN(imag(n))
}

func mandelbrot(n, c complex128) complex128 {
	return n*n + c
}
