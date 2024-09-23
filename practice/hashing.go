package practice

import "fmt"

func duplicate_count(arr []int, search int) int {
	hash := map[int]int{} // in Go int is by default starts from 0

	for _, val := range arr {
		hash[val] += 1
	}

	return hash[search]

}

func max_low_frequency(arr []int) []int {
	hash := map[int]int{}
	maxF := 0
	elMx := 0
	elMn := 0
	minF := len(arr)
	for _, v := range arr {
		if val, ok := hash[v]; ok {
			hash[v] = val + 1
		} else {
			hash[v] = 1
		}
	}
	for k, v := range hash {
		if v > maxF {
			maxF = v
			elMx = k
		}

		if v < minF {
			minF = v
			elMn = k
		}
	}

	return []int{elMx, elMn}

}

func HasingPractice() {
	ans := max_low_frequency([]int{1, 2, 1, 2, 5, 7, 2, 5})
	fmt.Println("Ans is  => ", ans)
}
