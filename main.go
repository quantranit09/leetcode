package main

import (
	"fmt"
	// closestsubsequencesum "leetcode/solutions/1755_closest_subsequence_sum"
	addtwonumbers "leetcode/solutions/0002_add_two_numbers"
)

func main() {
	// 0001 Two Sum
	// result := twosum.TwoSumEnhance([]int{1, 2, 4}, 6)
	// fmt.Println("Result from main:", result)

	// 1755 Closest Subsequence Sum
	// nums := []int{1556913,-259675,-7667451,-4380629,-4643857,-1436369,7695949,-4357992,-842512,-118463}
	// goal := -9681425
	// resultClosest := closestsubsequencesum.MinAbsDifference(nums, goal)
	// fmt.Println("Result from MinAbsDifference:", resultClosest)

	// 0002 Add Two Numbers
	// l1Array := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	// l2Array := []int{5, 6, 4}

	// l1Array := []int{2, 4, 3}
	// l2Array := []int{5, 6, 4}

	l1Array := []int{9, 9, 9, 9, 9, 9, 9}
	l2Array := []int{9, 9, 9, 9}

	l1 := addtwonumbers.ConvertArrayToList(l1Array)
	l2 := addtwonumbers.ConvertArrayToList(l2Array)
	fmt.Println("l1", l1)
	fmt.Println("l2", l2)
	resultAdd := addtwonumbers.AddTwoNumbers(l1, l2)
	fmt.Println("Result from AddTwoNumbers:", resultAdd)
}
