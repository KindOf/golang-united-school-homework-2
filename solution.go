package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type CustomInt int

const (
	SidesCircle = iota
	_
	_
	SidesTriangle
	SidesSquare
)

func CalcSquare(sideLen float64, sidesNum CustomInt) float64 {
	sidePow := math.Pow(sideLen, 2)

	switch sidesNum {
	case SidesCircle:
		return sidePow * math.Pi
	case SidesTriangle:
		return (math.Sqrt(3) / 4) * sidePow
	case SidesSquare:
		return sidePow
	default:
		return 0
	}
}
