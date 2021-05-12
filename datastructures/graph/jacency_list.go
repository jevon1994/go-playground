package graph

type Adj struct {
	Index  int // 邻接点下标
	Weight int // 权重
	Next   *Adj
}

type VertexNode struct {
	Data      int  // 顶点数据
	FirstEdge *Adj // 边指针
}

type AdjGraph struct {
	vCount int
	eCount int
	List   []VertexNode
}
