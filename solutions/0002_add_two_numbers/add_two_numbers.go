package addtwonumbers

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/** helpers */
func ConvertArrayToList(arr []int) *ListNode {
	list := &ListNode{}
	if len(arr) == 0 {
		return nil
	}

	list.Val = arr[0]
	list.Next = ConvertArrayToList(arr[1:])

	return list
}

// approach 1: convert to int, add, then convert back to list
// func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	sumOfLists := getSumOfList(l1) + getSumOfList(l2)
// 	fmt.Println("sumOfLists", sumOfLists)
// 	return convertIntToList(sumOfLists)
// }

// func getSumOfList(l *ListNode) int {
// 	sumOfList := 0
// 	currentIndex := 0
// 	for {
// 		sumOfList += l.Val * int(math.Pow(10, float64(currentIndex)))
// 		fmt.Println("Debug--->", sumOfList, l.Val, currentIndex, int(math.Pow(10, float64(currentIndex))))
// 		currentIndex++
// 		if l.Next == nil {
// 			break
// 		}
// 		l = l.Next
// 	}

// 	return sumOfList
// }

// func convertIntToList(n int) *ListNode {
// 	list := &ListNode{}
// 	if n == 0 {
// 		return list
// 	}

// 	list.Val = n % 10
// 	nextInt := n / 10

// 	if nextInt == 0 {
// 		return list
// 	}

// 	list.Next = convertIntToList(nextInt)

// 	return list
// }

// approach 2: add two numbers directly from linked list
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	result := &ListNode{}
	currentNode := result
	for {
		sumValue := l1.Val + l2.Val + carry
		digit := sumValue % 10
		carry = sumValue / 10

		currentNode.Val = digit
		if l1.Next == nil && l2.Next == nil && carry == 0 {
			break
		}
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}
		if l2.Next == nil {
			l2.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
	}

	return result
}