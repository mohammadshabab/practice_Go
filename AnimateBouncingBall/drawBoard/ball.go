package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
)

func main() {

	const (
		cellEmpty = ' '
		cellBall  = '⚾'
		maxFrames = 1200
		speed     = time.Second / 20
	)

	var cell rune
	var (
		px, py int
		vx, vy = 1, 1
	)

	width, height := screen.Size()
	bufLen := (width*2 + 1) * height

	ballWidth := runewidth.RuneWidth(cellBall)

	width = width / ballWidth
	height--

	buf := make([]rune, 0, bufLen)

	//	fmt.Println("capacity of buffer is ", cap(buf))
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	screen.Clear()
	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true

		buf = buf[:0]
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				//fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			//fmt.Println()

			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		fmt.Print(string(buf))
		time.Sleep(speed)
		//	fmt.Println("wid and hie after first loop ", cap(buf))
	}
}
