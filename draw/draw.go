package draw

import (
	"fmt"
	"strconv"
)

const (
	MAP_ROW        int = 59
	MAP_COLUMN     int = 235
	RGBCHANGESPEDD int = 7
)

var MaxLim = 50

type point struct {
	x, y int
}
type Point struct {
	points           []point
	r, g, b          int
	addr, addg, addb bool
}

func NewPoint() *Point {
	return &Point{
		points: make([]point, 1),
		r:      255,
		g:      128,
		b:      0,
	}
}

var MAP [MAP_ROW + 3][MAP_COLUMN + 3]byte
var Points []*Point

func Gotoxy(y, x int) {
	fmt.Printf("\x1b[%d;%dH", y+1, x+1)
}

func (P *Point) rgbfun() (int, int, int) {
	if P.r-RGBCHANGESPEDD < 0 {
		P.addr = true
	} else if P.r+RGBCHANGESPEDD > 255 {
		P.addr = false
	}
	if P.g-RGBCHANGESPEDD < 0 {
		P.addg = true
	} else if P.g+RGBCHANGESPEDD > 255 {
		P.addg = false
	}
	if P.b-RGBCHANGESPEDD < 0 {
		P.addb = true
	} else if P.b+RGBCHANGESPEDD > 255 {
		P.addb = false
	}
	if P.addr {
		P.r += RGBCHANGESPEDD
	} else {
		P.r -= RGBCHANGESPEDD
	}
	if P.addg {
		P.g += RGBCHANGESPEDD
	} else {
		P.g -= RGBCHANGESPEDD
	}
	if P.addb {
		P.b += RGBCHANGESPEDD
	} else {
		P.b -= RGBCHANGESPEDD
	}
	return P.r, P.g, P.b
}

func init() {
	for i := 0; i < MAP_ROW+2; i++ {
		for j := 0; j < MAP_COLUMN+2; j++ {
			if i == 0 || i == MAP_ROW+1 || j == 0 || j == MAP_COLUMN+1 {
				MAP[i][j] = '*'
			} else {
				MAP[i][j] = ' '
			}
		}
	}
	for i := 1; i <= 40; i++ {
		MAP[30][i] = '*'
		MAP[30][MAP_COLUMN-i+1] = '*'
	}
	for i := 20; i <= 40; i++ {
		MAP[i][41] = '*'
		MAP[i][MAP_COLUMN-40] = '*'
	}
}

func DrawMap() {
	fmt.Print("\x1bc")
	for i := 0; i < MAP_ROW+2; i++ {
		for j := 0; j < MAP_COLUMN+2; j++ {
			fmt.Print(string(MAP[i][j]))
		}
		fmt.Println()
	}
}

func Flush(num int) {
	if len(Points[num].points) >= MaxLim {
		Gotoxy(Points[num].points[0].y, Points[num].points[0].x)
		fmt.Print(" ")
		MAP[Points[num].points[0].y][Points[num].points[0].x] = ' '
		Points[num].points = Points[num].points[1:]
	}
	Gotoxy(Points[num].points[len(Points[num].points)-1].y, Points[num].points[len(Points[num].points)-1].x)
	var r, g, b = Points[num].rgbfun()
	fmt.Printf("\x1b[38;2;%s;%s;%sm*\x1b[38m\x1b[0m", strconv.Itoa(r), strconv.Itoa(g), strconv.Itoa(b))
	Gotoxy(MAP_ROW+2, 1)
}

func Addpoint(y, x, num int) {
	Points[num].points = append(Points[num].points, point{x, y})
	MAP[y][x] = '*' + byte(num+1)
	Flush(num)
}
