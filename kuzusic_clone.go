package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

const (
	_width = 80
)

type state struct {
	End       bool
	HighScore int
	Hoge      int
}

func drawLoop(sch chan state) {
	for {
		st := <-sch
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		drawLine(1, 0, "EXIT : ESC KEY")
		drawLine(_width-50, 0, fmt.Sprintf("HightScore : %05d", st.HighScore))
		drawLine(0, 1, "--------------------------------------------------------------------------------")
		termbox.Flush()
	}
}

func drawLine(x, y int, str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
	}
}

func controller(stateCh chan state) {
	st := initGame()
	stateCh <- st

MAINLOOP:
	for {
		ev := termbox.PollEvent()
		switch ev.Key {
		case termbox.KeyEsc, termbox.KeyCtrlC:
			break MAINLOOP
		}
	}
}

func initGame() state {
	st := state{End: true}
	return st
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	stateCh := make(chan state)

	go drawLoop(stateCh)

	controller(stateCh)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Close()
}
