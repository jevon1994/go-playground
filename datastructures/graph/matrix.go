package graph

import "fmt"

type Graph struct {
	vertex []int   //顶点表: 存顶点的数据
	dType  [][]int //邻接矩阵: 类型-顶点/边
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
	// 无向图
	g.dType[e.V2][e.V1] = e.Weight
}

//
//func (g *Graph) BFS(v Vertex){
//
//}
////树的先序遍历
//func (g *Graph) DFS(v Vertex){
//
//}
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
