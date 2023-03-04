package main

import (
	"fmt"
	"strings"
)

func main() {
	r := []int{4, 5, 3, 0, 1, 2}
	mergeSort(r, 0, 4)
	fmt.Println(strings.Trim(fmt.Sprint(r), "[]") ==
		"0 3 4 5 1 2")

}

// merge - принимает два отсортированных массива, сливает их
// в один отсортированный массив и возвращает его.
// Первый массив задаётся полуинтервалом [left, mid) массива
// array, а второй - полуинтервалом [mid, right) массива array
func merge(arr []int, lf int, mid int, rg int) []int {
	res := make([]int, -lf+rg)
	i, j, k := lf, mid, 0
	for i < mid && j < rg {
		if arr[i] <= arr[j] {
			res[k] = arr[i]
			i, k = i+1, k+1
		} else {
			res[k] = arr[j]
			j, k = j+1, k+1
		}
	}
	for i < mid {
		res[k] = arr[i]
		i, k = i+1, k+1
	}
	for j < rg {
		res[k] = arr[j]
		j, k = j+1, k+1
	}

	return res
}

// mergeSort - разбивает полуинтервал на две половинки
// и рекурсивно вызывает сортировку отдельно для каждой
func mergeSort(arr []int, lf int, rg int) {
	if rg-lf == 1 {
		return
	}

	mergeSort(arr, lf, (lf+rg+1)/2)
	mergeSort(arr, (lf+rg+1)/2, rg)

	res := merge(arr, lf, (lf+rg+1)/2, rg)
	for i := lf; i < rg; i++ {
		arr[i] = res[i-lf]
	}
}
