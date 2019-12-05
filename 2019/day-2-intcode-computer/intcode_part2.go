package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

func main() {
	var program []int
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		programRaw := strings.Split(line, ",")

		for _, code := range programRaw {
			number, err := strconv.Atoi(code)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error converting %s to integer number", line)

				return
			}

			program = append(program, number)
		}
	}

	noun, verb, err := findParameters(program)

	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)

		return
	}

	fmt.Fprintf(os.Stdout, "What is 100 * noun + verb?: %v", 100 * noun + verb)
}

func findParameters(program []int) (int, int, error) {
	const targetOutput = 19690720

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			fmt.Printf("noun %d, verb %d\n", noun, verb)

			modifiedProgram := make([]int, len(program))
			copy(modifiedProgram, program)

			modifiedProgram[1] = noun
			modifiedProgram[2] = verb

			output := RunIntCode(modifiedProgram)

			if output[0] == targetOutput {
				return noun, verb, nil
			}
		}
	}

	return 0, 0, errors.New("Noun and verb not found for target output")
}
