/*                 _____________________________________________________________________________________________________
                   |                                        EJERCICIO #3                                               |
                   |                                                                                                   |
	           |    Escriba una función que permita rotar una secuencia de elementos N posiciones a la izquierda   |
                   |    o a la derecha según sea el caso en función al parámetro que se reciba. Los parámetros deeben  |
                   |    ser un puntero a un arreglo previamente creado, la cantidad de moviento de rotación y la       |
                   |    clasificación (0 =izquierda y 1 =derecha)                                                      |
                   |                                                                                                   |
                   |    A partir de esta función, escriba un programa que haga varias totaciones cualquiera de una     |
                   |    secuencia de elementos previamente creada que sea inmutable. Al final debe imprimirse el       |
                   |    resultado de cada rotación además de la secuencia original.                                    |
                   |                                                                                                   |
                   |    Un ejemplo de rotación puede ser el siguiente:                                                 |
                   |                                                                                                   |
                   |    Secuencia Original = [a, b, c, d, e, f, g, h]                                                  |
                   |    Cantidad de rotaciones = 3                                                                     |
                   |    Dirección: izquierda                                                                           |
                   |    Secuencia final rotada: [d, e, f, g, h, a, b, c] <-- Nótese que cada iteración, el elemento    |
                   |    más a la izquierda pasó a forma parte del final de la secuencia, si la rotación fuera a la     |
                   |    derecha, sería lo contrario.                                                                   |
                   |___________________________________________________________________________________________________|
*/

package main

import "fmt"

func rotarSecuencia(l *[]int, rotaciones int, direccion int) {

	if direccion == 1 {
		rotaciones = rotaciones % len((*l))
		(*l) = append((*l)[rotaciones:], (*l)[0:rotaciones]...)
		fmt.Println("   Cantidad de rotaciones: ", rotaciones)
		fmt.Println("   Dirección a rotar: ", direccion)

		fmt.Println("     -> RESPUESTA: Rotación por la derecha: ", (*l))
	} else {
		result := append((*l)[len((*l))-rotaciones:], (*l)[:len((*l))-rotaciones]...)

		for i := 0; i < len((*l)); i++ {
			(*l)[i] = result[i]
		}
		fmt.Println("   Cantidad de rotaciones: ", rotaciones)
		fmt.Println("   Dirección a rotar: ", direccion)

		fmt.Println("     -> RESPUESTA: Rotación por la izquierda: ", (*l))
	}

}

func main() {
	secuenciaOriginal1 := []int{1, 2, 3, 4, 5}
	secuenciaOriginal2 := []int{10, 20, 30, 40, 50, 60, 70, 80}

	fmt.Println("\n----------------------------------------------------------------------------------------------\n",
		"                       ROTACIÓN DE UNA SECUENCIA\n")

	fmt.Println("Secuencia Original: ", secuenciaOriginal1)
	rotarSecuencia(&secuenciaOriginal1, 3, 1)

	fmt.Println("\nSecuencia Original: ", secuenciaOriginal2)
	rotarSecuencia(&secuenciaOriginal2, 5, 0)

	fmt.Println("\n----------------------------------------------------------------------------------------------\n")

}

/*
            RESPUESTA DEL CÓDIGO

GOPATH=C:\Users\Melany\Desktop\LENGUAJES DE PROGRAMACIÓN\TRABAJOS EXTRACLASE\S3-Rotar_Secuencia\.idea #gosetup
"C:\Program Files\Go\bin\go.exe" build -o C:\Users\Melany\AppData\Local\Temp\GoLand\___go_build_S3_Rotar_Secuencia.exe S3-Rotar_Secuencia #gosetup
C:\Users\Melany\AppData\Local\Temp\GoLand\___go_build_S3_Rotar_Secuencia.exe

----------------------------------------------------------------------------------------------
                        ROTACIÓN DE UNA SECUENCIA


Secuencia Original:  [1 2 3 4 5]
   Cantidad de rotaciones:  3
   Dirección a rotar:  1
     -> RESPUESTA: Rotación por la derecha:  [4 5 1 2 3]

Secuencia Original:  [10 20 30 40 50 60 70 80]
   Cantidad de rotaciones:  5
   Dirección a rotar:  0
     -> RESPUESTA: Rotación por la izquierda:  [40 50 60 70 80 10 20 30]

----------------------------------------------------------------------------------------------



Process finished with the exit code 0


*/
