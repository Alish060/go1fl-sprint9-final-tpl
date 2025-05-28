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
func generateRandomElements(size int) ([]int, error) {
	if size <= 0 {
		return []int{}, errors.New("некорреткный размер")
	}

	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Int()
	}

	return slice, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("слайс пустой")
	}

	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max, nil

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	// ваш код здесь
	if len(data) <= 1 {
		return 0, errors.New("слайс пустой")
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	maxSlice := []int{}
	for i := 0; i < CHUNKS; i++ {
		subSliceSize := len(data) / CHUNKS
		start := i * subSliceSize
		end := start + subSliceSize
		wg.Add(1)
		go func(subLice []int) {
			defer wg.Done()
			max := subLice[0]
			for _, v := range subLice {
				if v > max {
					max = v
				}
			}
			mu.Lock()
			maxSlice = append(maxSlice, max)
			mu.Unlock()
		}(data[start:end])
	}
	wg.Wait()
	result, err := maximum(maxSlice)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data, err := generateRandomElements(SIZE)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max, err := maximum(data)
	if err != nil {
		fmt.Print(err.Error())
	}
	elapsed := time.Duration(time.Since(start).Microseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max, err = maxChunks(data)
	if err != nil {
		fmt.Print(err.Error())
	}
	elapsed = time.Duration(time.Since(start).Microseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
