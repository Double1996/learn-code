package sort

/**
归并排序的实现: 先排序两个子数组，在合并 两个有序的数组. 和 快速排序的思想恰恰相反
写法:
*/

func merge(a []int, b []int) []int {
	var r = make([]int, len(a)+len(b)) // 定义合并后的数组
	var i, j = 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = a[j]
		j++
	}

	return r
}

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	var middle = len(items) / 2
	var a = MergeSort(items[:middle])
	var b = MergeSort(items[middle:])
	return merge(a, b)
}
