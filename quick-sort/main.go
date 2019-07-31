package problem

/**
* 划分-交换排序（Partition-exchange sort
* 分区+递归
**/
func main(o []int) []int {
	quick(o, 0, len(o)-1)
	return o
}

func quick(arr []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := arr[l]
	i := l + 1

	for j := l + 1; j <= r; j++ {
		if pivot > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[l], arr[i-1] = arr[i-1], arr[l]

	quick(arr, l, i-2)
	quick(arr, i, r)
}

func quick1(arr []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := arr[l]
	i := l + 1
	j := r

	for {
		for arr[i] < pivot {
			i++
		}

		for arr[j] > pivot {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			break
		}
	}

	arr[l], arr[i-1] = arr[i-1], arr[l]
	quick(arr, l, i-2)
	quick(arr, i, r)
}
