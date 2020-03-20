package physic

import (
	"math"
)

type vector struct {
	x, y float64
}
type object struct {
	X, Y         int
	position     *vector
	acceleration *vector
	volocity     *vector
}

func (obj *object) Getx() float64 {
	return obj.position.x
}

func (obj *object) Gety() float64 {
	return obj.position.y
}

func (obj *object) GetVx() float64 {
	return obj.volocity.x
}

func (obj *object) GetVy() float64 {
	return obj.volocity.y
}
func (obj *object) GetAx() float64 {
	return obj.acceleration.x

}

func (obj *object) GetAy() float64 {
	return obj.acceleration.y
}

func NewObj(initx, inity int, initvx, initvy, initax, initay float64) *object {
	newobj := &object{
		X: initx,
		Y: inity,
		position: &vector{
			x: float64(initx),
			y: float64(inity),
		},
		acceleration: &vector{
			x: initax,
			y: initay,
		},
		volocity: &vector{
			x: initvx,
			y: initvy,
		},
	}
	return newobj
}

func SolveEque(a, v, d float64) (x1, x2 float64) {
	x1 = (-v - math.Sqrt(v*v+2*a*d)) / a
	x2 = (-v + math.Sqrt(v*v+2*a*d)) / a
	return
}

func TimetoNext(p, v, a float64) (bool, float64) {
	d1, d2 := math.Floor(p+1)-p, math.Ceil(p-1)-p
	var res = 1e9
	var zfx = false
	if a == 0 {
		t1 := d1 / v
		t2 := d2 / v
		if t1 > 0 {
			return true, t1
		} else {
			return false, t2
		}
	}
	if v*v+2*a*d1 >= 0 {
		x1, x2 := SolveEque(a, v, d1)
		if x1 > 0 {
			res = x1
			zfx = true
		} else if x2 > 0 {
			res = x2
			zfx = true
		}
	}
	if v*v+2*a*d2 >= 0 {
		x1, x2 := SolveEque(a, v, d2)
		if x1 > 0 && x1 < res {
			res = x1
			zfx = false
		} else if x2 > 0 && x2 < res {
			res = x2
			zfx = false
		}
	}
	return zfx, res
}

func (obj *object) GetNex() (x, zfx bool, dt float64) {
	z1, t1 := TimetoNext(obj.position.y, obj.volocity.y, obj.acceleration.y)
	z2, t2 := TimetoNext(obj.position.x, obj.volocity.x, obj.acceleration.x)
	if t1 < t2 {
		return false, z1, t1
	} else {
		return true, z2, t2
	}
}

func (obj *object) Collision(xcol, ycol bool) {
	if xcol {
		obj.volocity.x *= -1
	}
	if ycol {
		obj.volocity.y *= -1
	}
}

func (obj *object) NextPos() (ny, nx, t float64, xmov bool) {
	x, zfx, t := obj.GetNex()
	//fmt.Print(x,zfx,t)
	xmov = x
	if x {
		if zfx {
			nx = math.Floor(obj.position.x+1) + 0.0001
		} else {
			nx = math.Ceil(obj.position.x-1) - 0.0001
		}
		ny = obj.position.y + obj.volocity.y*t + 0.5*obj.acceleration.y*t*t
	} else {
		if zfx {
			ny = math.Floor(obj.position.y+1) + 0.0001
		} else {
			ny = math.Ceil(obj.position.y-1) - 0.0001
		}
		nx = obj.position.x + obj.volocity.x*t + 0.5*obj.acceleration.x*t*t
	}
	return
}

func (obj *object) Move(nY, nX int, ny, nx, t float64) {
	obj.position.x = nx
	obj.position.y = ny
	obj.X = nX
	obj.Y = nY
	obj.volocity.y = obj.volocity.y + t*obj.acceleration.y
	obj.volocity.x = obj.volocity.x + t*obj.acceleration.x
}
