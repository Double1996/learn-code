/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	
  low := 0
  high := len(nums)-1

	for low <= high {
		
		mid := low + (high -low) / 2
		if nums[mid] == target { // 只有找到 中位数的时候 才行
			i := mid
			j := mid 
			for (i > 0 && nums[i] == target) || (j <len(nums)-1 && nums[j] == target) { // 左右的指针扩散
				if i > 0 && nums[i] == target {
						i--
				}
				
				if j < len(nums)-1 && nums[j] == target {
						j++	
				}

			}
			return []int{i + 1, j -1 }
		} else  {
			if nums[mid] > target {  //  没有找到这个数， 接着二分法进行遍历
				high = mid - 1
			} else {
			  low  = mid + 1 		
			}
		}

	}
	return []int{-1, -1}
}
// @lc code=end

