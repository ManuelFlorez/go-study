package main

import (
	"fmt"
	"strconv"
)

var Numero int
var Texto string
var Status bool = true

func main() {
	numero2, numero3, numero4, status, texto := 2, 5, 67, false, "Hola"

	var moneda int64 = 0

	numero2 = int(moneda)
	// texto = fmt.Sprintf("%d", moneda)
	texto = strconv.Itoa(int(moneda))
	fmt.Println(texto)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(status)

	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println(Status)
}
