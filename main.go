package main

import (
	"fmt"

	"github.com/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	variables.Nombre = "otro nombre"
	fmt.Println(variables.Nombre)

	estado, texto := variables.ConviertoaTexto(12345)
	fmt.Println(estado, texto)

}
