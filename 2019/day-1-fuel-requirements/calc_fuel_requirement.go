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

// CalcFuelRequirement - calculate the total fuel required to launch given module together with all fuel
// that would be used during travel.
func CalcFuelRequirement(mass int) int {
	requiredFuel := int(math.Floor(float64(mass) / 3.0) - 2)

	if (requiredFuel <= 0) {
		return 0
	}

	return requiredFuel + CalcFuelRequirement(requiredFuel)
}

// SumFuelRequirements - calculate resulting sum of all fuel requirements
func SumFuelRequirements(masses []int) int {
	sum := 0

	for _, mass := range masses {
		sum += CalcFuelRequirement(mass)
	}

	return sum
}