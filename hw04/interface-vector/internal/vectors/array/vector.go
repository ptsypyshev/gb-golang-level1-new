package array

import (
	"math"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/interface-vector/internal/vectors"
)

var _ vectors.Vector3 = ArrayVector3{}

type ArrayVector3 [3]float64

func New(x, y, z float64) ArrayVector3 {
	return ArrayVector3{x, y, z}
}

func (a ArrayVector3) X() float64 {
	return a[0]
}

func (a ArrayVector3) Y() float64 {
	return a[1]
}

func (a ArrayVector3) Z() float64 {
	return a[2]
}

// Сложение векторов и возврат нового вектора {a.X + u.X, a.Y + u.Y, a.Z + u.Z}
func (a ArrayVector3) Add(u vectors.Vector3) vectors.Vector3 {
	var res ArrayVector3
	res[0] = a.X() + u.X()
	res[1] = a.Y() + u.Y()
	res[2] = a.Z() + u.Z()
	return res
}

// Вычитаение векторов и возврат нового вектора {a.X - u.X, a.Y - u.Y, a.Z - u.Z}
func (a ArrayVector3) Subtract(u vectors.Vector3) vectors.Vector3 {
	var res ArrayVector3
	res[0] = a.X() - u.X()
	res[1] = a.Y() - u.Y()
	res[2] = a.Z() - u.Z()
	return res
}

// Умножение вектора на число. Умножьте кажду каоордианту на число и верните веткор
func (a ArrayVector3) Multiply(scalar float64) vectors.Vector3 {
	var res ArrayVector3
	res[0] = a.X() * scalar
	res[1] = a.Y() * scalar
	res[2] = a.Z() * scalar
	return res
}

// Скалярное произведение векторов. Перемножьте координаты вектора v на координаты вектора u
// Пример a.x * u.x
func (a ArrayVector3) Dot(u vectors.Vector3) vectors.Vector3 {
	var res ArrayVector3
	res[0] = a.X() * u.X()
	res[1] = a.Y() * u.Y()
	res[2] = a.Z() * u.Z()
	return res
}

// Вычисление длины вектора math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
func (a ArrayVector3) Length() float64 {
	return math.Sqrt(a.X()*a.X() + a.Y()*a.Y() + a.Z()*a.Z())
}
