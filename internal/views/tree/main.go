package tree

import (
	"bytes"
)

type View struct {
	Root      *Node
	selection int
	buf       []string
}

func NewView(root *Node, maxSize int) *View {
	return &View{
		Root: root,
		buf:  make([]string, maxSize),
	}
}

func (v *View) Selection() int {
	return v.selection
}

func (v *View) Render() []string {
	v.buf = make([]string, len(v.buf))
	v.render(0, 0, v.Root)
	return v.buf
}

func (v *View) render(lvl int, row int, node *Node) int {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < lvl; i++ {
		buf.WriteRune('.')
	}
	if node.Expandable {
		if len(node.Children) > 0 {
			buf.WriteRune('v')
		} else {
			buf.WriteRune('>')
		}
	} else {
		buf.WriteRune('o')
	}
	buf.WriteString(" " + node.Dat)
	v.buf[row] = buf.String()
	if node.Selected {
		v.selection = row
	}

	// children
	nRows := 0
	for _, child := range node.Children {
		nRows += v.render(lvl+1, row+1+nRows, child)
	}
	return 1 + nRows
}
