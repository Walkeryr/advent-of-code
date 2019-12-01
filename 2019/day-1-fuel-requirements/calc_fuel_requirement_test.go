package main

import "testing"

func TestCalcFuelRequirements(t *testing.T) {
	checkRequirements := func(t *testing.T, mass int, want int) {
		t.Helper()

		got := CalcFuelRequirement(mass)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Test case 1", func(t *testing.T) {
		mass := 12

		checkRequirements(t, mass, 2)
	})

	t.Run("Test case 2", func(t *testing.T) {
		mass := 14

		checkRequirements(t, mass, 2)
	})

	t.Run("Test case 3", func(t *testing.T) {
		mass := 1969

		checkRequirements(t, mass, 966)
	})

	t.Run("Test case 4", func(t *testing.T) {
		mass := 100756

		checkRequirements(t, mass, 50346)
	})
}

// TODO: mock CalcFuelRequirement function and test only main functionality
// of the function SumFuelRequirements.
func TestSumFuelRequirements(t *testing.T) {
	want := 970
	got := SumFuelRequirements([]int{12, 14, 1969})

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}