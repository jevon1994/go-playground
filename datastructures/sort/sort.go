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
