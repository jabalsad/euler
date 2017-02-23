package main

// https://projecteuler.net/problem=81

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	out, err := os.Create("sums.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(file)
	sums := [][]int{}
	row := 0
	var dim int

	for scanner.Scan() {
		stringArray := strings.Split(scanner.Text(), ",")
		dim = len(stringArray)
		sums = append(sums, make([]int, dim))
		for col, s := range stringArray {
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}

			switch {
			case row == 0 && col == 0:
				sums[row][col] = v
			case row == 0:
				sums[row][col] = v + sums[row][col-1]
			case col == 0:
				sums[row][col] = v + sums[row-1][col]
			default:
				sums[row][col] = v + min(sums[row-1][col], sums[row][col-1])
			}
		}
		rowStrings := make([]string, len(sums[row]))
		for i, r := range sums[row] {
			rowStrings[i] = strconv.Itoa(r)
		}
		_, err := out.WriteString(strings.Join(rowStrings, ",") + "\n")
		if err != nil {
			log.Fatal(err)
		}

		row = row + 1
	}

	fmt.Println(sums[dim-1][dim-1])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
