package termbox

import (
	. "github.com/nsf/termbox-go"
)

var (
	fg, bg        = ColorWhite, ColorBlack
	width, height int
)

func init() {
	Init()
	width, height = Size()
}

func Render(buf []string, selection int) {
	Clear(fg, bg)
	for y, line := range buf {
		fgc := fg
		if y == selection {
			fgc |= AttrUnderline
		}
		for x, ch := range line {
			if x >= width {
				break
			}
			SetCell(x, y, ch, fgc, bg)
		}
	}
	Flush()
}

func Input() string {
	ev := PollEvent()
	switch ev.Ch {
	case 'q':
		return "quit"
	case 'h':
		return "close"
	case 'j':
		return "down"
	case 'k':
		return "up"
	case 'l':
		return "open"
	}
	return string(ev.Ch)
}

func Shutdown() {
	Close()
}
