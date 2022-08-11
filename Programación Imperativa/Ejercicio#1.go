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



/*
            RESPUESTA DEL CÓDIGO

GOROOT=C:\Program Files\Go #gosetup
GOPATH=C:\Users\Melany\Desktop\LENGUAJES DE PROGRAMACIÓN\TRABAJOS EXTRACLASE\S3-Contar_Frase\.idea #gosetup
"C:\Program Files\Go\bin\go.exe" build -o C:\Users\Melany\AppData\Local\Temp\GoLand\___go_build_main_go.exe "C:\Users\Melany\Desktop\LENGUAJES DE PROGRAMACIÓN\TRABAJOS EXTRACLASE\S3-Contar_Frase\main.go" #gosetup
C:\Users\Melany\AppData\Local\Temp\GoLand\___go_build_main_go.exe

                 FRASE
----------------------------------------------
 Ella tiene el poder de destruirme, y la
verdad no me importa, ser destruido
por ella sería un jodido privilegio
----------------------------------------------

Caracteres:  113
Palabras:  20
Lineas:  3

Process finished with the exit code 0




*/
