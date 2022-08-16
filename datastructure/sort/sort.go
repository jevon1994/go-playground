package sort

func Bubble(arr []int) {
	flag := 0
	size := len(arr) - 1
	for i := size; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
}

func Insertion(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		tmp := arr[i]
		j := 0
		for j = i; j >= 1 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

func Shell(arr []int) {
	size := len(arr)
	//分组
	for D := size / 2; D > 0; D /= 2 {
		//插入排序
		for i := D; i < size; i++ {
			tmp := arr[i]
			j := 0
			// 从后往前 j+D
			for j = i - D; j >= 0 && arr[j] > tmp; j -= D {
				arr[j+D] = arr[j]
			}
			arr[j+D] = tmp
		}
	}
}

func swap(arr []int, a, b int) {
	tmp := 0
	tmp = arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}

func merge(arr, tmp []int, left, right, rightEnd int) {
	leftEnd := right - 1
	t := left
	total := rightEnd - left + 1
	for left <= leftEnd && right <= rightEnd {
		if arr[left] <= arr[right] {
			tmp[t] = arr[left]
			left++
		} else {
			tmp[t] = arr[right]
			right++
		}
		t++
	}
	for left <= leftEnd {
		tmp[t] = arr[left]
		t++
		left++
	}
	for right <= rightEnd {
		tmp[t] = arr[right]
		t++
		right++
	}
	//tmp 拷贝到 arr
	for i := 0; i < total; i++ {
		arr[rightEnd] = tmp[rightEnd]
		rightEnd--
	}
}

func Msort(arr, tmp []int, left, rightEnd int) {
	center := 0
	if left < rightEnd {
		//1. split
		center = (left + rightEnd) / 2
		Msort(arr, tmp, left, center)
		Msort(arr, tmp, center+1, rightEnd)
		//2.merge
		merge(arr, tmp, left, center+1, rightEnd)
	}
}

func Quick(arr []int, left, right int) {
	if left >= right {
		return
	}
	//1. 找主元分割
	var (
		i     = left
		j     = right
		pivot = arr[i]
	)
	for i < j {
		// 由右向左 小于等于pivot的 向左移
		for i < j && pivot <= arr[j] {
			j--
		}
		// 小于 pivot 换到左边
		if i < j {
			swap(arr, i, j)
			i++
		}
		// 由右向左 大于等于pivot的 向左移
		for i < j && pivot >= arr[i] {
			i++
		}
		// 大于 pivot 换到右边
		if i < j {
			swap(arr, j, i)
			j--
		}
	}
	arr[i] = pivot
	//2. 递归两边
	Quick(arr, left, i-1)
	Quick(arr, i+1, right)
}
