package main

import (
	"reflect"
	"testing"
)

func TestIntCode(t *testing.T) {
	checkRequirements := func(t *testing.T, program []int, want []int) {
		t.Helper()

		got := RunIntCode(program)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Should support add operation", func(t *testing.T) {
		program := []int{1, 0, 0, 0, 99}

		checkRequirements(t, program, []int{2, 0, 0, 0, 99})
	})

	t.Run("Should support add operation - 2", func(t *testing.T) {
		program := []int{1, 2, 1, 0, 99}

		checkRequirements(t, program, []int{3, 2, 1, 0, 99})
	})

	t.Run("Should support add operation - 3", func(t *testing.T) {
		program := []int{1, 5, 6, 0, 99, 30, 40}

		checkRequirements(t, program, []int{70, 5, 6, 0, 99, 30, 40})
	})

	t.Run("Should support multiply operation", func(t *testing.T) {
		program := []int{2, 5, 6, 0, 99, 10, 20}

		checkRequirements(t, program, []int{200, 5, 6, 0, 99, 10, 20})
	})

	t.Run("Should support multiply operation - 2", func(t *testing.T) {
		program := []int{2, 5, 6, 0, 99, 30, 40}

		checkRequirements(t, program, []int{1200, 5, 6, 0, 99, 30, 40})
	})

	t.Run("Should step through program and halt", func(t *testing.T) {
		program := []int{1, 0, 0, 0, 99}

		checkRequirements(t, program, []int{2, 0, 0, 0, 99})
	})

	t.Run("Test case 1", func(t *testing.T) {
		program := []int{2, 3, 0, 3, 99}

		checkRequirements(t, program, []int{2, 3, 0, 6, 99})
	})

	t.Run("Test case 2", func(t *testing.T) {
		program := []int{2, 4, 4, 5, 99, 0}

		checkRequirements(t, program, []int{2, 4, 4, 5, 99, 9801})
	})

	t.Run("Test case 3", func(t *testing.T) {
		program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}

		checkRequirements(t, program, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
	})
}
