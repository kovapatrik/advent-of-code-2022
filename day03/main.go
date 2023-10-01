package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func part1() int {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	prios := 0
	for scanner.Scan() {
		line := scanner.Text()
		p1, p2 := line[:(len(line))/2], line[(len(line))/2:]
		for i := range p1 {
			currentChar := p1[i]
			if strings.Contains(p2, string(currentChar)) {
				if int(currentChar) >= 'a' {
					prios += int(currentChar) - 'a' + 1
				} else {
					prios += int(currentChar) - 'A' + 27
				}
				break
			}
		}
	}
	return prios
}

func part2() int {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	items := make([]string, 3)

	prios := 0
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		items[counter] = line
		counter++
		if counter > 2 {
			sort.Slice(items, func(i, j int) bool {
				return len(items[i]) < len(items[j])
			})

			for i := range items[0] {
				currentChar := items[0][i]
				if strings.Contains(items[1], string(currentChar)) && strings.Contains(items[2], string(currentChar)) {
					if int(currentChar) >= 'a' {
						prios += int(currentChar) - 'a' + 1
					} else {
						prios += int(currentChar) - 'A' + 27
					}
					break
				}
			}

			counter = 0
		}
	}

	return prios
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
