//Собраны различные алгоритмы сортировки и проанализировано время выполнения кода
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {

}

func main() {
	arr := createArr()
	emptyArr := make([]int, len(arr))

	copy(emptyArr, arr)
	start := time.Now()
	bubbleSort(emptyArr) // 1. Сортировка пузырьком
	duration := time.Since(start)
	fmt.Println("Cортировка пузырьком - O(n^2) занимает:", duration)

	copy(emptyArr, arr) // 2. Сортировка выбором
	start = time.Now()
	selectionSort(emptyArr)
	duration = time.Since(start)
	fmt.Println("Cортировка выбором/двоичный поиск - O(n^2) занимает:", duration)

	copy(emptyArr, arr) // 3. Сортировка вставками
	start = time.Now()
	insertionSort(emptyArr)
	duration = time.Since(start)
	fmt.Println("Cортировка вставками O(n^2) занимает:", duration)

	copy(emptyArr, arr)
	start = time.Now()
	quickSort(emptyArr) // 4. Быстрая сортировка
	duration = time.Since(start)
	fmt.Println("Быстрая сортировка - O(n x log n) занимает:", duration)

	copy(emptyArr, arr) // 5. Сортировка слиянием
	start = time.Now()
	mergeSort(emptyArr)
	duration = time.Since(start)
	fmt.Println("Сортировка слиянием - O(n x log n) занимает:", duration)

	copy(emptyArr, arr) // 6. Встроенная в Go сортировка
	start = time.Now()
	sort.Ints(emptyArr)
	duration = time.Since(start)
	fmt.Println("Встроенная в Go - O(n x log n) занимает:", duration)
}

func createArr() (arr []int) {
	sizeArr := 100000
	arr = make([]int, sizeArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)
	}
	return
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}
