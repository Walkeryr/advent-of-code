package main

import "testing"

func TestCalcFrequency(t *testing.T) {
	checkFrequencies := func(t *testing.T, frequencies []int, want int) {
		t.Helper()

		got := CalcResultFrequency(frequencies)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Test case 1", func(t *testing.T) {
		frequencies := []int{1, -2, +3, 1}

		checkFrequencies(t, frequencies, 3)
	})

	t.Run("Test case 2", func(t *testing.T) {
		frequencies := []int{1, 1, 1}

		checkFrequencies(t, frequencies, 3)
	})

	t.Run("Test case 3", func(t *testing.T) {
		frequencies := []int{1, 1, -2}

		checkFrequencies(t, frequencies, 0)
	})

	t.Run("Test case 4", func(t *testing.T) {
		frequencies := []int{-1, -2, -3}

		checkFrequencies(t, frequencies, -6)
	})
}
