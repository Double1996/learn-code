/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	// 思路: 二分查找， 使用迭代的实现
	low, high := 0, len(nums) -1
	for low <= high {							// 定义退出的条件
		mid := low + (high - low) /2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid -1            // 剔除 边界
		} else if nums[mid] < target {
			low = mid + 1 						//
		}
	}
	return -1
	// 时间复杂度：  O(logn)，其中 n 是数组的长度
	// 空间复杂度:   O(1)
}
// @lc code=end

