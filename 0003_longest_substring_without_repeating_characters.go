package algorithm

//Given a string s, find the length of the longest substring without repeating c
//haracters.
//
//
// Example 1:
//
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//
//
// Example 2:
//
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//
//
// Example 3:
//
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a
//substring.
//
//
// Example 4:
//
//
//Input: s = ""
//Output: 0
//
//
//
// Constraints:
//
//
// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
//
// Related Topics Hash Table Two Pointers String Sliding Window
// 👍 13081 👎 679

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	longest := 0
	size := len(s)

	for i := 0; i < size; i++ {
		m := make(map[byte]bool, 0)
		l := 0
		for j := i; j < size; j++ {
			_, exist := m[s[j]]
			if exist {
				break
			} else {
				m[s[j]] = true
				l++
			}
		}
		if l > longest {
			longest = l
		}
	}

	return longest
}

//leetcode submit region end(Prohibit modification and deletion)
