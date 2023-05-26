package variables

import (
	"fmt"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables(){
	Nombre = "Erick"
	Estado = true
	Sueldo = 5700
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}