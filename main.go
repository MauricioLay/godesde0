package main

import (
	"fmt"

	"github.com/godesde0/ejericios01"
)

func main() {

	//dato, res := ejercicios.Ejercicios01("100")
	//variables.MuestroEnteros()
	dato, salida := ejericios01.Ejercicios01("2")
	fmt.Println("dato:", dato)
	fmt.Println("salida:", salida)
	//variables.MuestroEnteros()
	/*variables.RestoVariables()
	variables.Nombre = "otro nombre"
	fmt.Println(variables.Nombre)

	estado, texto := variables.ConviertoaTexto(12345)
	fmt.Println(estado, texto)
	fmt.Println(runtime.GOOS)
	if os := runtime.GOOS; os == "Linux." || os == "darwin" {
		fmt.Println("no es un windows")
	} else {
		fmt.Println("es un windows")

	}*/

}
