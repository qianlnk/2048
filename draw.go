package main

import (
	"math"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"github.com/qianlnk/to"
)

type BorderType int

const (
	BorderLighter BorderType = iota
	BorderBolder
	BorderDouble
	BorderCurved
)

type TextPos int

const (
	TextLeft TextPos = iota
	TextMiddle
	TextRight
)

var (
	Borders = map[BorderType][][]rune{
		BorderLighter: [][]rune{
			{'┌', '┬', '┐'},
			{'├', '┼', '┤'},
			{'└', '┴', '┘'},
			{'─', '│', ' '},
		},
		BorderBolder: [][]rune{
			{'┏', '┳', '┓'},
			{'┣', '╋', '┫'},
			{'┗', '┻', '┛'},
			{'━', '┃', ' '},
		},
		BorderDouble: [][]rune{
			{'╔', '╦', '╗'},
			{'╠', '╬', '╣'},
			{'╚', '╩', '╝'},
			{'═', '║', ' '},
		},
		BorderCurved: [][]rune{
			{'╭', '┬', '╮'},
			{'├', '┼', '┤'},
			{'╰', '┴', '╯'},
			{'─', '│', ' '},
		},
	}
)

//宽高不包含线条
func drawTable(bt BorderType, left, top int, height, width int, row, col int, fg, bg termbox.Attribute) {
	avgH, avgW := height/row, width/col
	realH, realW := avgH*row+row+1, avgW*col+col+1
	offsetH := 0
	border := make([][]rune, realH)
	for i := 0; i < realH; i++ {
		tmp := make([]rune, realW)
		border[i] = tmp
	}

	for h := 0; h < realH; h++ {
		offsetW := 0
		for w := 0; w < realW; w++ {
			switch {
			case w == 0 && h == 0:
				border[h][w] = Borders[bt][0][0]
			case w == 0 && h == realH-1:
				border[h][w] = Borders[bt][2][0]
			case w == 0 && h == offsetH:
				border[h][w] = Borders[bt][1][0]
			case w == realW-1 && h == 0:
				border[h][w] = Borders[bt][0][2]
			case w == realW-1 && h == realH-1:
				border[h][w] = Borders[bt][2][2]
			case w == offsetW && h == 0:
				border[h][w] = Borders[bt][0][1]
			case w == offsetW && h == realH-1:
				border[h][w] = Borders[bt][2][1]
			case w == realW-1 && h == offsetH:
				border[h][w] = Borders[bt][1][2]
			case w == offsetW && h == offsetH:
				border[h][w] = Borders[bt][1][1]
			case w == offsetW:
				border[h][w] = Borders[bt][3][1]
			case h == offsetH:
				border[h][w] = Borders[bt][3][0]
			default:
				border[h][w] = Borders[bt][3][2]
			}
			if w == offsetW {
				offsetW += avgW + 1
			}
		}

		if h == offsetH {
			offsetH += avgH + 1
		}
	}

	drawborder(border, left, top, fg, bg)
}

func drawborder(border [][]rune, left int, top int, fg termbox.Attribute, bg termbox.Attribute) {
	termbox.Init()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for h := 0; h < len(border); h++ {
		for w := 0; w < len(border[h]); w++ {
			termbox.SetCell(left+w, top+h, border[h][w], fg, bg)
		}
	}
}

func drawText(text string, width int, tt TextPos, left int, top int, fg termbox.Attribute, bg termbox.Attribute) {
	if width <= runewidth.StringWidth(text) {
		text = runewidth.Truncate(text, width, "")
	} else {
		switch tt {
		case TextLeft:
			text = runewidth.FillRight(text, width)
		case TextMiddle:
			lw := (width-runewidth.StringWidth(text))/2 + runewidth.StringWidth(text)
			text = runewidth.FillLeft(text, lw)
			text = runewidth.FillRight(text, width)
		case TextRight:
			text = runewidth.FillLeft(text, width)
		default:
			text = runewidth.FillRight(text, width)
		}
	}
	offset := 0
	for _, c := range []rune(text) {
		termbox.SetCell(left+offset, top, c, fg, bg)
		if runewidth.IsAmbiguousWidth(c) { //特殊字符填补底色
			termbox.SetCell(left+offset+1, top, ' ', fg, bg)
		}

		offset += runewidth.RuneWidth(c)
	}
}

func drawLine(bt BorderType, width int, left int, top int, fg termbox.Attribute, bg termbox.Attribute) {
	for i := 0; i < width; i++ {
		termbox.SetCell(left+i, top, Borders[bt][3][0], fg, bg)
	}
}

func drawCell(left, top int, width, height int, val string) {
	for l := left; l < left+width; l++ {
		for t := top; t < top+height; t++ {
			termbox.SetCell(l, t, ' ', termbox.ColorWhite, termbox.Attribute(10+int(math.Log2(to.Float64(val)))))
		}
	}

	drawText(to.String(val), width, TextMiddle, left, top+height/2, termbox.ColorWhite, termbox.Attribute(10+int(math.Log2(to.Float64(val)))))
}
