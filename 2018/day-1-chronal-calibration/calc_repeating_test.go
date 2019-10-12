package main

import "testing"

func TestCalcRepeating(t *testing.T) {
	checkRepeating := func(t *testing.T, frequencies []int, want int) {
		t.Helper()

		got := CalcRepeatingFrequency(frequencies)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Test case 1", func(t *testing.T) {
		frequencies := []int{+1, -1}

		checkRepeating(t, frequencies, 0)
	})

	t.Run("Test case 2", func(t *testing.T) {
		frequencies := []int{+3, +3, +4, -2, -4}

		checkRepeating(t, frequencies, 10)
	})

	t.Run("Test case 3", func(t *testing.T) {
		frequencies := []int{-6, +3, +8, +5, -6}

		checkRepeating(t, frequencies, 5)
	})

	t.Run("Test case 4", func(t *testing.T) {
		frequencies := []int{+7, +7, -2, -7, -4}

		checkRepeating(t, frequencies, 14)
	})
}
