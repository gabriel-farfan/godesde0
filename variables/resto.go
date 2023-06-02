package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time


func RestoVariables() {
	Nombre = "Juan"
	Estado = true
	Sueldo = 1000.50
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)

}

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}