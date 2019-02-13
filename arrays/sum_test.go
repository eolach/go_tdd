package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3, 5}

		got := Sum(numbers)
		want := 11

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{3, 5})
	want := []int{3, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func Test(t *testing.T) {
	got := SumAllTails([]int{5, 6}, []int{3, 5},
		[]int{1, 2}, []int{3, 5})
	want := []int{7, 12}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
