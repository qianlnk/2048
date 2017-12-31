package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/qianlnk/log"
	"github.com/qianlnk/to"
)

const (
	GSIZE = 4
)

type GameCells [GSIZE][GSIZE]int

func NewGame() *GameCells {
	termbox.Init()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputEsc)
	game := new(GameCells)
	game.Genarate()
	game.Genarate()

	log.SetOutputPath("./game.log")
	log.SetFormatter("logstash")

	log.Info("init", *game)
	game.Draw(false)
	return game
}

func (g *GameCells) Up() bool {
	var ok bool
	for col := 0; col < GSIZE; col++ {
		for row := 0; row < GSIZE-1; {
			add := false
			if g[row][col] != 0 {
				for r := row + 1; r < GSIZE; r++ {
					if g[r][col] == 0 {
						continue
					}
					if g[row][col] == g[r][col] {
						g[row][col] += g[r][col]
						g[r][col] = 0
						ok = true
						row = r + 1
						add = true
					}

					break
				}
			}

			if !add {
				row++
			}
		}
	}

	for col := 0; col < GSIZE; col++ {
		for row := 0; row < GSIZE-1; row++ {
			if g[row][col] == 0 {
				for r := row + 1; r < GSIZE; r++ {
					if g[r][col] == 0 {
						continue
					}
					g[row][col] = g[r][col]
					g[r][col] = 0
					ok = true
					break
				}
			}
		}
	}

	return ok
}

func (g *GameCells) Down() bool {
	var ok bool
	for col := GSIZE - 1; col >= 0; col-- {
		for row := GSIZE - 1; row >= 1; {
			add := false
			if g[row][col] != 0 {
				for r := row - 1; r >= 0; r-- {
					if g[r][col] == 0 {
						continue
					}

					if g[row][col] == g[r][col] {
						g[row][col] += g[r][col]
						g[r][col] = 0
						ok = true
						row = r - 1
						add = true
					}

					break
				}
			}

			if !add {
				row--
			}
		}
	}

	for col := GSIZE - 1; col >= 0; col-- {
		for row := GSIZE - 1; row >= 1; row-- {
			if g[row][col] == 0 {
				for r := row - 1; r >= 0; r-- {
					if g[r][col] == 0 {
						continue
					}
					g[row][col] = g[r][col]
					g[r][col] = 0
					ok = true
					break
				}
			}
		}
	}

	return ok
}

func (g *GameCells) Left() bool {
	var ok bool
	for row := 0; row < GSIZE; row++ {
		for col := 0; col < GSIZE-1; {
			add := false
			if g[row][col] != 0 {
				for c := col + 1; c < GSIZE; c++ {
					if g[row][c] == 0 {
						continue
					}

					if g[row][col] == g[row][c] {
						g[row][col] += g[row][c]
						g[row][c] = 0
						ok = true
						col = c + 1
						add = true
					}

					break
				}
			}

			if !add {
				col++
			}
		}
	}

	for row := 0; row < GSIZE; row++ {
		for col := 0; col < GSIZE; col++ {
			if g[row][col] == 0 {
				for c := col + 1; c < GSIZE; c++ {
					if g[row][c] == 0 {
						continue
					}
					g[row][col] = g[row][c]
					g[row][c] = 0
					ok = true
					break
				}
			}
		}
	}

	return ok
}

func (g *GameCells) Right() bool {
	var ok bool
	for row := GSIZE - 1; row >= 0; row-- {
		for col := GSIZE - 1; col >= 1; {
			add := false
			if g[row][col] != 0 {
				for c := col - 1; c >= 0; c-- {
					if g[row][c] == 0 {
						continue
					}

					if g[row][col] == g[row][c] {
						g[row][col] += g[row][c]
						g[row][c] = 0
						ok = true
						col = c - 1
						add = true
					}

					break
				}
			}

			if !add {
				col--
			}
		}
	}

	for row := GSIZE - 1; row >= 0; row-- {
		for col := GSIZE - 1; col >= 1; col-- {
			if g[row][col] == 0 {
				for c := col - 1; c >= 0; c-- {
					if g[row][c] == 0 {
						continue
					}

					g[row][col] = g[row][c]
					g[row][c] = 0
					ok = true
					break
				}
			}
		}
	}

	return ok
}

func (g *GameCells) GameOver() bool {
	for row := 0; row < GSIZE; row++ {
		for col := 0; col < GSIZE-1; col++ {
			if g[row][col] == 0 || g[row][col+1] == 0 {
				return false
			}
			if g[row][col] == g[row][col+1] {
				return false
			}
		}
	}

	for col := 0; col < GSIZE; col++ {
		for row := 0; row < GSIZE-1; row++ {
			if g[row][col] == g[row+1][col] {
				return false
			}
		}
	}

	return true
}

func (g *GameCells) Genarate() {
	type point struct {
		row int
		col int
	}

	var emptyCells []point
	for row := 0; row < GSIZE; row++ {
		for col := 0; col < GSIZE; col++ {
			if g[row][col] == 0 {
				emptyCells = append(emptyCells, point{row, col})
			}
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(emptyCells))
	val := r.Intn(10)
	if val >= 8 {
		val = 4
	} else {
		val = 2
	}
	g[emptyCells[index].row][emptyCells[index].col] = val
}

func (g *GameCells) Draw(gameOver bool) {
	drawTable(BorderDouble, 0, 0, 32, 80, 4, 4, termbox.ColorRed, termbox.ColorDefault)
	top := 1

	for row := 0; row < GSIZE; row++ {
		left := 1
		for col := 0; col < GSIZE; col++ {
			if g[row][col] != 0 {
				drawCell(left, top, 20, 8, to.String(g[row][col]))
			}
			left += 20 + 1
		}
		top += 8 + 1
	}

	if gameOver {
		drawCell(1, 16, 83, 6, "Game Over")
	}

	termbox.Flush()
}

func (g *GameCells) Play() {
	for {
		redraw := false
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				g.Close()
				return
			case termbox.KeyArrowUp:
				if redraw = g.Up(); redraw {
					log.Info("U", *g)
				}
			case termbox.KeyArrowDown:
				if redraw = g.Down(); redraw {
					log.Info("D", *g)
				}
			case termbox.KeyArrowLeft:
				if redraw = g.Left(); redraw {
					log.Info("L", *g)
				}
			case termbox.KeyArrowRight:
				if redraw = g.Right(); redraw {
					log.Info("R", *g)
				}
			}
		}

		if redraw {
			g.Genarate()
			log.Info("G", *g)
			g.Draw(g.GameOver())
		}
	}
}

func (g *GameCells) Close() {
	termbox.Close()
}
