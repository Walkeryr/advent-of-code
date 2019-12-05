package main

import "fmt"

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

