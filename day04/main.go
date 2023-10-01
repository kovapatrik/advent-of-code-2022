package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	containCnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		l, r, _ := strings.Cut(line, ",")
		l1, l2, _ := strings.Cut(l, "-")
		l1Num, _ := strconv.Atoi(l1)
		l2Num, _ := strconv.Atoi(l2)

		r1, r2, _ := strings.Cut(r, "-")
		r1Num, _ := strconv.Atoi(r1)
		r2Num, _ := strconv.Atoi(r2)

		if (l1Num >= r1Num && l2Num <= r2Num) || (l1Num <= r1Num && l2Num >= r2Num) {
			containCnt++
		}
	}
	return containCnt
}

func part2() int {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	overlapCnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		l, r, _ := strings.Cut(line, ",")
		l1, l2, _ := strings.Cut(l, "-")
		l1Num, _ := strconv.Atoi(l1)
		l2Num, _ := strconv.Atoi(l2)

		r1, r2, _ := strings.Cut(r, "-")
		r1Num, _ := strconv.Atoi(r1)
		r2Num, _ := strconv.Atoi(r2)

		if l1Num <= r2Num && r1Num <= l2Num {
			overlapCnt++
		}
	}
	return overlapCnt
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
