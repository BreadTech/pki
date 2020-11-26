package controller

import (
	"github.com/BreadTech/pki/internal/asn"
	"github.com/BreadTech/pki/internal/termbox"
	"github.com/BreadTech/pki/internal/views/tree"
	"github.com/BreadTech/pki/pkg"
)

type Service struct {
	root      *asn.Node
	view      *tree.View
	selection *pkg.Stack
}

func New(der []byte) (*Service, error) {
	root, err := asn.New(der)
	if err != nil {
		return nil, err
	}
	return &Service{
		root:      root,
		view:      tree.NewView(tree.NewNode(true, root.String()), root.Size()),
		selection: pkg.NewStack(),
	}, nil
}

func (s *Service) GetViewNode() *tree.Node {
	vn := s.view.Root
	for _, i := range s.selection.Iter() {
		vn = vn.Children[i]
	}
	return vn
}

func (s *Service) GetDataNode() *asn.Node {
	dn := s.root
	for _, i := range s.selection.Iter() {
		dn = dn.Children[i]
	}
	return dn
}

func (s *Service) Run() {
	for sig := ""; sig != "quit"; sig = termbox.Input() {
		switch sig {
		case "open":
			if vn := s.GetViewNode(); !vn.Expandable || len(vn.Children) > 0 {
				break
			} else {
				// expand
				for _, dchild := range s.GetDataNode().Children {
					vn.Add(dchild.IsCompound, dchild.String())
				}
			}

		case "close":
			if vn := s.GetViewNode(); len(vn.Children) > 0 {
				// collapse
				vn.Children = nil
				break
			}

		case "up":
			if s.selection.IsEmpty() {
				break
			}
			vn := s.GetViewNode()
			vn.Selected = false

			if last := s.selection.Peek(); last == 0 {
				// selected node is first child, move to parent
				s.selection.Pop()
			} else {
				// selected node is not first child, move to sibling
				s.selection.Pop()
				s.selection.Push(last - 1)
				// move to bottom most sibling
				for len(s.GetViewNode().Children) > 0 {
					s.selection.Push(len(s.GetViewNode().Children) - 1)
				}
				if vn = s.GetViewNode(); len(vn.Children) > 0 {
					// if sibling node has children, move to last child
					s.selection.Push(len(vn.Children) - 1)
				}
			}
			s.GetViewNode().Selected = true

		case "down":
			vn := s.GetViewNode()
			vn.Selected = false

			if s.selection.IsEmpty() {
				if len(vn.Children) == 0 {
					break
				}

				// empty stack, move to first element
				s.selection.Push(0)
			} else if len(vn.Children) > 0 {
				// selected node has children
				s.selection.Push(0)
			} else if last := s.selection.Peek(); last == len(vn.Parent.Children)-1 {
				// selected node is last child, inc parent
				s.selection.Pop()
				if s.selection.IsEmpty() {
					break
				}
				// pop until valid parent to inc
				last := s.selection.Pop()
				for ; last+1 == len(s.GetViewNode().Children); last = s.selection.Pop() {
				}
				s.selection.Push(last + 1)
			} else {
				// selected node is not last child, move to sibling
				s.selection.Pop()
				s.selection.Push(last + 1)
			}
			s.GetViewNode().Selected = true
		}
		termbox.Render(s.view.Render(), s.view.Selection())
	}
	termbox.Shutdown()
}
