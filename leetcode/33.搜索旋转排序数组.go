/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	length := len(nums)
	low :=  0
	high := length -1  // 定义左右的下标
	
	for low <= high   {
		mid := low + (high -low) / 2		 // 查询中间值
		
		if nums[mid] == target {
			return mid
		}
		
		// 先判断一边是否有序的
		if (nums[low] <= target && target <= nums[mid]) || 
			(nums[mid] <= nums[high] && (target < nums[mid] || target > nums[high])) {
				high = mid -1 
		} else {
				low  = mid  + 1
		}
	}
	return -1
}
// @lc code=end

