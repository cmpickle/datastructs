package search

import (
	"testing"
)

func TestSearchArrayFirstElement(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 0
	expected := 0
	result := Binary(arr, target)
	if result != expected {
		t.Errorf("First element search failed, expected %d but got %d", expected, result)
	} else {
		t.Logf("First element search success, expected %d and got %d", expected, result)
	}
}

func TestSearchArrayMiddleElement(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 4
	expected := 4
	result := Binary(arr, target)
	if result != expected {
		t.Errorf("First element search failed, expected %d but got %d", expected, result)
	} else {
		t.Logf("First element search success, expected %d and got %d", expected, result)
	}
}

func TestSearchArrayEndElement(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 9
	expected := 9
	result := Binary(arr, target)
	if result != expected {
		t.Errorf("First element search failed, expected %d but got %d", expected, result)
	} else {
		t.Logf("First element search success, expected %d and got %d", expected, result)
	}
}

func TestSearchArrayNegativeTarget(t *testing.T) {
	arr := []int{-30, -15, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 42, 100}
	target := -15
	expected := 1
	result := Binary(arr, target)
	if result != expected {
		t.Errorf("First element search failed, expected %d but got %d", expected, result)
	} else {
		t.Logf("First element search success, expected %d and got %d", expected, result)
	}
}

func TestSearchArrayZeroTarget(t *testing.T) {
	arr := []int{-30, -15, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 42, 100}
	target := 0
	expected := 3
	result := Binary(arr, target)
	if result != expected {
		t.Errorf("First element search failed, expected %d but got %d", expected, result)
	} else {
		t.Logf("First element search success, expected %d and got %d", expected, result)
	}
}

func TestSearchArrayLargeTarget(t *testing.T) {
	arr := []int{-30, -15, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 42, 100000000}
	target := 100000000
	expected := 14
	result := Binary(arr, target)
	if result != expected {
		t.Errorf("First element search failed, expected %d but got %d", expected, result)
	} else {
		t.Logf("First element search success, expected %d and got %d", expected, result)
	}
}
