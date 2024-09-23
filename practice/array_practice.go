package practice

import (
	"fmt"
	"math"
)

func find_s_small_large(arr []int) []int {
	small := arr[0]
	s_small := int(math.Inf(1))
	large := arr[0]
	s_large := -1
	for i := 1; i < len(arr); i++ {

		if arr[i] < small {
			s_small = small
			small = arr[i]
		} else if arr[i] < s_small && arr[i] > small {
			s_small = arr[i]
		}

		if arr[i] > large {
			s_large = large
			large = arr[i]
		} else if arr[i] > s_large && arr[i] < large {
			s_large = arr[i]
		}
	}

	return []int{s_small, s_large}
}

func check_sorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func check_sorted_rotated(nums []int) bool {
	count := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums, nums[i], nums[(i+1)%len(nums)])
		if nums[i] > nums[(i+1)%len(nums)] {
			count++
		}
	}
	return count <= 1
}

func remove_dup(arr []int) []int {
	// brute
	left := 0
	right := 1
	ans := []int{arr[left]}
	for left < right && right < len(arr) {
		if arr[left] == arr[right] {
			right++
		} else {
			arr[left+1] = arr[right]
			ans = append(ans, arr[right])
			left++
		}
	}
	fmt.Println(arr)
	return ans
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func rotate_by_n(arr []int, k int) []int {
	n := len(arr)
	k = k % n
	reverse(arr, 0, n-k-1)
	reverse(arr, n-k, n-1)
	reverse(arr, 0, n-1)
	return arr
}

func find_union(arr1, arr2 []int) []int {
	left, right := 0, 0
	l := len(arr1)
	r := len(arr2)
	ans := []int{}

	for left < l && right < r {
		if arr1[left] <= arr2[right] {
			if len(ans) == 0 || ans[len(ans)-1] != arr1[left] {
				ans = append(ans, arr1[left])
			}
			left++
		} else {
			if len(ans) == 0 || ans[len(ans)-1] != arr2[right] {
				ans = append(ans, arr2[right])
			}
			right++
		}
	}
	for left < l {
		if len(ans) == 0 || ans[len(ans)-1] != arr1[left] {
			ans = append(ans, arr1[left])
		}
		left++
	}
	for right < r {
		if len(ans) == 0 || ans[len(ans)-1] != arr2[right] {
			ans = append(ans, arr2[right])
		}
		right++
	}
	return ans
}

func missingNumber(nums []int) int {
	v := len(nums)
	for i, n := range nums {
		v ^= i
		v ^= n
	}
	return v
}

func lenOfLongSubArr(arr []int, length, target int) int {
	// if only positives are present
	left, right := 0, 0
	maxLen := 0
	sum := arr[0]
	for right < length {
		fmt.Println("sum => ", sum)
		for sum > target && left <= right {
			sum -= arr[left]
			left++
		}
		if sum == target {
			maxLen = max(maxLen, right-left+1)
		}
		right++
		if right < length {
			sum += arr[right]
		}

	}
	return maxLen
}

func lenOfLongSubArr1(arr []int, length, target int) int {
	hash := map[int]int{0: -1}
	maxLen, sum := 0, 0
	for i, j := range arr {
		sum += j
		if _, ok := hash[sum-target]; ok {
			maxLen = max(maxLen, i-hash[sum-target])
		}
		if _, ok := hash[sum]; !ok {
			hash[sum] = i
		}
	}
	return maxLen
}

func ArrayPractice() {
	// arr := []int{1, 2, 3, 4, 5, 6, 7}
	ans := lenOfLongSubArr1([]int{1, -1, 4, 3, 3, 5, 5}, 7, 7)
	fmt.Println("Ans is => ", ans)

}
