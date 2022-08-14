/*                 _____________________________________________________________________________________________________
                   |                                        EJERCICIO #4                                               |
                   |                                                                                                   |
	           |    Escriba un programa que utilice una estructura y un arreglo/slice para almacenar en memoria    |
                   |    un inventario de una tienda que vende zapatos. Cada zapato debe contar con la información      |
                   |    de su modelo (marca), su precio y la talla del mismo que debe ir únicamente del 34 al 44.      |
                   |    La estructura debe de llamarse "calzado". El programa debe de poder almacenar la información   |
                   |    "quemada" para 10+ zapatos del inventario y un stock de N cantidad de unidades de dicho        |
                   |    zapato (quiere decir por ejemplo: que puede existir en el inventario el zapato de marca Nike   |
                   |    talla 42 y que cuesta 50000 colones, pero además que puede exirtir no solo un par de esos,     |
                   |    sino N pares del mismo zapato).                                                                |
                   |                                                                                                   |
                   |    Haga una función que venda zapatos (eliminando del stock cada vez que haya venta e indicando   |
                   |    que no se puede vender porque ya no hay stock) y pruebelo dentro del main haciendo las         |
                   |    ventas que usted considere para demostrar su funcionamiento.                                   |
                   |___________________________________________________________________________________________________|
*/

package main

import (
	"fmt"
)

type calzado struct {
	marca  string
	precio int
	talla  int
	stock  int
}

type listaCalzado []calzado

var lCalzado listaCalzado

func (l *listaCalzado) agregarCalzado(marca string, precio int, talla int, stock int) {

	var zap = l.buscarCalzado(marca, talla)
	if zap != -1 {
		(*l)[zap].stock = (*l)[zap].stock + stock
		fmt.Println("\n--------------------------------------------------------------------------")
		fmt.Println("\n             ********** SOLICITUD DE CAMBIO EN CALZADO **********")
		fmt.Println("\n--- Información de cambio ---",
			"\n   Marca: ", (*l)[zap].marca, "\n   Precio: ", (*l)[zap].precio, "\n   Talla: ", (*l)[zap].talla, "\n   Nuevo Stock: ", (*l)[zap].stock)
		fmt.Println("\n ESTADO: Cambiado correctamente\n")
		fmt.Println("\n--------------------------------------------------------------------------")
	} else {
		if 34 <= talla && talla <= 44 {
			*l = append(*l, calzado{marca: marca, precio: precio, talla: talla, stock: stock})
		} else {
			fmt.Println("\n--------------------------------------------------------------------------")
			fmt.Println("\n               ********** SOLICITUD DE NUEVO CALZADO **********\n")
			fmt.Println("\n ESTADO: No se puede añadir. La talla debe ser únicamente del 34 a 44.\n")
			fmt.Println("\n--------------------------------------------------------------------------")

		}
	}

}

func (l *listaCalzado) buscarCalzado(marca string, talla int) int {
	var result = -1
	var i int

	for i = 0; i < len(*l); i++ {
		if (*l)[i].marca == marca && (*l)[i].talla == talla {
			result = i
		}
	}
	return result
}

func (l *listaCalzado) venderCalzado(marca string, talla int, stock int) {
	var zapa = l.buscarCalzado(marca, talla)
	if zapa != -1 && stock > 0 {
		if (*l)[zapa].stock >= stock { //Sí la cantidad del calzado que hay en la lista es mayor o igual que la cantidad que nos piden
			(*l)[zapa].stock = (*l)[zapa].stock - stock //Le quitamos a la cantidad que hay en la lista la cantidad que nos solicitan
			fmt.Println("\n--------------------------------------------------------------------------")
			fmt.Println("\n            ********** SOLICITUD DE VENTA **********\n")
			fmt.Println("\n--- Información de venta ---",
				"\n   Marca: ", (*l)[zapa].marca, "\n   Precio: ", (*l)[zapa].precio, "\n   Talla: ", (*l)[zapa].talla)
			fmt.Println("\n ESTADO: Se vendió con éxito.\n")
			fmt.Println("\n--------------------------------------------------------------------------")

		} else {
			fmt.Println("\n--------------------------------------------------------------------------")
			fmt.Println("\n            ********** SOLICITUD DE VENTA **********\n")
			fmt.Println("\n--- Información de venta ---",
				"\n   Marca: ", (*l)[zapa].marca, "\n   Precio: ", (*l)[zapa].precio, "\n   Talla: ", (*l)[zapa].talla)
			fmt.Println("\nESTADO: No se puede vender; No hay suficiente stock.\n")

			fmt.Println("\n--------------------------------------------------------------------------")

		}
		if (*l)[zapa].stock == 0 { //Sí la cantidad del producto es igual a cero
			fmt.Println("\n--------------------------------------------------------------------------")
			fmt.Println("\n            ********** INFORMACIÓN DEL CALZADO **********")
			fmt.Println("\nEl producto ", (*l)[zapa].marca, " de talla ", (*l)[zapa].talla, " ha sido eliminado; \nNo hay más existencias.")
			fmt.Println("\n--------------------------------------------------------------------------")
			(*l) = append((*l)[:zapa], (*l)[zapa+1:]...) //Nos borra el calzado
		}
	}

}

func llenarInventario() {
	lCalzado.agregarCalzado("Tommy", 39000, 35, 20)
	lCalzado.agregarCalzado("Nike", 48000, 38, 17)
	lCalzado.agregarCalzado("Adidas", 50300, 40, 12)
	lCalzado.agregarCalzado("Sperry", 52000, 37, 16)
	lCalzado.agregarCalzado("Tommy", 53450, 34, 10)
	lCalzado.agregarCalzado("Adidas", 65856, 42, 25)
	lCalzado.agregarCalzado("Adidas", 67596, 44, 19)
	lCalzado.agregarCalzado("Nike", 45926, 36, 14)
	lCalzado.agregarCalzado("Sperry", 63789, 35, 18)
	lCalzado.agregarCalzado("Tommy", 78956, 41, 22)
	lCalzado.agregarCalzado("Sperry", 45896, 34, 8)

}

func (l *listaCalzado) imprimirInventario() {

	fmt.Println("\n--------------------------------------------------------------------------")
	fmt.Println("\n                      LISTA DE PRODUCTOS            \n")

	for zap := range *l {
		fmt.Println("    Marca:", (*l)[zap].marca, " - Precio: ₡", (*l)[zap].precio, " - Talla:", (*l)[zap].talla, " - Stock:", (*l)[zap].stock)
	}
	fmt.Println("\n--------------------------------------------------------------------------")

}

func main() {

	//MOSTRAMOS EL INVENTARIO DEL CALZADO
	llenarInventario()
	lCalzado.imprimirInventario()

	//AGREGAMOS UN NUEVO CALZADO AL INVENTARIO
	//Se muestra nuevamente el inventario para comprobarlo
	lCalzado.agregarCalzado("Converse", 64598, 37, 6)
	lCalzado.imprimirInventario()

	//AGREGAMOS UN NUEVO CALZADO AL INVENTARIO
	//Se muestra un mensaje "No se puede. La talla debe ser del 34 al 44" y muestra nuevamente el inventario para comprobar que no se añadio
	lCalzado.agregarCalzado("Converse", 54789, 30, 6)
	lCalzado.imprimirInventario()

	//CAMBIAMOS EL STOCK DE UN CALZADO
	//Se muestra información del cambio del zapato y  muestra nuevamente el inventario para comprobar el cambio
	lCalzado.agregarCalzado("Tommy", 78956, 41, 3)
	lCalzado.imprimirInventario()

	//VENDEMOS UN CALZADO
	//Se muestra información de la solicitud de venta y un mensaje "vendido con éxito". Además muestra nuevamenete el inventario para ver los cambios
	lCalzado.venderCalzado("Nike", 38, 3)
	lCalzado.imprimirInventario()

	//VENDEMOS UN CALZADO
	//Se muestra información de la solicitud de venta y un mensaje "No se puede. No hay suficiente stock".
	//Además muestra nuevamenete el inventario para ver los cambios
	lCalzado.venderCalzado("Adidas", 42, 30)
	lCalzado.imprimirInventario()

	//VENDEMOS UN CALZADO
	//Se muestra información de la solicitud de venta y un mensaje "vendido con éxito".
	//También, muestra un mensaje indicando que se elimino el producto porque no hay más stock.
	//Además muestra nuevamenete el inventario para ver los cambios.
	lCalzado.venderCalzado("Sperry", 37, 16)
	lCalzado.imprimirInventario()

}

/*
            RESPUESTA DEL CÓDIGO

GOROOT=C:\Program Files\Go #gosetup
GOPATH=C:\Users\Melany\Desktop\LENGUAJES DE PROGRAMACIÓN\TRABAJOS EXTRACLASE\S3-Venta_Calzado\.idea #gosetup
"C:\Program Files\Go\bin\go.exe" build -o C:\Users\Melany\AppData\Local\Temp\GoLand\___go_build_S3_Venta_Calzado.exe S3-Venta_Calzado #gosetup
C:\Users\Melany\AppData\Local\Temp\GoLand\___go_build_S3_Venta_Calzado.exe

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Marca: Tommy  - Precio: ₡ 39000  - Talla: 35  - Stock: 20
    Marca: Nike  - Precio: ₡ 48000  - Talla: 38  - Stock: 17
    Marca: Adidas  - Precio: ₡ 50300  - Talla: 40  - Stock: 12
    Marca: Sperry  - Precio: ₡ 52000  - Talla: 37  - Stock: 16
    Marca: Tommy  - Precio: ₡ 53450  - Talla: 34  - Stock: 10
    Marca: Adidas  - Precio: ₡ 65856  - Talla: 42  - Stock: 25
    Marca: Adidas  - Precio: ₡ 67596  - Talla: 44  - Stock: 19
    Marca: Nike  - Precio: ₡ 45926  - Talla: 36  - Stock: 14
    Marca: Sperry  - Precio: ₡ 63789  - Talla: 35  - Stock: 18
    Marca: Tommy  - Precio: ₡ 78956  - Talla: 41  - Stock: 22
    Marca: Sperry  - Precio: ₡ 45896  - Talla: 34  - Stock: 8

--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Marca: Tommy  - Precio: ₡ 39000  - Talla: 35  - Stock: 20
    Marca: Nike  - Precio: ₡ 48000  - Talla: 38  - Stock: 17
    Marca: Adidas  - Precio: ₡ 50300  - Talla: 40  - Stock: 12
    Marca: Sperry  - Precio: ₡ 52000  - Talla: 37  - Stock: 16
    Marca: Tommy  - Precio: ₡ 53450  - Talla: 34  - Stock: 10
    Marca: Adidas  - Precio: ₡ 65856  - Talla: 42  - Stock: 25
    Marca: Adidas  - Precio: ₡ 67596  - Talla: 44  - Stock: 19
    Marca: Nike  - Precio: ₡ 45926  - Talla: 36  - Stock: 14
    Marca: Sperry  - Precio: ₡ 63789  - Talla: 35  - Stock: 18
    Marca: Tommy  - Precio: ₡ 78956  - Talla: 41  - Stock: 22
    Marca: Sperry  - Precio: ₡ 45896  - Talla: 34  - Stock: 8
    Marca: Converse  - Precio: ₡ 64598  - Talla: 37  - Stock: 6

--------------------------------------------------------------------------

--------------------------------------------------------------------------

               ********** SOLICITUD DE NUEVO CALZADO **********


 ESTADO: No se puede añadir. La talla debe ser únicamente del 34 a 44.


--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Marca: Tommy  - Precio: ₡ 39000  - Talla: 35  - Stock: 20
    Marca: Nike  - Precio: ₡ 48000  - Talla: 38  - Stock: 17
    Marca: Adidas  - Precio: ₡ 50300  - Talla: 40  - Stock: 12
    Marca: Sperry  - Precio: ₡ 52000  - Talla: 37  - Stock: 16
    Marca: Tommy  - Precio: ₡ 53450  - Talla: 34  - Stock: 10
    Marca: Adidas  - Precio: ₡ 65856  - Talla: 42  - Stock: 25
    Marca: Adidas  - Precio: ₡ 67596  - Talla: 44  - Stock: 19
    Marca: Nike  - Precio: ₡ 45926  - Talla: 36  - Stock: 14
    Marca: Sperry  - Precio: ₡ 63789  - Talla: 35  - Stock: 18
    Marca: Tommy  - Precio: ₡ 78956  - Talla: 41  - Stock: 22
    Marca: Sperry  - Precio: ₡ 45896  - Talla: 34  - Stock: 8
    Marca: Converse  - Precio: ₡ 64598  - Talla: 37  - Stock: 6

--------------------------------------------------------------------------

--------------------------------------------------------------------------

             ********** SOLICITUD DE CAMBIO EN CALZADO **********

--- Información de cambio ---
   Marca:  Tommy
   Precio:  78956
   Talla:  41
   Nuevo Stock:  25

 ESTADO: Cambiado correctamente


--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Marca: Tommy  - Precio: ₡ 39000  - Talla: 35  - Stock: 20
    Marca: Nike  - Precio: ₡ 48000  - Talla: 38  - Stock: 17
    Marca: Adidas  - Precio: ₡ 50300  - Talla: 40  - Stock: 12
    Marca: Sperry  - Precio: ₡ 52000  - Talla: 37  - Stock: 16
    Marca: Tommy  - Precio: ₡ 53450  - Talla: 34  - Stock: 10
    Marca: Adidas  - Precio: ₡ 65856  - Talla: 42  - Stock: 25
    Marca: Adidas  - Precio: ₡ 67596  - Talla: 44  - Stock: 19
    Marca: Nike  - Precio: ₡ 45926  - Talla: 36  - Stock: 14
    Marca: Sperry  - Precio: ₡ 63789  - Talla: 35  - Stock: 18
    Marca: Tommy  - Precio: ₡ 78956  - Talla: 41  - Stock: 25
    Marca: Sperry  - Precio: ₡ 45896  - Talla: 34  - Stock: 8
    Marca: Converse  - Precio: ₡ 64598  - Talla: 37  - Stock: 6

--------------------------------------------------------------------------

--------------------------------------------------------------------------

            ********** SOLICITUD DE VENTA **********


--- Información de venta ---
   Marca:  Nike
   Precio:  48000
   Talla:  38

 ESTADO: Se vendió con éxito.


--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Marca: Tommy  - Precio: ₡ 39000  - Talla: 35  - Stock: 20
    Marca: Nike  - Precio: ₡ 48000  - Talla: 38  - Stock: 14
    Marca: Adidas  - Precio: ₡ 50300  - Talla: 40  - Stock: 12
    Marca: Sperry  - Precio: ₡ 52000  - Talla: 37  - Stock: 16
    Marca: Tommy  - Precio: ₡ 53450  - Talla: 34  - Stock: 10
    Marca: Adidas  - Precio: ₡ 65856  - Talla: 42  - Stock: 25
    Marca: Adidas  - Precio: ₡ 67596  - Talla: 44  - Stock: 19
    Marca: Nike  - Precio: ₡ 45926  - Talla: 36  - Stock: 14
    Marca: Sperry  - Precio: ₡ 63789  - Talla: 35  - Stock: 18
    Marca: Tommy  - Precio: ₡ 78956  - Talla: 41  - Stock: 25
    Marca: Sperry  - Precio: ₡ 45896  - Talla: 34  - Stock: 8
    Marca: Converse  - Precio: ₡ 64598  - Talla: 37  - Stock: 6

--------------------------------------------------------------------------

--------------------------------------------------------------------------

            ********** SOLICITUD DE VENTA **********


--- Información de venta ---
   Marca:  Adidas
   Precio:  65856
   Talla:  42

ESTADO: No se puede vender; No hay suficiente stock.


--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Marca: Tommy  - Precio: ₡ 39000  - Talla: 35  - Stock: 20
    Marca: Nike  - Precio: ₡ 48000  - Talla: 38  - Stock: 14
    Marca: Adidas  - Precio: ₡ 50300  - Talla: 40  - Stock: 12
    Marca: Sperry  - Precio: ₡ 52000  - Talla: 37  - Stock: 16
    Marca: Tommy  - Precio: ₡ 53450  - Talla: 34  - Stock: 10
    Marca: Adidas  - Precio: ₡ 65856  - Talla: 42  - Stock: 25
    Marca: Adidas  - Precio: ₡ 67596  - Talla: 44  - Stock: 19
    Marca: Nike  - Precio: ₡ 45926  - Talla: 36  - Stock: 14
    Marca: Sperry  - Precio: ₡ 63789  - Talla: 35  - Stock: 18
    Marca: Tommy  - Precio: ₡ 78956  - Talla: 41  - Stock: 25
    Marca: Sperry  - Precio: ₡ 45896  - Talla: 34  - Stock: 8
    Marca: Converse  - Precio: ₡ 64598  - Talla: 37  - Stock: 6

--------------------------------------------------------------------------

--------------------------------------------------------------------------

            ********** SOLICITUD DE VENTA **********


--- Información de venta ---
   Marca:  Sperry
   Precio:  52000
   Talla:  37

 ESTADO: Se vendió con éxito.


--------------------------------------------------------------------------

--------------------------------------------------------------------------

            ********** INFORMACIÓN DEL CALZADO **********

El producto  Sperry  de talla  37  ha sido eliminado;
No hay más existencias.

--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Marca: Tommy  - Precio: ₡ 39000  - Talla: 35  - Stock: 20
    Marca: Nike  - Precio: ₡ 48000  - Talla: 38  - Stock: 14
    Marca: Adidas  - Precio: ₡ 50300  - Talla: 40  - Stock: 12
    Marca: Tommy  - Precio: ₡ 53450  - Talla: 34  - Stock: 10
    Marca: Adidas  - Precio: ₡ 65856  - Talla: 42  - Stock: 25
    Marca: Adidas  - Precio: ₡ 67596  - Talla: 44  - Stock: 19
    Marca: Nike  - Precio: ₡ 45926  - Talla: 36  - Stock: 14
    Marca: Sperry  - Precio: ₡ 63789  - Talla: 35  - Stock: 18
    Marca: Tommy  - Precio: ₡ 78956  - Talla: 41  - Stock: 25
    Marca: Sperry  - Precio: ₡ 45896  - Talla: 34  - Stock: 8
    Marca: Converse  - Precio: ₡ 64598  - Talla: 37  - Stock: 6

--------------------------------------------------------------------------

Process finished with the exit code 0


*/
