package algorithm

//Given a binary tree, check whether it is a mirror of itself (ie, symmetric aro
//und its center).
//
// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//
//
//
//
// But the following [1,2,2,null,3,null,3] is not:
//
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//
//
//
//
// Follow up: Solve it both recursively and iteratively.
// Related Topics Tree Depth-first Search Breadth-first Search
// 👍 4905 👎 117

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	level := []*TreeNode{root}
	for len(level) > 0 {
		var next []*TreeNode
		for _, l := range level {
			if l != nil {
				next = append(next, l.Left)
				next = append(next, l.Right)
			}
		}
		if !symmetric(next) {
			return false
		}

		level = next
	}

	return true
}

func symmetric(nodes []*TreeNode) bool {
	size := len(nodes)
	if size%2 != 0 {
		return false
	}

	half := size / 2
	for i := 0; i < half; i++ {
		me := nodes[i]
		bro := nodes[size-i-1]
		if me == nil && bro == nil {
			continue
		} else if me != nil && bro != nil && me.Val == bro.Val {
			continue
		} else {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
