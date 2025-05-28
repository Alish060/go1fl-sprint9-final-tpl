package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	slice := generateRandomElements(0)
	if len(slice) != 0 {
		t.Errorf("Expected size 0, got size %d", len(slice))
	}

	slice = generateRandomElements(-100)
	if len(slice) != 0 {
		t.Errorf("Expected size 0 , got size %d", len(slice))
	}

	size := 10
	slice = generateRandomElements(size)
	if len(slice) != size {
		t.Errorf("Expected slice size %d, got size %d", size, len(slice))
	}
}

func TestMaximum(t *testing.T) {

	if max := maximum([]int{}); max != 0 {
		t.Errorf("Expected 0 for empty slice, got %d", max)
	}

	if max := maximum([]int{42}); max != 42 {
		t.Errorf("Expected 42, got %d", max)
	}

	data := []int{1, 5, 2, 9, 3}
	expected := 9
	if max := maximum(data); max != expected {
		t.Errorf("Expected %d, got %d", expected, max)
	}
}

func TestMaxChunks(t *testing.T) {

	if max := maxChunks([]int{}); max != 0 {
		t.Errorf("Expected 0 for empty slice, got %d", max)
	}

	if max := maxChunks([]int{42}); max != 0 {
		t.Errorf("Expected 42, got %d", max)
	}

	data := []int{1, 8, 3, 4, 5, 6, 7, 2}
	expected := 8
	if max := maxChunks(data); max != expected {
		t.Errorf("Expected %d, got %d", expected, max)
	}

	slice := generateRandomElements(100)
	max1 := maximum(slice)
	max2 := maxChunks(slice)
	if max1 != max2 {
		t.Errorf("Expected same max: got %d  %d", max1, max2)
	}
}
