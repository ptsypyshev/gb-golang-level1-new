package main

import (
	"fmt"

	ar "github.com/ptsypyshev/gb-golang-level1-new/hw04/interface-vector/internal/vectors/array"
	f "github.com/ptsypyshev/gb-golang-level1-new/hw04/interface-vector/internal/vectors/functions"
	st "github.com/ptsypyshev/gb-golang-level1-new/hw04/interface-vector/internal/vectors/structure"
)

func main() {
	vec := ar.New(1, 1, 1)
	vec1 := st.New(0, 0, 0)
	f.PrintVec(vec, vec1)

	vec2 := ar.New(5, 2, 3)
	vec3 := st.New(4, 7, 1)
	fmt.Println(f.Sum(vec, vec1, vec2, vec3))

	bytes, err := f.MarshalA(vec2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Marshaled: %s\n", bytes)

		var vec4 ar.ArrayVector3
		err = f.UnmarshalA(bytes, &vec4)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Unmarshaled: %+v\n", vec4)
		}
	}

	bytes, err = f.MarshalS(vec3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Marshaled: %s\n", bytes)

		var vec5 st.StructVector3
		err = f.UnmarshalS(bytes, &vec5)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Unmarshaled: %+v\n", vec5)
		}
	}
}
