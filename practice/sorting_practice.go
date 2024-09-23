package practice

import "fmt"

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	fmt.Println("Runs")
}

func selection_sort(arr []int) []int {
	for i := 0; i <= len(arr)-2; i++ {
		mini := i
		fmt.Println(arr[i], arr[i+1], arr)
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[j] < arr[mini] {
				mini = j

			}
		}
		swap(arr, i, mini)
	}
	return arr
}

func bubble_sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		didSwap := 0
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
				didSwap = 1
			}
		}
		if didSwap == 0 {
			break
		}
	}
	return arr
}

func insertion_sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			swap(arr, j, j-1)
			j--
		}
	}
	return arr
}

func merge_sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := merge_sort(arr[:mid])
	right := merge_sort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

func quick_sort(arr []int, low, high int) []int {

	if low < high {
		index := partition(arr, low, high)
		quick_sort(arr, low, index-1)
		quick_sort(arr, index+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high
	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}
		for arr[j] > pivot && j < low {
			j--
		}
		if i < j {
			swap(arr, i, j)
		}
	}
	swap(arr, low, j)
	return j

}

func SortingPractice() {
	ans := merge_sort([]int{3, 5, 1, 4, 2})
	fmt.Println("Ans is  => ", ans)
}
