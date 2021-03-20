package algorithm

//Given a string s, return the longest palindromic substring in s.
//
//
// Example 1:
//
//
//Input: s = "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
//
//
// Example 2:
//
//
//Input: s = "cbbd"
//Output: "bb"
//
//
// Example 3:
//
//
//Input: s = "a"
//Output: "a"
//
//
// Example 4:
//
//
//Input: s = "ac"
//Output: "a"
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 1000
// s consist of only digits and English letters (lower-case and/or upper-case),
//
// Related Topics String Dynamic Programming
// 👍 10051 👎 663

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	m := make(map[string]bool)
	size := len(s)
	longest := 0
	start := 0
	end := 0
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			_, exist := m[s[i:j+1]]
			if !exist {
				if isPalindrome(s[i : j+1]) {
					m[s[i:j+1]] = true
					currentLength := j - i + 1
					if currentLength > longest {
						longest = currentLength
						start = i
						end = j + 1
					}
				}
			}

		}
	}
	return s[start:end]
}

func isPalindrome(s string) bool {
	size := len(s)
	half := size / 2
	for i := 0; i < half; i++ {
		if s[i] != s[size-1-i] {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
