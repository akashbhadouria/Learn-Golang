package main

import "testing"

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4}

	got := Sum(numbers)
	want := 10

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}