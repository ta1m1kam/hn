//package main
//
//import (
//	"math/rand"
//
//	log "github.com/sirupsen/logrus"
//	"github.com/nsf/termbox-go"
//)
//
//const coldef = termbox.ColorDefault
//
//// (3) 描画関数の定義
//// (x, y)座標に四角を表示するだけの関数
//func drawBox(x, y int) {
//	termbox.Clear(coldef, coldef)
//	termbox.SetCell(x, y, '┏', coldef, coldef)
//	termbox.SetCell(x+1, y, '┓', coldef, coldef)
//	termbox.SetCell(x, y+1, '┗', coldef, coldef)
//	termbox.SetCell(x+1, y+1, '┛', coldef, coldef)
//	termbox.Flush() // Flushを呼ばないと描画されない
//}
//
//func main() {
//	if err := termbox.Init(); err != nil {
//		log.Fatal(err)
//	}
//	defer termbox.Close()
//
//	drawBox(0, 0)
//MAINLOOP:
//	for {
//		w, h := termbox.Size()
//		switch ev := termbox.PollEvent(); ev.Type {
//		case termbox.EventKey:
//			switch ev.Key {
//			case termbox.KeyEsc:
//				break MAINLOOP
//			}
//		}
//		drawBox(rand.Intn(w), rand.Intn(h)) // (4) メインループ内で描画関数を呼び出す
//	}
//}
