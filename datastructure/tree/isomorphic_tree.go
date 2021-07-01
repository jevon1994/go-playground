package tree

type Node struct {
	val   int
	left  int
	right int
}

type T struct {
	arry []interface{}
}

/**
in
-	4 2
	3 1 4 2
+    2 4 1 2
+	3 2 4 1
---
-	2 1
+	2 1
+	1 2
	0
out
	yes
	no
	no

解: 1. 需要建树
	2. 不需要建树
*/
// 对应的左子树 和 右子树
func (t *T) Isomorphic(t1 T, t2 T) {

}
