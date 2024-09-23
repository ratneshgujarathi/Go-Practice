package practice

import (
	"fmt"
	"math"
	"slices"
)

func extract_nums(num int) []int {
	ans := []int{}
	for num != 0 {
		ans = append(ans, num%10)
		num /= 10
	}
	return ans
}

func reverse_num(a int) int {
	// reverse a number
	ans := 0
	for a != 0 {
		ans = ans*10 + a%10
		a /= 10
	}

	return ans
}

func check_palindrome(num int) bool {
	// check palindrome
	onum := num
	dup := 0
	for num != 0 {
		dup = dup*10 + num%10
		num /= 10
	}
	return dup == onum
}

func count_nums(num int) int {
	// count nums in optimal way
	count := 0
	count = int(math.Log10(float64(num))) + 1
	return count
}

func check_amrstrong(num int) bool {
	original := num
	sum := 0
	count := count_nums(num)
	for num != 0 {
		rem := num % 10
		sum += int(math.Pow(float64(rem), float64(count)))
		num /= 10
	}

	return original == sum
}

func all_divisors_brute(num int) []int {
	ans := []int{}
	for i := 1; i < num; i++ {
		if num%i == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func all_divisors(num int) []int {
	ans := []int{}
	for i := 1; i*i < num; i++ {
		if num%i == 0 {
			ans = append(ans, i)
			if num/i != i {
				ans = append(ans, num/i)
			}
		}
	}
	slices.Sort(ans)
	return ans
}

func check_prime_num(num int) bool {
	cnt := 0
	for i := 1; i*i < num; i++ {
		if num%i == 0 {
			cnt += 1
			if num/i != i {
				cnt += 1
			}
		}
	}
	return cnt == 2
}

func find_GCD(num1 int, num2 int) int {
	for num1 > 0 && num2 > 0 {
		if num1 > num2 {
			num1 %= num2
		} else {
			num2 %= num1
		}
	}
	if num1 == 0 {
		return num2
	} else {
		return num1
	}
}

func MathPractice() {
	ans := find_GCD(10, 52)
	fmt.Println("Ans is => ", ans)
}
