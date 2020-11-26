package asn

import (
	"encoding/asn1"
	"fmt"
)

type Node struct {
	asn1.RawValue
	Children []*Node
}

func (n *Node) Size() int {
	sum := 0
	for _, child := range n.Children {
		sum += child.Size()
	}
	return 1 + sum
}

func (n *Node) String() string {
	nc := len(n.Children)
	if n.Class == asn1.ClassUniversal {
		switch n.Tag {
		case asn1.TagSequence:
			return fmt.Sprintf("SEQUENCE (%d)", nc)
		case asn1.TagSet:
			return fmt.Sprintf("SET (%d)", nc)
		case asn1.TagOID:
			oid := new(asn1.ObjectIdentifier)
			if _, err := asn1.Unmarshal(n.FullBytes, oid); err != nil {
				panic(err)
			}
			return fmt.Sprintf("OID (%s)", oid.String())
		}
	}
	return fmt.Sprintf("%d:%d::0x%x", n.Class, n.Tag, n.Bytes)
}

func (n *Node) Visit(path []int, do func(*Node) error) error {
	visited := n
	for _, i := range path {
		visited = visited.Children[i]
	}
	return do(visited)
}

func New(dat []byte) (*Node, error) {
	root := new(Node)
	asnRest, err := asn1.Unmarshal(dat, &root.RawValue)
	if err != nil {
		return nil, err
	}
	if len(asnRest) > 0 {
		panic("expected top level asn object to be SET or SEQUENCE")
	}
	return root, parseASN(root, root.Bytes)
}

func parseASN(parent *Node, dat []byte) error {
	node := new(Node)

	parent.Children = append(parent.Children, node)
	asnRest, err := asn1.Unmarshal(dat, &node.RawValue)
	if err != nil {
		return err
	}
	if node.IsCompound {
		if err := parseASN(node, node.Bytes); err != nil {
			return err
		}
	}
	if len(asnRest) > 0 {
		if err := parseASN(parent, asnRest); err != nil {
			return err
		}
	}
	return nil
}
