package algorithm

//Given two sorted arrays nums1 and nums2 of size m and n respectively, return t
//he median of the two sorted arrays.
//
// Follow up: The overall run time complexity should be O(log (m+n)).
//
//
// Example 1:
//
//
//Input: nums1 = [1,3], nums2 = [2]
//Output: 2.00000
//Explanation: merged array = [1,2,3] and median is 2.
//
//
// Example 2:
//
//
//Input: nums1 = [1,2], nums2 = [3,4]
//Output: 2.50000
//Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
//
//
// Example 3:
//
//
//Input: nums1 = [0,0], nums2 = [0,0]
//Output: 0.00000
//
//
// Example 4:
//
//
//Input: nums1 = [], nums2 = [1]
//Output: 1.00000
//
//
// Example 5:
//
//
//Input: nums1 = [2], nums2 = []
//Output: 2.00000
//
//
//
// Constraints:
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106
//
// Related Topics Array Binary Search Divide and Conquer
// 👍 9165 👎 1410

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var pass int
	m := len(nums1)
	n := len(nums2)

	total := m + n
	index_1 := 0
	index_2 := 0

	if total%2 == 0 {
		// total length is even, we want to middle two
		pass = total/2 - 1
		get_middle_num := 0
		middle_two := make([]int, 0)
		for index_1 < m && index_2 < n {
			if pass == 0 {
				if nums1[index_1] < nums2[index_2] {
					get_middle_num++
					middle_two = append(middle_two, nums1[index_1])
					index_1++
				} else {
					get_middle_num++
					middle_two = append(middle_two, nums2[index_2])
					index_2++
				}
				if get_middle_num == 2 {
					goto result
				}

			} else {
				if nums1[index_1] < nums2[index_2] {
					index_1++
				} else {
					index_2++
				}
				pass--
			}
		}

		for index_1 < m {
			if pass == 0 {
				get_middle_num++
				middle_two = append(middle_two, nums1[index_1])
				index_1++
				if get_middle_num == 2 {
					goto result
				}
			} else {
				index_1++
				pass--
			}
		}

		for index_2 < n {
			if pass == 0 {
				get_middle_num++
				middle_two = append(middle_two, nums2[index_2])
				index_2++
				if get_middle_num == 2 {
					goto result
				}
			} else {
				index_2++
				pass--
			}
		}

	result:
		median := (float64(middle_two[0]) + float64(middle_two[1])) / 2
		return median

	} else {
		// total length is odd, find the middle one
		pass = total / 2
		middle_num := 0

		for index_1 < m && index_2 < n {
			if pass == 0 {
				if nums1[index_1] < nums2[index_2] {
					middle_num = nums1[index_1]
					goto result2
				} else {
					middle_num = nums2[index_2]
					goto result2
				}
			} else {
				if nums1[index_1] < nums2[index_2] {
					index_1++
				} else {
					index_2++
				}
				pass--
			}
		}

		for index_1 < m {
			if pass == 0 {
				middle_num = nums1[index_1]
				goto result2
			} else {
				index_1++
				pass--
			}
		}

		for index_2 < n {
			if pass == 0 {
				middle_num = nums2[index_2]
				goto result2
			} else {
				index_2++
				pass--
			}
		}

	result2:
		return float64(middle_num)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
