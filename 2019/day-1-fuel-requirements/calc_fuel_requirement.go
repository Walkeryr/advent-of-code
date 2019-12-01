package main

import (
	"math"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var masses []int
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		number, err := strconv.Atoi(line)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting %s to integer number", line)
		}

		masses = append(masses, number)
	}

	fmt.Fprintf(os.Stdout, "Sum of the fuel requirements: %d\n", SumFuelRequirements(masses))
}

// CalcFuelRequirement - calculate the fuel required to launch given module mass
func CalcFuelRequirement(mass int) int {
	return int(math.Floor(float64(mass) / 3.0) - 2)
}

// SumFuelRequirements - calculate resulting sum of all fuel requirements
func SumFuelRequirements(masses []int) int {
	sum := 0

	for _, mass := range masses {
		sum += CalcFuelRequirement(mass)
	}

	return sum
}