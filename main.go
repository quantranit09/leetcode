package main

import (
	"fmt"
	closestsubsequencesum "leetcode/solutions/1755_closest_subsequence_sum"
)

func main() {
	// 0001 Two Sum
	// result := twosum.TwoSumEnhance([]int{1, 2, 4}, 6)
	// fmt.Println("Result from main:", result)

	// 1755 Closest Subsequence Sum
	nums := []int{1556913,-259675,-7667451,-4380629,-4643857,-1436369,7695949,-4357992,-842512,-118463}
	goal := -9681425
	resultClosest := closestsubsequencesum.MinAbsDifference(nums, goal)
	fmt.Println("Result from MinAbsDifference:", resultClosest)
}
