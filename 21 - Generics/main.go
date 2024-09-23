package main

//Constraints são restrições que vc pode colocar no generics para limitar os tipos que podem ser usados
type Number interface {
	int | float64
}

// Tipos Generics são formas mais práticas de vc não ter que criar duas funções para tipos diferentes
// Utilizando o generics, vc pode criar uma função que aceita qualquer tipo de dado, sem precisar criar uma função para cada tipo.
func Soma[T Number](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}
	return sum
}

// Exemplo de uso:
// A função Soma aceita tanto um map de int quanto um map de float64
func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	println(Soma(m))

	m2 := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	println(Soma(m2))
}
