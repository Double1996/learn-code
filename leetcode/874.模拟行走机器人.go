/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
		// 存储障碍物的, 为一个Map
		m := make(map[[2]int]bool, len(obstacles))
		for _, obstacle := range obstacles { // ERROR map 设计不合理
			m[[2]int{obstacle[0], obstacle[1]}] = true
		}
	
		var x, y, res int
		var dir int // N = 0, E = 1, S =2, W =3
		// (dir + 1) %4  // 右转
		// (dir + 3) %4 // 左转
	
		dx := []int{0, 1, 0, -1}  // 网格行走的题目, 技巧方向数组
		dy := []int{1, 0, -1, -0} // 获得方向走向的值
	
		for _, command := range commands {
			if command == -1 {
				dir = (dir + 1) % 4
			} else if command == -2 {
				dir = (dir + 3) % 4
			} else {
				for i := 1; i <= command; i++ {
					nx := x + dx[dir]
					ny := y + dy[dir]
					if _, ok := m[[2]int{nx, ny}]; ok { // 判断是否遇到障碍物, 遇到障碍物停止下来
						break
					}
	
					x = nx
					y = ny
	
					if x*x+y*y > res {
						res = x*x + y*y
					}
				}
			}
		}
		return res
}

// @lc code=end

