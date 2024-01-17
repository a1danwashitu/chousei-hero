package mincostflow

type Node struct {
	Type  string
	Id    int
	Prev  *Edge
	Nexts []*Edge
	Dist  any
	Pot   any
	Used  bool
}

type Edge struct {
	Cost any
	Cap  int
	From *Node
	To   *Node
	Rev  *Edge
}

type Graph []*Node
