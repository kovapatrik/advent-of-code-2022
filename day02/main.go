package main

import (
	"bufio"
	"fmt"
	"os"
)

func modNeg(v, m int) int {
	return (v%m + m) % m
}

func part1() int {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	point := 0
	for scanner.Scan() {
		line := scanner.Text()
		l, r := int(line[0])-'A', int(line[2])-'X'

		point += 1 + r + 3*modNeg(1+r-l, 3)
	}

	return point
}

func part2() int {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	point := 0
	for scanner.Scan() {
		line := scanner.Text()
		l, r := int(line[0])-'A', int(line[2])-'X'

		point += r*3 + 1 + (2+l+r)%3
	}

	return point
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
