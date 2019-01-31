package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{2, 3, 4, 6, 8}
		got := Sum(numbers)
		want := 23

		if got != want {
			t.Errorf("want '%d' got '%d' given %v", want, got, numbers)
		}
	})
	t.Run("Collection of any sizes", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("want '%d' got '%d' given %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 3}, []int{2, 5})
	want := []int{5, 7}

	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
