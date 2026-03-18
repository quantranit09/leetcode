package closestsubsequencesum

import (
	"fmt"
	"sort"
	"math"
)

/*
You are given an integer array nums and an integer goal.

You want to choose a subsequence of nums such that the sum of its elements is the closest possible to goal. That is, if the sum of the subsequence's elements is sum, then you want to minimize the absolute difference abs(sum - goal).

Return the minimum possible value of abs(sum - goal).

Note that a subsequence of an array is an array formed by removing some elements (possibly all or none) of the original array.



Example 1:

Input: nums = [5,-7,3,5], goal = 6
Output: 0
Explanation: Choose the whole array as a subsequence, with a sum of 6.
This is equal to the goal, so the absolute difference is 0.
Example 2:

Input: nums = [7,-9,15,-2], goal = -5
Output: 1
Explanation: Choose the subsequence [7,-9,-2], with a sum of -4.
The absolute difference is abs(-4 - (-5)) = abs(1) = 1, which is the minimum.
Example 3:

Input: nums = [1,2,3], goal = -7
Output: 7


Constraints:

1 <= nums.length <= 40
-107 <= nums[i] <= 107
-109 <= goal <= 109

*/

func MinAbsDifference(nums []int, goal int) int {
	sortDirection := "desc"
	if goal < 0 {
		sortDirection = "asc"
	}
	sort.Slice(nums, func(i, j int) bool {
		if sortDirection == "asc" {
			return nums[i] < nums[j]
		}
		return nums[i] > nums[j]
	})
	subsequence := []int{}
	lastAbsDiff := int(math.Abs(float64(0 - goal)))

	for i := 0; i < len(nums); i++ {
		subsequence = append(subsequence, nums[i])

		sumOfSubsequence := 0
		for j := 0; j < len(subsequence); j++ {
			sumOfSubsequence += subsequence[j]
		}

		fmt.Println("sumOfSubsequence", subsequence, sumOfSubsequence)
		absDiff := int(math.Abs(float64(sumOfSubsequence - goal)))
		if absDiff == 0 {
			return 0
		}
	
		if absDiff < lastAbsDiff {
			lastAbsDiff = int(absDiff)
		}
		fmt.Println("lastAbsDiff", lastAbsDiff, "new absDiff", absDiff)
	}

	return lastAbsDiff
}
