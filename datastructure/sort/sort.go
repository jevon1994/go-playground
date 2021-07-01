package sort

func bubble(arr []int) {
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

func insertion(arr []int) {
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

func shell(arr []int) {
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

//1. 找主元分割, left < right, 一
//2. 递归两边
func partion(arr []int, left, right int) int {
	var (
		i     = left
		j     = right
		pivot = arr[left]
	)

	for i != j {
		for i < j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i += 1
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j -= 1
		}
	}
	arr[i] = pivot
	return pivot
}

func Quick(arr []int, left, right int) {
	i := partion(arr, left, right)
	Quick(arr, left, i-1)
	Quick(arr, i+1, right)
}

//1. 拆分
//2. 合并
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
		center = (left + rightEnd) / 2
		Msort(arr, tmp, left, center)
		Msort(arr, tmp, center+1, rightEnd)
		merge(arr, tmp, left, center+1, rightEnd)
	}
}
