package algorithm

//You are given two non-empty linked lists representing two non-negative integer
//s. The digits are stored in reverse order, and each of their nodes contains a si
//ngle digit. Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the nu
//mber 0 itself.
//
//
// Example 1:
//
//
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.
//
//
// Example 2:
//
//
//Input: l1 = [0], l2 = [0]
//Output: [0]
//
//
// Example 3:
//
//
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]
//
//
//
// Constraints:
//
//
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading
// zeros.
//
// Related Topics Linked List Math Recursion
// 👍 11160 👎 2664

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	index_1 := l1
	index_2 := l2

	temp := 0
	var head *ListNode
	current := &ListNode{}

	for index_1 != nil && index_2 != nil {
		sum := temp + index_1.Val + index_2.Val
		temp = sum / 10
		value := sum % 10

		new_node := &ListNode{
			Val: value,
		}

		// save return point
		if head == nil {
			head = new_node
		}

		current.Next = new_node
		current = new_node

		index_1 = index_1.Next
		index_2 = index_2.Next
	}

	for index_1 != nil {
		sum := temp + index_1.Val
		temp = sum / 10
		value := sum % 10

		new_node := &ListNode{
			Val: value,
		}

		// save return point
		if head == nil {
			head = new_node
		}

		current.Next = new_node
		current = new_node

		index_1 = index_1.Next
	}

	for index_2 != nil {
		sum := temp + index_2.Val
		temp = sum / 10
		value := sum % 10

		new_node := &ListNode{
			Val: value,
		}

		// save return point
		if head == nil {
			head = new_node
		}

		current.Next = new_node
		current = new_node

		index_2 = index_2.Next

	}

	// easy to forget about
	if temp != 0 {
		new_node := &ListNode{
			Val: temp,
		}
		current.Next = new_node
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
