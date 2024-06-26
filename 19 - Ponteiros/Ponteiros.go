// Escolhendo um valor ou ponteiro receptor:

// Há duas razões para usar um receptor de ponteiro.

// O primeiro é para que o método possa modificar o valor que os seus pontos de receptor referenciam. = *

// A segunda é para evitar copiar o valor de cada chamada de método. Isto pode ser mais eficaz se o receptor for uma estrutura grande, por exemplo. =  &

// Neste exemplo, tanto Scale e Abs estão com tipo de receptor *Vertex, mesmo que o método Abs não precise modificar seu receptor.

// Em geral, todos os métodos em um determinado tipo devem ter o valor ou ponteiro receptores, mas não uma mistura de ambos. (Vamos ver porque nas próximas páginas).
package main

import (
	"fmt"
	"math"
)

type Vertexx struct {
	b, n float64
}

func (v *Vertexx) Scales(f float64) {
	v.b = v.b * f
	v.n = v.n * f
}

func (v *Vertexx) Absa() float64 {
	return math.Sqrt(v.b*v.b + v.n*v.n)
}

func teste() {
	v := &Vertexx{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Absa())
	v.Scales(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Absa())
}
