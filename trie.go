package lite

type Tree struct {
	root *node
}

type node struct {
	segment string
	handler Handler
	child   []*node
}

func NewTree() *Tree {
	return &Tree{
		root: &node{},
	}
}

func (*Tree) AddRouter(method, uri string, h Handler)
