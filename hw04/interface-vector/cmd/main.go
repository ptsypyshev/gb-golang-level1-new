package main

import (
	"fmt"
	"math"
)

type Vector3 interface {
	Add(v Vector3) Vector3
	Subtract(u Vector3) Vector3
	Multiply(scalar float64) Vector3
	Dot(u Vector3) Vector3
	Length() float64
	X() float64
	Y() float64
	Z() float64
}

var _ Vector3 = arrayVector3{}

type arrayVector3 [3]float64

func (a arrayVector3) X() float64 {
	return a[0]
}

func (a arrayVector3) Y() float64 {
	return a[1]
}

func (a arrayVector3) Z() float64 {
	return a[2]
}

// Сложение векторов и возврат нового вектора {a.X + u.X, a.Y + u.Y, a.Z + u.Z}
func (a arrayVector3) Add(u Vector3) Vector3 {
	var res arrayVector3
	res[0] = a.X() + u.X()
	res[1] = a.Y() + u.Y()
	res[2] = a.Z() + u.Z()
	return res
}

// Вычитаение векторов и возврат нового вектора {a.X - u.X, a.Y - u.Y, a.Z - u.Z}
func (a arrayVector3) Subtract(u Vector3) Vector3 {
	var res arrayVector3
	res[0] = a.X() - u.X()
	res[1] = a.Y() - u.Y()
	res[2] = a.Z() - u.Z()
	return res
}

// Умножение вектора на число. Умножьте кажду каоордианту на число и верните веткор
func (a arrayVector3) Multiply(scalar float64) Vector3 {
	var res arrayVector3
	res[0] = a.X() * scalar
	res[1] = a.Y() * scalar
	res[2] = a.Z() * scalar
	return res
}

// Скалярное произведение векторов. Перемножьте координаты вектора v на координаты вектора u
// Пример a.x * u.x
func (a arrayVector3) Dot(u Vector3) Vector3 {
	var res arrayVector3
	res[0] = a.X() * u.X()
	res[1] = a.Y() * u.Y()
	res[2] = a.Z() * u.Z()
	return res
}

// Вычисление длины вектора math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
func (a arrayVector3) Length() float64 {
	return math.Sqrt(a.X()*a.X() + a.Y()*a.Y() + a.Z()*a.Z())
}

// Аналогично сделайте для вектора на основе структуры. Формулы аналогичные
var _ Vector3 = structVector3{}

type structVector3 struct {
	x float64
	y float64
	z float64
}

func (s structVector3) X() float64 {
	return s.x
}

func (s structVector3) Y() float64 {
	return s.y
}

func (s structVector3) Z() float64 {
	return s.z
}

func (s structVector3) Add(v Vector3) Vector3 {
	var res structVector3
	res.x = s.x + v.X()
	res.y = s.y + v.Y()
	res.z = s.z + v.Z()
	return res
}

func (s structVector3) Subtract(v Vector3) Vector3 {
	var res structVector3
	res.x = s.x - v.X()
	res.y = s.y - v.Y()
	res.z = s.z - v.Z()
	return res
}

func (s structVector3) Multiply(scalar float64) Vector3 {
	var res structVector3
	res.x = s.X() * scalar
	res.y = s.Y() * scalar
	res.z = s.Z() * scalar
	return res
}

func (s structVector3) Dot(v Vector3) Vector3 {
	var res structVector3
	res.x = s.x * v.X()
	res.y = s.y * v.Y()
	res.z = s.z * v.Z()
	return res
}

func (s structVector3) Length() float64 {
	return math.Sqrt(s.X()*s.X() + s.Y()*s.Y() + s.Z()*s.Z())
}

func Sum(vectors ...Vector3) Vector3 {
	var res structVector3
	for _, v := range vectors {
		res = res.Add(v).(structVector3)
	}
	return res
}

func printVec(v, v1 Vector3) {
	fmt.Println(v.Add(v1))
	fmt.Println(v.Subtract(v1))
	fmt.Println(v.Multiply(5))
	fmt.Println(v.Dot(v1))
	fmt.Println(v.Length())
	fmt.Println(v1.Length())
}

func main() {
	vec := arrayVector3{0: 1, 1: 1, 2: 1}
	vec1 := structVector3{
		x: 0,
		y: 0,
		z: 0,
	}
	printVec(vec, vec1)

	vec2 := arrayVector3{0: 5, 1: 2, 2: 3}
	vec3 := structVector3{
		x: 4,
		y: 7,
		z: 1,
	}
	fmt.Println(Sum(vec, vec1, vec2, vec3))
}
