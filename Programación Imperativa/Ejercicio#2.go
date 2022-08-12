/*                 _____________________________________________________________________________________________________
                   |                                        EJERCICIO #2                                               |
                   |                                                                                                   |
	               |     Escriba el programa más eficiente que pueda para imprimir en pantalla la siguiente figura:    |
                   |                                                                                                   |
                   |                                                                                                   |
				   |											 *                                                     |
                   |                                          *  *  *                                                  |
                   |                                       *  *  *  *  *                                               |
                   |                                          *  *  *                                                  |
                   |                                             *                                                     |
                   |                                                                                                   |
                   |     Para dicho fin, escriba y use una función que reciba de parámetro la cantidad de elementos    |
                   |      de la línea del centro, la cual debe ser impar positiva.                                     |
                   |___________________________________________________________________________________________________|
*/

package main

import "fmt"

func imprimirRombo() {

	h := 3 //La altura que queramos ponerle

	/*
	   1. El "i" representa la posición de las filas
	   2. El "j" representa los espacios en blanco en la posición horizontal (posición de cada fila)
	   3. El "k" representa los asteriscos en la posición horizontal (posción en cada fila )
	*/

	//ARMAMOS EL TRIÁNGULO SUPERIOR
	for i := 1; h-1 >= i; i++ {
		for j := h; j >= i; j-- {
			fmt.Print("  ")

		}
		for k := 1; 2*i-1 >= k; k++ {
			fmt.Print(" *")
		}
		fmt.Println(" ")
	}

	//ARMAMOS EL TRIÁNGULO INFERIOR
	for i := 1; h >= i; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("  ")
		}
		for k := 2*h - 1; 2*i-1 <= k; k-- {
			fmt.Print(" *")
		}
		fmt.Println(" ")
	}

	fmt.Println("\n")
}

func main() {
	fmt.Println("\n --------------------------- ROMBO --------------------------- \n")

	imprimirRombo()

	fmt.Println("\n ------------------------------------------------------------- \n")
}

/*
   RESPUESTA DEL CÓDIGO

GOROOT=C:\Program Files\Go #gosetup
GOPATH=C:\Users\Melany\Desktop\LENGUAJES DE PROGRAMACIÓN\TRABAJOS EXTRACLASE\S3-Figura_Rombo\.idea #gosetup
"C:\Program Files\Go\bin\go.exe" build -o C:\Users\Melany\AppData\Local\Temp\GoLand\___3go_build_main_go.exe "C:\Users\Melany\Desktop\LENGUAJES DE PROGRAMACIÓN\TRABAJOS EXTRACLASE\S3-Figura_Rombo\main.go" #gosetup
C:\Users\Melany\AppData\Local\Temp\GoLand\___3go_build_main_go.exe

 --------------------------- ROMBO ---------------------------

       *
     * * *
   * * * * *
     * * *
       *



 -------------------------------------------------------------


Process finished with the exit code 0


*/
