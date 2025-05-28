package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"math/rand"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Print(errors.New("incorrect size"))
		return []int{}
	}

	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Int()
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		fmt.Print(errors.New("empty slice"))
		return 0
	}

	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) <= 1 {
		fmt.Print(errors.New("empty slice"))
		return 0
	}
	var wg sync.WaitGroup
	maxSlice := make([]int, CHUNKS)
	subSliceSize := len(data) / CHUNKS
	remainder := len(data) % CHUNKS
	start := 0
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		end := start + subSliceSize
		if i < remainder {
			end++
		}
		go func(subSlice []int, index int) {
			defer wg.Done()
			max := maximum(subSlice)
			maxSlice[index] = max
		}(data[start:end], i)
		start = end
	}
	wg.Wait()
	result := maximum(maxSlice)
	return result
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Duration(time.Since(start).Microseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Duration(time.Since(start).Microseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
