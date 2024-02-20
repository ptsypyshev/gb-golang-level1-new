package functions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/interface-vector/internal/vectors"
	st "github.com/ptsypyshev/gb-golang-level1-new/hw04/interface-vector/internal/vectors/structure"
	ar "github.com/ptsypyshev/gb-golang-level1-new/hw04/interface-vector/internal/vectors/array"
)

func Sum(vectors ...vectors.Vector3) vectors.Vector3 {
	var res st.StructVector3
	for _, v := range vectors {
		res = res.Add(v).(st.StructVector3)
	}
	return res
}

func PrintVec(v, v1 vectors.Vector3) {
	fmt.Println(v.Add(v1))
	fmt.Println(v.Subtract(v1))
	fmt.Println(v.Multiply(5))
	fmt.Println(v.Dot(v1))
	fmt.Println(v.Length())
	fmt.Println(v1.Length())
}

func MarshalS(s st.StructVector3) ([]byte, error) {
	var sb strings.Builder
	sb.WriteString(strconv.FormatFloat(s.X(), 'f', -1, 64))
	sb.WriteByte(' ')
	sb.WriteString(strconv.FormatFloat(s.Y(), 'f', -1, 64))
	sb.WriteByte(' ')
	sb.WriteString(strconv.FormatFloat(s.Z(), 'f', -1, 64))
	return []byte(sb.String()), nil
}

func UnmarshalS(data []byte, s *st.StructVector3) error {
	splitted := strings.Split(string(data), " ")
	if len(splitted) != 3 {
		return errors.New("cannot unmarshal to structVector3: bad input")
	}

	for i, elem := range splitted {
		f, err := strconv.ParseFloat(elem, 64)
		if err != nil {
			return errors.New("cannot unmarshal to structVector3: bad float")
		}
		switch i {
		case 0:
			s.SetX(f)
		case 1:
			s.SetY(f)
		case 2:
			s.SetZ(f)
		}
	}
	return nil
}

func MarshalA(a ar.ArrayVector3) ([]byte, error) {
	var sb strings.Builder
	for _, v := range a {
		if sb.Len() == 0 {
			sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		} else {
			sb.WriteByte(' ')
			sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		}
	}
	return []byte(sb.String()), nil
}

func UnmarshalA(data []byte, a *ar.ArrayVector3) error {
	splitted := strings.Split(string(data), " ")
	if len(splitted) != 3 {
		return errors.New("cannot unmarshal to arrayVector3: bad input")
	}
	for i, elem := range splitted {
		f, err := strconv.ParseFloat(elem, 64)
		if err != nil {
			return errors.New("cannot unmarshal to arrayVector3: bad float")
		}
		a[i] = f
	}
	return nil
}
