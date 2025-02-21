package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"math/rand"
	"os"
	"time"
)

const (
	poolWidth            = 40
	poolHeight           = 40
	border               = '|'
	snakeB               = '█'
	appleB               = '*'
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

type apple struct {
	x, y int
}

func generateApple(f *field) *apple {
	for {
		x := rand.Intn(poolWidth - 2)
		y := rand.Intn(poolHeight - 2)
		if !f.pool[y][x] {
			return &apple{x, y}
		}
	}
}

type gameState struct {
	state        string
	speed        int64 // snake moving speed in milliseconds
	snake        *snake
	apple        *apple
	score        int32
	cancelTicker context.CancelFunc
	tickerCtx    context.Context
	events       chan string
}

func (s *gameState) IncreaseScore() {
	game.score++
	s.increaseSnakeSpeed()
}

func (s *gameState) increaseSnakeSpeed() {
	if s.score > 0 && s.score%2 == 0 {
		if s.cancelTicker != nil {
			s.cancelTicker()
		} else {
			fmt.Println("s.cancelTicker == nil")
		}
		game.speed -= 10
		s.initTicker()
	}
}

func (s *gameState) initTicker() {
	s.tickerCtx, s.cancelTicker = context.WithCancel(context.Background())

	ticker := time.NewTicker(time.Millisecond * time.Duration(game.speed))
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-s.tickerCtx.Done():
				return
			case <-ticker.C:
				s.events <- "move"
			}
		}
	}()
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

func (f *field) move(sn *snake) error {
	snakeHeadY := sn.body[0][0]
	snakeHeadX := sn.body[0][1]
	snakeTailY := sn.body[len(sn.body)-1][0]
	snakeTailX := sn.body[len(sn.body)-1][1]
	var newX, newY int
	switch sn.dir {
	case up:
		newX = snakeHeadX
		newY = snakeHeadY - 1
	case down:
		newX = snakeHeadX
		newY = snakeHeadY + 1
	case left:
		newX = snakeHeadX - 1
		newY = snakeHeadY
	case right:
		newX = snakeHeadX + 1
		newY = snakeHeadY
	}

	if newX < 0 || newY < 0 {
		return errors.New("snake out of range")
	}

	if newX >= poolWidth-2 || newY >= poolHeight-2 {
		return errors.New("snake out of range")
	}

	if f.pool[newY][newX] {
		return errors.New("snake is already in a snake")
	}

	if game.apple.x == newX && game.apple.y == newY {
		sn.putAppleIntoBody(newX, newY)
		game.apple = generateApple(f)
		game.IncreaseScore()
	} else {
		sn.moveBody(newX, newY)

		f.pool[snakeTailY][snakeTailX] = false
	}

	//move into a pool
	f.pool[newY][newX] = true

	return nil
}

func newField(w, h int) *field {
	xStart := (w - poolWidth) / 2
	yStart := (h - poolHeight) / 2

	pool := make([][]bool, poolHeight-2) //pool heights - 2 (it's a borders)

	for i := range pool {
		pool[i] = make([]bool, poolWidth-2) //pool widths - 2 (it's a borders)
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

// putAppleIntoBody add elem into snake body
func (s *snake) putAppleIntoBody(x int, y int) {
	s.body = append(s.body, []int{-100, -100}) //make snake body more by one el
	copy(s.body[1:], s.body)                   //shift all el to the tail
	s.body[0] = []int{y, x}
}

func (s *snake) moveBody(x int, y int) {
	copy(s.body[1:len(s.body)], s.body)
	s.body[0] = []int{y, x}
}

func newSnake(x, y int, dir direction) *snake {
	body := make([][]int, 1)
	body[0] = []int{y, x}
	return &snake{body, dir}
}

func newGame(snX, snY int, f *field) {
	game = gameState{state: "playing", speed: 300, snake: newSnake(snX, snY, up), apple: generateApple(f), score: 0}
	game.events = make(chan string)
	game.initTicker()
}

func main() {
	s := initScreen()
	x, y := getCenter()
	w, h := s.Size()
	f := newField(w, h)
	f.add(x, y)

	newGame(x, y, f)
	defer game.cancelTicker()
	defer close(game.events)

	drawBorder(s, f)
	drawPool(s, f)
	ctx, cancel := context.WithCancel(context.Background())
	readUserActions(ctx, s, game.events)

	drawState(s)
	s.Show()

	for {
		select {
		case event := <-game.events:
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
					if err := f.move(game.snake); err != nil {
						game.state = "game_over"
						displayTextOnCenter(s, "Game Over")
					} else {
						drawPool(s, f)
						drawState(s)
						s.Show()
					}
				}
			}
		}
	}
}

func drawState(s tcell.Screen) {
	emitStr(s, 0, 0, tcell.StyleDefault, fmt.Sprintf("Score: %d", game.score))
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
		game.state = "screen_is_small"
		return
	}
	game.state = "playing"

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
	if game.state != "playing" {
		return
	}
	for y := 0; y < len(f.pool); y++ {
		for x := 0; x < len(f.pool[y]); x++ {
			if f.pool[y][x] {
				s.SetContent(x+f.xStart+1, y+f.yStart+1, snakeB, nil, tcell.StyleDefault)
			} else {
				s.SetContent(x+f.xStart+1, y+f.yStart+1, ' ', nil, tcell.StyleDefault)
			}
		}
	}

	s.SetContent(game.apple.x+f.xStart+1, game.apple.y+f.yStart+1, appleB, nil, tcell.StyleDefault)
}
