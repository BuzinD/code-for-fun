package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/mattn/go-runewidth"
	"os"
)

const (
	poolWidth            = 40
	poolHeight           = 40
	border               = '|'
	snakeB               = '█'
	up         direction = "up"
	down       direction = "down"
	left       direction = "left"
	right      direction = "right"
)

var defStyle = tcell.StyleDefault.Foreground(tcell.ColorCadetBlue).Background(tcell.ColorWhite)

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func displayTextOnCenter(s tcell.Screen, str string) {
	w, h := s.Size()
	s.Clear()
	emitStr(s, w/2-runewidth.StringWidth(str)/2, h/2, defStyle, str)
	emitStr(s, w/2-9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")
	s.Show()
}

type game struct {
	state string
}

type field struct {
	xStart, yStart, xEnd, yEnd int
	pool                       [][]bool
	scWidth                    int
	scHeight                   int
}

func (f *field) screenIsSmall() bool {
	return f.xEnd > f.scWidth || f.yEnd > f.scHeight
}

func newField(w, h int) *field {
	xStart := (w - poolWidth) / 2
	yStart := (h - poolHeight) / 2

	return &field{
		xStart:   xStart,
		yStart:   yStart,
		xEnd:     xStart + poolWidth,
		yEnd:     yStart + poolHeight,
		pool:     make([][]bool, poolHeight-2),
		scWidth:  w,
		scHeight: h,
	}
}

func getCenter() (x int, y int) {
	return poolWidth / 2, poolHeight / 2
}

type direction string

type snake struct {
	x, y int
	dir  direction
}

func main() {
	//encoding.Register()
	s := initScreen()

	w, h := s.Size()
	f := newField(w, h)

	x, y := getCenter()
	sn := &snake{x, y, up}

	drawBorder(s, f)
	f.pool[1] = []bool{true, true, true, true, true, true}
	drawPool(s, f)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
			w, h := s.Size()
			f = newField(w, h)
			drawBorder(s, f)
			drawPool(s, f)
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyLeft:
				sn.dir = left
				s.SetContent(1, 1, 'l', nil, tcell.StyleDefault)
			case tcell.KeyRight:
				sn.dir = right
				s.SetContent(1, 1, 'r', nil, tcell.StyleDefault)
			case tcell.KeyUp:
				sn.dir = up
				s.SetContent(1, 1, 'u', nil, tcell.StyleDefault)
			case tcell.KeyDown:
				sn.dir = down
				s.SetContent(1, 1, 'd', nil, tcell.StyleDefault)
			case tcell.Key('p'):
				fallthrough
			case tcell.Key('P'):
				//todo pause
			default:
				//don't do something
			}
			//case snake:

		}
	}
}

func initScreen() tcell.Screen {
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	blackAndWhiteStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(blackAndWhiteStyle)

	return s
}

func drawBorder(s tcell.Screen, f *field) {

	if f.screenIsSmall() {
		displayTextOnCenter(s, "Make your screen greater")
		return
	}

	s.Clear()

	// Horizontal borders
	for x := f.xStart; x < f.xEnd; x++ {
		s.SetContent(x, f.yStart, border, nil, tcell.StyleDefault)
		s.SetContent(x, f.yStart+poolHeight-1, border, nil, tcell.StyleDefault)
	}

	// Vertical borders
	for y := f.yStart; y < f.yStart+poolHeight; y++ {
		s.SetContent(f.xStart, y, border, nil, tcell.StyleDefault)
		s.SetContent(f.xStart+poolWidth-1, y, border, nil, tcell.StyleDefault)
	}
}

func drawPool(s tcell.Screen, f *field) {
	//s.Clear()
	for y := 0; y < len(f.pool); y++ {
		for x := 0; x < len(f.pool[y]); x++ {
			if f.pool[y][x] {
				s.SetContent(x+f.xStart+1, y+f.yStart+1, snakeB, nil, tcell.StyleDefault)
			} else {
				s.SetContent(x+f.xStart+1, y+f.yStart+1, ' ', nil, tcell.StyleDefault)
			}
		}
	}
}
