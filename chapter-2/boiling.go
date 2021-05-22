// Boiling exibe o ponto de ebuição da água
package main

import "fmt"

// Constante visível a nível de pacote porque seu nome começa com letra minúscula
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}
