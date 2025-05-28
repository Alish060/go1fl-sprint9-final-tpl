package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoolSize(t *testing.T) {
	expectedError := errors.New("некорреткный размер")
	expectedValue := []int{}
	data, err := generateRandomElements(0)
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
		assert.Equal(t, expectedValue, data)
	}

}

func TestEmptySlice(t *testing.T) {
	expectedError := errors.New("слайс пустой")
	expectedValue := 0
	data, err := maximum([]int{})
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
		assert.Equal(t, expectedValue, data)
	}
}
