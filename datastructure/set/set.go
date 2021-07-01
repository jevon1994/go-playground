package set

type Set struct {
	Data   int
	Parent int
}

var max = 1000

func NewSet(data int, parent int) *Set {
	return &Set{
		Data:   data,
		Parent: parent,
	}
}

func Find(s []Set, d int) int {
	i := 0
	//找下标
	for ; i < max && s[i].Data != d; i++ {
	}
	if i >= max {
		return -1
	}
	for s[i].Parent >= 0 {
		i = s[i].Parent
	}
	return i
}

// 小的集合并到大的集合,避免树过高
func Union(s []Set, a int, b int) {
	r1 := Find(s, a)
	r2 := Find(s, b)
	if r1 != r2 {
		s[r2].Parent = r1
	}
}

// 优化后的结构
type SetN int

func FindN(s []SetN, x int) int {
	for ; s[x] >= 0; x = int(s[x]) {
	}
	return x
}

func UnionN(s []SetN, a int, b int) {
	s[a] = SetN(b)
}

func UnionNBySize(s []SetN, r1 int, r2 int) {
	// 负数: 集合 2 规模 比集合 1大
	if s[r2] < s[r1] {
		s[r2] += s[r1] // 小集合并入大集合
		s[r1] = s[r2]
	} else {
		s[r1] += s[r2]
		s[r2] = s[r1]
	}
}
