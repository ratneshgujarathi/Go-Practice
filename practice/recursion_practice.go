package practice

import "fmt"

func recur_num_asc(num1 int, num2 int) {
	if num1 > num2 {
		return
	}
	fmt.Println(num1)
	recur_num_asc(num1+1, num2)
}

func recur_num_desc(num1 int, num2 int) {
	if num1 < num2 {
		return
	}
	fmt.Println(num1)
	recur_num_asc(num1-1, num2)
}

func recur_sum(num int) int {
	//return num * ((num + 1) / 2) //optimal sol
	if num == 0 {
		return 0
	}
	return num + recur_sum(num-1)
}

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}

func reverse_arr(arr []int) []int {
	// two pointer method
	p1 := 0
	p2 := len(arr) - 1

	for p1 < p2 {
		arr[p1], arr[p2] = arr[p2], arr[p1]
		p1 += 1
		p2 -= 1
	}
	return arr
}

func recur_reverse_arr(arr []int, start int, end int) []int {
	if start < end {
		arr[start], arr[end] = arr[end], arr[start]
		recur_reverse_arr(arr, start+1, end-1)
	}
	return arr
}

func check_str_palindrome(s string) bool {
	// two pointer
	chrs := []rune(s)
	l := len(chrs)
	for i := 0; i < l/2; i++ {
		chrs[i], chrs[l-i-1] = chrs[l-i-1], chrs[i]
	}
	return string(chrs) == s
}

func check_recur_str_palindrome(i int, s string) bool {
	chrs := []rune(s)
	l := len(chrs)
	if i >= l/2 {
		return true
	}
	if chrs[i] != chrs[l-i-1] {
		return false
	}
	return check_recur_str_palindrome(i+1, string(s))
}

func fib(num int) int {
	if num <= 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}

func fib_series(num int) []int {
	series := make([]int, num)
	for i := 0; i < num; i++ {
		series[i] = fib(i)
	}
	return series
}

func RecursionPractice() {
	ans := recur_sum(5)
	fmt.Println("Ans is => ", ans)
}
