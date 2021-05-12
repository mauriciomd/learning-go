// Echo2 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// Range produz um par de valores: index, valor do elemento
	// _ => identificador vazio, utilizado quando a sintaxe exige uma variável mas ela não será utilizada
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
