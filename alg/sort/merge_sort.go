package sort

/**
归并排序的实现: 先排序两个子数组，在合并 两个有序的数组. 和 快速排序的思想恰恰相反
写法:
*/

func merge(a []int, b []int) []int {
	var r = make([]int, len(a)+len(b)) // 定义合并后的数组
	var i, j = 0, 0                    // 定义两个 i,j index

	for i < len(a) && j < len(b) { // 如果 i 小于a 的长度， j  小于 b 的长度
		if a[i] < b[j] { // 判断 i, j 当中最小的数字, 然后 进行归并操作
			r[i+j] = a[i]
			i++ // 迭代++
		} else {
			r[i+j] = b[j] //  最小的进行++
			j++
		}
	}

	for i < len(a) { // a -> i , b -> j
		r[i+j] = a[i] //
		i++
	}

	for j < len(b) {
		r[i+j] = a[j]
		j++
	}

	return r
}

func MergeSort(items []int) []int {
	if len(items) < 2 { // 找到小于2
		return items
	}
	var middle = len(items) >> 1      // 中间的mid
	var a = MergeSort(items[:middle]) // 合并左边界
	var b = MergeSort(items[middle:]) // 合并右边界
	return merge(a, b)
}
