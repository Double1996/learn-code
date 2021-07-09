/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// 使用 target 找到 所在的行 
	l, r := 0, len(matrix) -1 
	for l <= r {
		mid := (l + r ) / 2 	
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			l = mid + 1 
		} else {
			r = mid - 1
		}
	}

	if r == -1 {
		r = 0
	}

	left, right := 0, len(matrix[r])-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[r][mid] == target {
			return true
		} else if matrix[r][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
// @lc code=end

