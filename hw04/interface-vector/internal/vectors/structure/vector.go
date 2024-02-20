package structure

import (
	"math"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/interface-vector/internal/vectors"
)

var _ vectors.Vector3 = StructVector3{}

type StructVector3 struct {
	x float64
	y float64
	z float64
}

func New(x, y, z float64) StructVector3 {
	return StructVector3{x: x, y: y, z: z}
}

func (s StructVector3) X() float64 {
	return s.x
}

func (s StructVector3) Y() float64 {
	return s.y
}

func (s StructVector3) Z() float64 {
	return s.z
}

func (s StructVector3) SetX(v  float64) {
	s.x = v
}

func (s StructVector3) SetY(v  float64) {
	s.y = v
}

func (s StructVector3) SetZ(v  float64) {
	s.z = v
}

func (s StructVector3) Add(v vectors.Vector3) vectors.Vector3 {
	var res StructVector3
	res.x = s.x + v.X()
	res.y = s.y + v.Y()
	res.z = s.z + v.Z()
	return res
}

func (s StructVector3) Subtract(v vectors.Vector3) vectors.Vector3 {
	var res StructVector3
	res.x = s.x - v.X()
	res.y = s.y - v.Y()
	res.z = s.z - v.Z()
	return res
}

func (s StructVector3) Multiply(scalar float64) vectors.Vector3 {
	var res StructVector3
	res.x = s.X() * scalar
	res.y = s.Y() * scalar
	res.z = s.Z() * scalar
	return res
}

func (s StructVector3) Dot(v vectors.Vector3) vectors.Vector3 {
	var res StructVector3
	res.x = s.x * v.X()
	res.y = s.y * v.Y()
	res.z = s.z * v.Z()
	return res
}

func (s StructVector3) Length() float64 {
	return math.Sqrt(s.X()*s.X() + s.Y()*s.Y() + s.Z()*s.Z())
}
