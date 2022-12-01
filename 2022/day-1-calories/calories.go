package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	calories_grouped := make([][]int, 0)
	calories_grouped = append(calories_grouped, make([]int, 0))

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			// If the line is empty, we start a new group of calories
			calories_grouped = append(calories_grouped, make([]int, 0))
		} else {
			// Otherwise, we add the calories to the last group
			last_index := len(calories_grouped) - 1
			calories_grouped[last_index] = append(calories_grouped[last_index], getCalories(line))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Maximum calories:", getMaximumCalories(calories_grouped))
}

func getCalories(line string) int {
	number, err := strconv.Atoi(line)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting %s to integer number", line)
	}

	return number
}

func getMaximumCalories(calories_grouped [][]int) int {
	maximum_calories := 0

	for _, group := range calories_grouped {
		// calculate the sum of the calories for each group
		sum := 0
		for _, calories := range group {
			sum += calories
		}

		if sum > maximum_calories {
			maximum_calories = sum
		}
	}

	return maximum_calories
}
