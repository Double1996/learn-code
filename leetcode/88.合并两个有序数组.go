/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
		// 双指针的思路
		// 时间复杂度  Olog(n)
		// 空间复杂度  2n

		// var a,b int
		// sorted := make([]int, 0, m+n)
		
		// for  {
		// 	if a == m {
		// 		sorted = append(sorted, nums2[b:]...)
		// 		break
		// 	}

		// 	if b == n {
		// 		sorted = append(sorted, nums1[a:]...)
		// 		break
		// 	}
			
		// 	if nums1[a] < nums2[b] {
		// 		sorted= append(sorted, nums1[a])
		// 		a++
		// 	} else {
		// 		sorted= append(sorted, nums2[b])
		// 		b++			
		// 	}
		// }
		// copy(nums1, sorted)

		index := len(nums1) - 1
		m = m -1
		n = n -1
		for m >= 0 && n >= 0 {	// 迭代遍历
			num1 := nums1[m]
			num2 := nums2[n]
			if num1 > num2 {
				nums1[index] = num1
				m--
			} else {
				nums1[index] = num2
				n--
		}
			index--
		}
		if n >= 0 {  // m +n == nums1.Length // 最大
				copy(nums1[0:n+1], nums2[0:n+1])  // copy 作用，把 nums2 后面的值 n+1 进行最后的替换
		}

}
// @lc code=end

