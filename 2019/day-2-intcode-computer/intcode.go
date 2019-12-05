package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const haltCode = 99
const addCode = 1
const multiplyCode = 2

// RunIntCode - runs until halt command and returns final program state
func RunIntCode(program []int) []int {
	i := 0

	for {
		inputCode := program[i]

		if inputCode == haltCode {
			break
		}

		outputPosition := program[i + 3]
		firstOperand := program[program[i + 1]]
		secondOperand := program[program[i + 2]]

		if inputCode == addCode {
			program[outputPosition] = firstOperand + secondOperand
		} else if inputCode == multiplyCode {
			program[outputPosition] = firstOperand * secondOperand
		} else {
			panic(fmt.Sprintf("Error, unhandled opcode: %v", inputCode))
		}

		i += 4
	}

	return program
}

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

	program[1] = 12
	program[2] = 2

	fmt.Fprintf(os.Stdout, "Output result of RunIntCode: \n%v\n", RunIntCode(program))
}
