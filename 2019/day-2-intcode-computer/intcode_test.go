package main

import (
	"testing"
	"reflect"
)

func TestIntCode(t *testing.T) {
	checkIntCode := func(t *testing.T, program []int, want []int) {
		t.Helper()

		got := RunIntCode(program)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Should handle adding", func(t *testing.T) {
		program := []int{1, 10, 20, 0}

		checkIntCode(t, program, []int{30, 10, 20, 0})
	})
}
