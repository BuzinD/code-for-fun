package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"os"
	"time"
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
var game gameState

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

type gameState struct {
	state string
	speed int64 // snake moving speed in milliseconds
	snake *snake
}

type field struct {
	xStart, yStart, xEnd, yEnd int
	pool                       [][]bool
	scWidth                    int
	scHeight                   int
}

func (f *field) add(x, y int) {
	f.pool[y][x] = true
}

func (f *field) screenIsSmall() bool {
	return f.xEnd > f.scWidth || f.yEnd > f.scHeight
}

func newField(w, h int) *field {
	xStart := (w - poolWidth) / 2
	yStart := (h - poolHeight) / 2

	pool := make([][]bool, poolHeight-2)

	for i := range pool {
		pool[i] = make([]bool, w-2)
	}

	return &field{
		xStart:   xStart,
		yStart:   yStart,
		xEnd:     xStart + poolWidth,
		yEnd:     yStart + poolHeight,
		pool:     pool,
		scWidth:  w,
		scHeight: h,
	}
}

// get center of field
func getCenter() (x int, y int) {
	return poolWidth / 2, poolHeight / 2
}

type direction string

type snake struct {
	body [][]int //presented by slice of []int{y, x} where y, x are coords in a pool
	dir  direction
}

func newSnake(x, y int, dir direction) *snake {
	body := make([][]int, 1)
	body[0] = []int{y, x}
	return &snake{body, dir}
}

func main() {
	//encoding.Register()
	s := initScreen()
	x, y := getCenter()
	game = gameState{"playing", 300, newSnake(x, y, up)}

	w, h := s.Size()
	f := newField(w, h)
	f.add(x, y)

	drawBorder(s, f)
	f.pool[1] = []bool{true, true, true, true, true, true}
	drawPool(s, f)
	ctx, cancel := context.WithCancel(context.Background())
	events := make(chan string)
	defer close(events)
	readUserActions(ctx, s, events)
	i := 0
	for {
		select {
		case event := <-events:
			emitStr(s, 1, 1, tcell.StyleDefault, event)
			s.Show()
			switch event {
			case "resize":
				s.Sync()
				w, h := s.Size()
				f = newField(w, h)
				drawBorder(s, f)
				drawPool(s, f)
				s.Show()
			case "quit":
				s.Fini()
				cancel()
				os.Exit(0)
			case "left":
				game.snake.dir = left
			case "right":
				game.snake.dir = right
			case "up":
				game.snake.dir = up
			case "down":
				game.snake.dir = down
			case "pause":
				game.state = "paused"
			case "move":
				if game.state == "playing" {
					i++
					emitStr(s, 0, 0, tcell.StyleDefault, fmt.Sprintf("Moving %d", i))
					if err := f.move(game.snake); err != nil {
						//todo game over
					} else {
						drawPool(s, f)
						s.Show()
					}
				}
			}
		}
	}
}

func readUserActions(ctx context.Context, s tcell.Screen, events chan string) {

	go func(ctx context.Context) {
		for {
			ev := s.PollEvent() // Блокирующий вызов, ждет событие
			switch ev := ev.(type) {
			case *tcell.EventResize:
				events <- "resize"
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					events <- "quit"
				case tcell.KeyLeft:
					events <- "left"
				case tcell.KeyRight:
					events <- "right"
				case tcell.KeyUp:
					events <- "up"
				case tcell.KeyDown:
					events <- "down"
				case tcell.Key('p'), tcell.Key('P'):
					events <- "pause"
				}
			}
		}
	}(ctx)

	ticker := time.NewTicker(time.Millisecond * time.Duration(game.speed))
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				events <- "move"
			}
		}
	}()
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
