package main

import (
	"fmt"
	"math"
)

// Um tipo de interface é definida por um conjunto de métodos.
// Um valor de tipo de interface pode conter qualquer valor que implementa esses métodos.
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implementa Abser
	a = &v // a *Vertex implementa Abser

	// a = v // Não implementa Abser, pois nao é um ponteiro

	fmt.Println(a)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
