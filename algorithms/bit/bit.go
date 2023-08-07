package bit

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

func add(a, b int) int {
	var sum, carry int
	for b != 0 {
		sum = a ^ b
		carry = (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func minus(a, b int) int {
	i := add(^b, 1)
	res := add(i, a)
	return res
}

func multiply(a, b int) int {
	a = abs(a)
	b = abs(b)
	res := 0
	for b > 0 {
		if (b & 0x1) > 0 {
			res = add(res, a)
		}
		a = a << 1
		b = b >> 1
	}

	if (a ^ b) < 0 {
		res = add(^res, 1)
	}
	return res
}

func divide(a, b int) int {
	a = abs(a)
	b = abs(b)

	var res, r int
	for i := 31; i >= 0; i-- {
		if (a >> i) >= b {
			res = add(res, 1<<i)
			a = minus(a, b<<i)
		}
	}

	if (a ^ b) < 0 {
		res = add(^res, 1)
	}

	if b > 0 {
		r = add(^a, 1)
	} else {
		r = a
	}
	fmt.Println(r)
	return res
}

func abs(a int) int {
	if a < 0 {
		return add(^a, 1)
	}
	return a
}
