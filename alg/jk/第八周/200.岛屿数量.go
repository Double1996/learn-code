package 第八周

func numIslands(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])
	ans := 0
	//方向数组
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	//初始化并查集
	fa := make([]int, m*n+1)
	//二维索引转一维
	for i := 0; i <= m*n; i++ {
		fa[i] = i
	}

	//循环网格
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//如果碰到水则continue
			if grid[i][j] == '0' {
				continue
			}

			for k := 0; k < 4; k++ {

				nx := i + dx[k]
				ny := j + dy[k]

				//到了边界continue
				if nx >= m || ny >= n || nx < 0 || ny < 0 {
					continue

					//如果相邻是1，则加入并查集 uninonset
				} else if grid[nx][ny] == '1' {
					unionSet(fa, nums(n, i, j), nums(n, nx, ny))
				}
			}
		}
	}

	//最后查找并查集根的个数即可
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			//如果字符为1，且根是自己本身，说明找到一个根，ans++
			if grid[i][j] == '1' && Find(fa, nums(n, i, j)) == nums(n, i, j) {
				ans++
			}
		}

	}

	return ans

}

//二维索引转一维
func nums(n, i, j int) int {
	return i*n + j
}

//并查集函数模板
func Find(fa []int, x int) int {
	if fa[x] == x {
		return x
	}
	fa[x] = Find(fa, fa[x])
	return fa[x]
}

func unionSet(fa []int, x, y int) {
	x, y = Find(fa, x), Find(fa, y)
	if x != y {
		fa[x] = y
	}
}
