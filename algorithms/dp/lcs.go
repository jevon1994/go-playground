package main

import "fmt"

func lcs(x []byte, y []byte, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	// until a = b , total count
	if x[m-1] == y[n-1] {
		return lcs(x, y, m-1, n-1) + 1
	}

	i := lcs(x, y, m, n-1)
	j := lcs(x, y, m-1, n)
	if i > j {
		return i
	}
	return j
}

func lcsDp(x []byte, y []byte, m int, n int, dp map[string]int) int {
	if m == 0 || n == 0 {
		return 0
	}

	var key = fmt.Sprintf("%s|%s", string(x[:]), string(y[:]))
	exists := dp[key]

	if exists != 0 {
		return dp[key]
	}
	// until a = b , total count
	if x[m-1] == y[n-1] {
		dp[key] = lcsDp(x, y, m-1, n-1, dp) + 1
	} else {
		i := lcsDp(x, y, m, n-1, dp)
		j := lcsDp(x, y, m-1, n, dp)

		if i > j {
			dp[key] = i
		} else {
			dp[key] = j
		}
	}
	return dp[key]
}

func main() {
	x := "ABCBDAB"
	y := "BDCABA"
	//fmt.Println(lcs([]byte(x), []byte(y), len(x), len(y)))
	dp := map[string]int{}
	fmt.Println(lcsDp([]byte(x), []byte(y), len(x), len(y), dp))
}
