package vectors

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
