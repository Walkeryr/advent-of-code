package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

	program[1] = 12
	program[2] = 2

	fmt.Fprintf(os.Stdout, "Output result of RunIntCode: \n%v\n", RunIntCode(program))
}