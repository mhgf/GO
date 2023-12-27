package main

import (
	"fmt"
)

const ebulicaoK = 373.0

func main() {

	ebulicaoC := ebulicaoK - 273.0

	fmt.Printf("A temperatura de ebuliçao da água em K = %G, e em °C = %G.", ebulicaoK, ebulicaoC)

}
