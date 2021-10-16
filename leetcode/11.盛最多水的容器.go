/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// 分析 1. 两重循环，找冗余，2. 发现关键， 3. 移动最短的那一块的
	// 解法:  双指针
	// 时间空间复杂度
	i := 0
	j := len(height) -1
	var res int 
	for i < j {
		area := (j - i) * min(height[i], height[j])// 计算面积
		if res < area {
			res =  area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

