/*                 _____________________________________________________________________________________________________
                   |                                        EJERCICIO #1                                               |
                   |                                                                                                   |
	               |     Haga un programa que cuente e indique el número de caracteres, el número de palabras          |
                   |     y el número de líneas de un texto cualquiera (obtenido de cualquier forma que considere).     |
                   |     Considere hacer el cálculo de las tres variables en el mismo proceso                          |
                   |                                                                                                   |
                   |___________________________________________________________________________________________________|
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string
	var characters []string
	var words []string
	var lines []string

	frase = "Ella tiene el poder de destruirme, y la \n" +
		"verdad no me importa, ser destruido \n" +
		"por ella sería un jodido privilegio"

	characters = append(strings.Split(frase, ""))
	words = append(strings.Split(frase, " "))
	lines = append(strings.Split(frase, "\n"))

	fmt.Println("\n                 FRASE\n"+
		"----------------------------------------------\n", frase,
		"\n----------------------------------------------\n")

	fmt.Println("Caracteres: ", len(characters))
	fmt.Println("Palabras: ", len(words))
	fmt.Println("Lineas: ", len(lines))

}
