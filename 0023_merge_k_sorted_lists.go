package algorithm

import "math"

//You are given an array of k linked-lists lists, each linked-list is sorted in
//ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
//
// Example 1:
//
//
//Input: lists = [[1,4,5],[1,3,4],[2,6]]
//Output: [1,1,2,3,4,4,5,6]
//Explanation: The linked-lists are:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//merging them into one sorted list:
//1->1->2->3->4->4->5->6
//
//
// Example 2:
//
//
//Input: lists = []
//Output: []
//
//
// Example 3:
//
//
//Input: lists = [[]]
//Output: []
//
//
//
// Constraints:
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] is sorted in ascending order.
// The sum of lists[i].length won't exceed 10^4.
//
// Related Topics Linked List Divide and Conquer Heap
// 👍 6551 👎 343

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	size := len(lists)
	headLists := make([]*ListNode, size)

	for index, list := range lists {
		headLists[index] = list
	}

	head := &ListNode{}
	current := head

	for {
		index := findMin(headLists)
		if index == -1 {
			break
		}

		value := headLists[index].Val
		headLists[index] = headLists[index].Next
		newNode := &ListNode{Val: value}
		current.Next = newNode
		current = newNode

	}

	return head.Next

}

func findMin(lists []*ListNode) int {
	min := math.MaxInt32
	position := -1
	for index, list := range lists {
		if list != nil {
			if list.Val < min {
				position = index
				min = list.Val
			}
		}
	}

	return position
}

//leetcode submit region end(Prohibit modification and deletion)
