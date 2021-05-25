package graph

import "fmt"

type Graph struct {
	vertex []int   //顶点表: 存顶点的数据
	dType  [][]int //邻接矩阵: 类型-顶点/边, [行][列], 行: 出度,列: 入度
	vCount int     //图的顶点数
	eCount int     //边数
}
type Vertex int
type Edge struct {
	V1, V2 Vertex //有向边
	Weight int
}

func Create(vCount int) *Graph {
	g := &Graph{
		vertex: make([]int, vCount),
		dType:  make([][]int, vCount),
		vCount: vCount,
		eCount: 0,
	}
	for v := 0; v < vCount; v++ {
		ints := make([]int, vCount)
		for w := 0; w < vCount; w++ {
			ints[w] = w
		}
		g.dType[v] = ints
	}
	return g
}

func (g *Graph) TraversalGraph() {
	count := g.vCount
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			fmt.Println(g.dType[i][j])
		}
	}
}

func BuildGraph(vCount int, edge []Edge) *Graph {
	g := Create(vCount)
	for i := 0; i < len(edge); i++ {
		g.InsertEdge(edge[i])
	}
	return g
}

//func (g *Graph) InsertVertex(v Vertex){
//
//}
//
func (g *Graph) InsertEdge(e Edge) {
	//边<V1,V2>
	g.dType[e.V1][e.V2] = e.Weight
	// 无向图 需要插入 <V2,V1>
	g.dType[e.V2][e.V1] = e.Weight
}

//层序遍历, 队列
func (g *Graph) BFS(v Vertex) {
	visited[v] = true
	q.push(v)
	while(!q.isEmpty())
	{
		v = q.pop()
		for (v
		的每个邻接点
		w){
		if !visited[w] {
			visited[w] = true
			q.push(w)
		}
	}
	}
}

//树的先序遍历
func (g *Graph) DFS(v Vertex) {
	visited[v] = true
	// 如果都访问过了, 相当于原路返回, 即是堆栈
	for (v
	的每个邻接点
	w)
	if !visited[w] {
		// 递归 == 系统实现堆栈
		g.DFS(w)
	}
}

//
//func (g *Graph) MST(){
//
//}

func NewGraph(vCount int, eCount int, vertex []int, dType [][]int) *Graph {
	g := &Graph{
		vertex: vertex,
		dType:  make([][]int, vCount),
		vCount: vCount,
		eCount: eCount,
	}
	//matrix
	for i, _ := range g.dType {
		tem := make([]int, vCount)
		g.dType[i] = tem
	}
	for i := 0; i < vCount; i++ {
		for j := 0; j < vCount; j++ {
			if i == j {
				g.dType[i][j] = 0
			} else {
				g.dType = dType
			}
		}
	}
	g.dType = dType
	return g
}
