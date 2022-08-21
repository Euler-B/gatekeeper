package main 

import (
	"fmt"
)

func main() {
	var nombre string
	var edad   int

	fmt.Print("Buenos Dias, me indica su nombre: \n --> ")
    fmt.Scanf("%s\n", &nombre)

	fmt.Print("Ademas quisiera que me dijera su edad: \n --> ")
	fmt.Scanf("%d", &edad)

	switch {
	case edad < 18:
		fmt.Printf("%s no cumples con la edad reglamentaria para entrar a este lugar. \n", nombre)
	
	case edad > 60:
		fmt.Printf("Don %s, deberia explorar otras opciones para distraerse. \n", nombre)
	default:
		fmt.Printf("Bienvenido Sr. %s, en un instante le atenderemos. \n", nombre)
	}	
}