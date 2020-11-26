package tree

type Node struct {
	Expandable bool
	Children   []*Node
	Dat        string
	Selected   bool
	Parent     *Node
}

func NewNode(isParent bool, dat string) *Node {
	return &Node{
		Expandable: isParent,
		Dat:        dat,
	}
}

func (n *Node) Add(isParent bool, dat string) {
	n.Children = append(n.Children, &Node{
		Expandable: isParent,
		Dat:        dat,
		Parent:     n,
	})
}
