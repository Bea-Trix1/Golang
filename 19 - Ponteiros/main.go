package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Métodos com receptores ponteiro podem modificar o valor que o receptor aponta (como Scale faz aqui).
// Esta é uma prática muito utilizada, e geralmente é ainda mais comum que receptores normais (com valor em cópia).
// Podemos entender os ponteiros como uma referência à variável original, ao invés de receber uma cópia daquele valor.
func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
