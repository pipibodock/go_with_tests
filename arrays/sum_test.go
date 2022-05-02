package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertValuesMessage := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}
	t.Run("sum all values", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertValuesMessage(t, got, want, numbers)
	})
	t.Run("sum any size array values", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertValuesMessage(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	assertValuesMessage := func(t testing.TB, got, want[]int) {
		t.Helper()
		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
		}
		for i:= range got {
			if got[i] != want[i] {
				t.Errorf("got %v want %v", got, want)
				break
			}
		}
	}
	t.Run("Summ all should return an array of values", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{3, 9}
		got := SumAll(numbers1, numbers2)
		want := []int{3, 12}
		assertValuesMessage(t, got, want)
	})
	
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of tails", func(t *testing.T){
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("safely sum tails of empty slices", func(t *testing.T){
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}