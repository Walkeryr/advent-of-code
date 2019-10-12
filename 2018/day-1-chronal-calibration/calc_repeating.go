package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var frequencies []int
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		number, err := strconv.Atoi(line)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting %s to integer number", line)
		}

		frequencies = append(frequencies, number)
	}

	fmt.Fprintf(os.Stdout, "First repeating frequency: %d\n", CalcRepeatingFrequency(frequencies))
}

// CalcRepeatingFrequency - calculate the first frequency that repeats twice
func CalcRepeatingFrequency(frequencies []int) int {
	i := 0
	current := 0
	checkedFrequencies := make(map[int]int)
	checkedFrequencies[current]++

	for {
		current += frequencies[i]

		if checkedFrequencies[current] > 0 {
			break
		}

		checkedFrequencies[current]++

		if i == len(frequencies)-1 {
			i = 0
		} else {
			i++
		}
	}

	return current
}
