package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() int {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	maxCalorie := 0
	calorie := 0
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		calorie += val

		if err != nil {
			if calorie > maxCalorie {
				maxCalorie = calorie
			}
			calorie = 0
		}
	}

	return maxCalorie
}

func part2() int {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	caloriesTop3 := make([]int, 3)
	calorie := 0
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		calorie += val

		if err != nil {
			for i := range caloriesTop3 {
				if calorie > caloriesTop3[i] {
					if i < len(caloriesTop3)-1 {
						copy(caloriesTop3[i+1:], caloriesTop3[i:])
					}
					caloriesTop3[i] = calorie
					break
				}
			}
			calorie = 0
		}
	}

	return caloriesTop3[0] + caloriesTop3[1] + caloriesTop3[2]
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
