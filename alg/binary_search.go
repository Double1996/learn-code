package alg

// 递归的写法
func binarySearch1(array []int, target, low, high int) {
	if low >= high {
		return
	}

	mid := low + (high-low)/2 // 寻找 中间

	if array[mid] == target {
		return //
	}

	if target > array[mid] {
		binarySearch1(array, target, mid+1, high)
	} else {
		binarySearch1(array, target, low, mid-1)
	}
}

// 非递归的写法
func binarySearch2(array []int, target, low, high int) {
	for low < high {
		mid := low + (high-low)/2 // 寻找 中间

		if array[mid] == target {
			break
		}
		if target > array[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return
}
