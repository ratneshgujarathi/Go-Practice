package practice

import "fmt"

// func reverse(s string, start, end int) string{
//     ans:=""
// 	for start < end{

// 	}
// 	return ans
// }

// func reverseWords(s string) string {
//     ans:=""
//     breakpoint:=0
//     for i, val := range s{
//         if val == ' '{
//             ans += reverse(s, breakpoint, i)
//             breakpoint = i+1
//         }
//     }
// }

func generate(numRows int) [][]int {
	var triangle [][]int
	if numRows == 0 {
		return triangle
	}

	triangle = append(triangle, []int{1})

	for i := 1; i < numRows; i++ {
		prevRow := triangle[i-1]
		var newRow []int
		newRow = append(newRow, 1)
		fmt.Println(triangle, i, prevRow)

		for j := 1; j < len(prevRow); j++ {
			fmt.Println("in j => ", j)
			newRow = append(newRow, prevRow[j-1]+prevRow[j])
		}

		newRow = append(newRow, 1)
		triangle = append(triangle, newRow)
	}

	return triangle
}

func StringPractice() {
	ans := generate(5)
	fmt.Println("Ans is => ", ans)
}
