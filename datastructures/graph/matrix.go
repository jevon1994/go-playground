package graph

type Graph struct {
	vertex []int   //顶点表
	arc    [][]int //邻接矩阵
	vCount int     //图的顶点数
	eCount int     //边数
}

func NewGraph(vCount int, eCount int, vertex []int, arc [][]int) *Graph {
	g := &Graph{
		vertex: vertex,
		arc:    make([][]int, vCount),
		vCount: vCount,
		eCount: eCount,
	}
	//matrix
	for i, _ := range g.arc {
		tem := make([]int, vCount)
		g.arc[i] = tem
	}
	//for i := 0; i < vCount; i++ {
	//	for j := 0; j < vCount; j++ {
	//		if i == j {
	//			g.arc[i][j] = 0
	//		}else{
	//			g.arc = arc
	//		}
	//	}
	//}
	g.arc = arc
	return g
}

//
//func (g *Graph) InsertVertex(v Vertex){
//
//}
//
//func (g *Graph) InsertEdge(e Edge){
//
//}
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
