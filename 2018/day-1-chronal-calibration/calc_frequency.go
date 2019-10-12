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

	fmt.Fprintf(os.Stdout, "Result frequency: %d\n", CalcResultFrequency(frequencies))
}

// CalcResultFrequency - calculate the resulting frequency
func CalcResultFrequency(frequencies []int) (sum int) {
	for _, number := range frequencies {
		sum += number
	}

	return sum
}
