package problem

/**
* 划分-交换排序（Partition-exchange sort
* 分区+递归
**/
func main(o []int) []int {
	quick(o, 0, len(o)-1)
	return o
}

func quick(arr []int, left int, right int) {
	pivot := arr[left]
	i := left + 1
	j := right

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

	if i != left+1 && j != right {
		arr = append(arr[left+1:i+1], arr[i:right+1]...)
		arr[i] = pivot
	}

	quick(arr, left, i)
	quick(arr, i+1, right)
}
