package main

import (
	"github.com/Godesde0/variables"
	"fmt"
)

func main(){
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}