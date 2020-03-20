package main

import (
	"fmt"

	"github.com/ffacs/Bling/draw"
	"github.com/ffacs/Bling/physic"

	//"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	//"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func AddBall(num, inity, initx int, initvx, initvy, initax, initay float64) {
	obj := physic.NewObj(initx, inity, initvx, initvy, initax, initay)
	for {
		y, x, t, xmov := obj.NextPos()
		var ny, nx int
		if xmov {
			ny = obj.Y
			nx = int(x)
		} else {
			ny = int(y)
			nx = obj.X
		}
		lock.Lock()
		if draw.MAP[ny][nx] != draw.MAP[obj.Y][obj.X] && draw.MAP[ny][nx] != ' ' {
			obj.Collision(xmov, !xmov)
		} else {
			if ny == obj.Y && nx == obj.X {
				fmt.Println(ny, nx)
			}
			obj.Move(ny, nx, y, x, t)
		}
		draw.Addpoint(obj.Y, obj.X, num)
		lock.Unlock()
		time.Sleep(time.Duration(int(t * 300000000)))
	}
}

var wg sync.WaitGroup

const BallLim = 20

func main() {
	draw.DrawMap()
	wg.Add(1)
	cnt := 0
	for i := 1; i <= draw.MAP_ROW; i++ {
		for j := 1; j <= draw.MAP_COLUMN; j++ {
			if draw.MAP[i][j] != '*' {
				draw.Points = append(draw.Points, draw.NewPoint())
				go AddBall(cnt, i, j, 5, 0, 0, 1)
				cnt++
			}
			if cnt == BallLim {
				break
			}
		}
		if cnt == BallLim {
			break
		}
	}
	wg.Wait()
}
