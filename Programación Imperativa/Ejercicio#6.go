/*
	               _____________________________________________________________________________________________________
	               |                                        EJERCICIO #6                                               |
	               |                                                                                                   |
		       |    Amplie el funcionamiento del ejercicio de Productos visto en clase para que el programa        |
                       |    ahora permita:                                                                                 |
                       |                                                                                                   |
                       |    a. A partir de la lista de productos con mínimas existencias de inventario, ampliar dicho      |
                       |       inventario con la compra de más unidades de dicho producto hasta que cumplan con el         |
                       |       mínimo establecido de manera constante. Se sugiere crear una función denominada             |
                       |      "aumentarInventarioDeMinimas(listaMínimos)".                                                 |
	               |                                                                                                   |
	               |    b. Crear una función que ordene la lista de productos usando como llave para el                |
                       |       ordenamiento cualquiera de los elementos de la estructura producto. La lista/slice          |
                       |       debe quedar modificada al finalizar el método. Se solicita investigar y hacer uso de        |
                       |       alguna función predefinida de alguna librería del lenguaje Go.                              |
	               |___________________________________________________________________________________________________|
*/
package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var produ = l.buscarProducto(nombre)
	if produ != -1 {
		(*l)[produ].cantidad = (*l)[produ].cantidad + cantidad
		if (*l)[produ].precio != precio {
			(*l)[produ].precio = precio
			fmt.Println("\n--------------------------------------------------------------------------")
			fmt.Println("\n             ********** SOLICITUD DE CAMBIO EN PRODUCTO **********")
			fmt.Println("\n--- Información de cambio ---",
				"\n   Nombre: ", (*l)[produ].nombre, "\n   Nuevo Precio: ", (*l)[produ].precio, "\n   Nueva Cantidad: ", (*l)[produ].cantidad)
			fmt.Println("\n ESTADO: Cambiado correctamente\n")
			fmt.Println("\n--------------------------------------------------------------------------")

		}
	} else {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})

	}

}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant { //Sí la cantidad del producto que hay en la lista es mayor o igual que la cantidad que nos piden
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant //Le quitamos a la cantidad que hay en la lista la cantidad que nos solicitan
			fmt.Println("\n--------------------------------------------------------------------------")
			fmt.Println("\n            ********** SOLICITUD DE VENTA **********\n")
			fmt.Println("\n--- Información de venta ---",
				"\n   Nombre: ", (*l)[prod].nombre, "\n   Precio: ", (*l)[prod].precio)
			fmt.Println("\n ESTADO: Se vendió con éxito.\n")
			fmt.Println("\n--------------------------------------------------------------------------")

		} else {
			fmt.Println("\n--------------------------------------------------------------------------")
			fmt.Println("\n            ********** SOLICITUD DE VENTA **********\n")
			fmt.Println("\n--- Información de venta ---",
				"\n   Nombre: ", (*l)[prod].nombre, "\n   Precio: ", (*l)[prod].precio)
			fmt.Println("\nESTADO: No se puede vender; No hay suficiente cantidad.\n")

			fmt.Println("\n--------------------------------------------------------------------------")
		}
		if (*l)[prod].cantidad == 0 { //Sí la cantidad del producto es igual a cero
			fmt.Println("\n--------------------------------------------------------------------------")
			fmt.Println("\n            ********** INFORMACIÓN DEL PRODUCTO **********")
			fmt.Println("\nEl producto ", (*l)[prod].nombre, " del precio ", (*l)[prod].precio, " ha sido eliminado; \nNo hay más existencias.")
			fmt.Println("\n--------------------------------------------------------------------------")
			(*l) = append((*l)[:prod], (*l)[prod+1:]...) //Nos borra el producto

		}
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	}

}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("cafe", 12, 4500)
	lProductos.agregarProducto("sal", 1, 700)
	lProductos.agregarProducto("azucar", 7, 1800)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var listaMinimos []producto //nueva lista
	for prod := range *l {      //Recorre la lista original
		if (*l)[prod].cantidad <= existenciaMinima { //Pregunta sí la cantidad del producto que está en la lista tiene la existencia minima
			listaMinimos = append(listaMinimos, (*l)[prod]) //Añade ese producto a la lista de minimos
		}
	}

	return listaMinimos
}

func (l *listaProductos) aumentarInventarioDeMinimos() {
	var listaMinimosInventario = l.listarProductosMínimos()

	fmt.Println("\n--------------------------------------------------------------------------\n")
	fmt.Println("                  SE AUMENTÓ EL  INVENTARIO DE MÍNIMAS EXISTECNIAS \n")
	for i := 0; i < len(listaMinimosInventario); i++ { //Al recorrer el "len" de la lista de los productos mínimos
		(*l)[i].cantidad += existenciaMinima - (*l)[i].cantidad //A la cantidad de la lista le sumamos los 10 de la existencia mínima,
		//pero le restamos la cantidad que el producto tenía anteriormente, por lo cual, nos aumentaria las unidades que faltan para que cumpla con el mínimo establecido

	}
	fmt.Println("--------------------------------------------------------------------------\n")

}

func (l *listaProductos) ordenarLista() {
	fmt.Println("--------------------------------------------------------------------------\n")
	fmt.Println("                SE ORDENÓ LA LISTA DE PRODUCTOS\n")
	sort.SliceStable(*l, func(i, j int) bool { //El sort nos ayuda a ordenar
		return (*l)[i].cantidad < (*l)[j].cantidad
	})
	fmt.Println("\n--------------------------------------------------------------------------\n")
}

func (l *listaProductos) imprimirInventario() {

	fmt.Println("\n--------------------------------------------------------------------------")
	fmt.Println("\n                      LISTA DE PRODUCTOS            \n")

	for pro := range *l {
		fmt.Println("    Nombre:", (*l)[pro].nombre, " - Precio: ₡", (*l)[pro].precio, " - Cantidad:", (*l)[pro].cantidad)
	}
	fmt.Println("\n--------------------------------------------------------------------------")

}

func main() {
	//MOSTRAMOS EL INVENTARIO DE LOS PRODUCTOS
	llenarDatos()
	lProductos.imprimirInventario()

	//AGREGAMOS NUEVO PRODUCTO
	lProductos.agregarProducto("mantequilla", 20, 1000)
	lProductos.imprimirInventario()

	//CAMBIEMOS EL PRECIO Y LA CANTIDAD DEL PRODUCTO
	//El nuevo producto a agragar ya existe, pero el precio es diferente, por lo cual se cambia en la lista de productos,
	//ademas, sumamos la cantidad que queremos agregar con la cantidad anterior
	lProductos.agregarProducto("cafe", 4, 1200)
	lProductos.imprimirInventario()

	//VENDEMOS EL PRODUCTO SIN PROBLEMA
	//Se muestra información de la solicitud de venta y un mensaje "vendido con éxito". Además muestra nuevamenete el inventario para ver los cambios
	lProductos.venderProducto("arroz", 4)
	lProductos.imprimirInventario()

	//VENDEMOS EL PRODUCTO, PERO NOS DICE QUE NO SE PUEDE
	//Se muestra información de la solicitud de venta y un mensaje "No se puede. No hay suficiente cantidad".
	//Además muestra nuevamenete el inventario para ver los cambios
	lProductos.venderProducto("frijoles", 6)
	lProductos.imprimirInventario()

	//VENDEMOS EL PRODUCTO Y SE ELIMINA EL PRODUCTO DE LA LISTA PORQUE DICE QUE LA CANTIDAD ES CERO
	//Se muestra información de la solicitud de venta y un mensaje "vendido con éxito".
	//También, muestra un mensaje indicando que se elimino el producto porque no hay más existencias.
	//Además muestra nuevamenete el inventario para ver los cambios.
	lProductos.venderProducto("leche", 8)
	lProductos.imprimirInventario()

	//MUESTRA UNA LISTA CON PRODUCTOS DE EXISTENCIAS MINIMAS
	fmt.Println("\n\n--------------------------------------------------------------------------\n",
		"                         LISTA DE PRODUCTOS MÍNIMOS\n\n", lProductos.listarProductosMínimos(),
		"\n--------------------------------------------------------------------------------\n")

	//SE ORDENA LA LISTA DE PRODUCTOS
	//La lista de productos se ordena, se muestra un mensaje "se ordeno la lista".
	//Además, muestra el inventario para comprobar los cambios
	lProductos.ordenarLista()
	lProductos.imprimirInventario()

	//SE AUMENTA EL INVENTARIO PARA EXISTENCIAS MÍNIMAS
	//La lista de productos con existencias mínimas se le aumentan las unidades hasta llegar a la existencia mínima indicada.
	//Muestra un mensaje "se aumentó la lista de existencias mínimas"
	//Además, muestra el inventario para comprobar los cambios
	lProductos.aumentarInventarioDeMinimos()
	lProductos.imprimirInventario()
}

/*\
  RESPUESTA DEL CÓDIGO

GOROOT=C:\Program Files\Go #gosetup
GOPATH=C:\Users\Melany\Desktop\LENGUAJES DE PROGRAMACIÓN\TRABAJOS EN CLASE\S3-Venta_Productos\.idea #gosetup
"C:\Program Files\Go\bin\go.exe" build -o C:\Users\Melany\AppData\Local\Temp\GoLand\___go_build_main_go.exe "C:\Users\Melany\Desktop\LENGUAJES DE PROGRAMACIÓN\TRABAJOS EN CLASE\S3-Venta_Productos\main.go" #gosetup
C:\Users\Melany\AppData\Local\Temp\GoLand\___go_build_main_go.exe

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Nombre: arroz  - Precio: ₡ 2500  - Cantidad: 15
    Nombre: frijoles  - Precio: ₡ 2000  - Cantidad: 4
    Nombre: leche  - Precio: ₡ 1200  - Cantidad: 8
    Nombre: cafe  - Precio: ₡ 4500  - Cantidad: 12
    Nombre: sal  - Precio: ₡ 700  - Cantidad: 1
    Nombre: azucar  - Precio: ₡ 1800  - Cantidad: 7

--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Nombre: arroz  - Precio: ₡ 2500  - Cantidad: 15
    Nombre: frijoles  - Precio: ₡ 2000  - Cantidad: 4
    Nombre: leche  - Precio: ₡ 1200  - Cantidad: 8
    Nombre: cafe  - Precio: ₡ 4500  - Cantidad: 12
    Nombre: sal  - Precio: ₡ 700  - Cantidad: 1
    Nombre: azucar  - Precio: ₡ 1800  - Cantidad: 7
    Nombre: mantequilla  - Precio: ₡ 1000  - Cantidad: 20

--------------------------------------------------------------------------

--------------------------------------------------------------------------

             ********** SOLICITUD DE CAMBIO EN PRODUCTO **********

--- Información de cambio ---
   Nombre:  cafe
   Nuevo Precio:  1200
   Nueva Cantidad:  16

 ESTADO: Cambiado correctamente


--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Nombre: arroz  - Precio: ₡ 2500  - Cantidad: 15
    Nombre: frijoles  - Precio: ₡ 2000  - Cantidad: 4
    Nombre: leche  - Precio: ₡ 1200  - Cantidad: 8
    Nombre: cafe  - Precio: ₡ 1200  - Cantidad: 16
    Nombre: sal  - Precio: ₡ 700  - Cantidad: 1
    Nombre: azucar  - Precio: ₡ 1800  - Cantidad: 7
    Nombre: mantequilla  - Precio: ₡ 1000  - Cantidad: 20

--------------------------------------------------------------------------

--------------------------------------------------------------------------

            ********** SOLICITUD DE VENTA **********


--- Información de venta ---
   Nombre:  arroz
   Precio:  2500

 ESTADO: Se vendió con éxito.


--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Nombre: arroz  - Precio: ₡ 2500  - Cantidad: 11
    Nombre: frijoles  - Precio: ₡ 2000  - Cantidad: 4
    Nombre: leche  - Precio: ₡ 1200  - Cantidad: 8
    Nombre: cafe  - Precio: ₡ 1200  - Cantidad: 16
    Nombre: sal  - Precio: ₡ 700  - Cantidad: 1
    Nombre: azucar  - Precio: ₡ 1800  - Cantidad: 7
    Nombre: mantequilla  - Precio: ₡ 1000  - Cantidad: 20

--------------------------------------------------------------------------

--------------------------------------------------------------------------

            ********** SOLICITUD DE VENTA **********


--- Información de venta ---
   Nombre:  frijoles
   Precio:  2000

ESTADO: No se puede vender; No hay suficiente cantidad.


--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Nombre: arroz  - Precio: ₡ 2500  - Cantidad: 11
    Nombre: frijoles  - Precio: ₡ 2000  - Cantidad: 4
    Nombre: leche  - Precio: ₡ 1200  - Cantidad: 8
    Nombre: cafe  - Precio: ₡ 1200  - Cantidad: 16
    Nombre: sal  - Precio: ₡ 700  - Cantidad: 1
    Nombre: azucar  - Precio: ₡ 1800  - Cantidad: 7
    Nombre: mantequilla  - Precio: ₡ 1000  - Cantidad: 20

--------------------------------------------------------------------------

--------------------------------------------------------------------------

            ********** SOLICITUD DE VENTA **********


--- Información de venta ---
   Nombre:  leche
   Precio:  1200

 ESTADO: Se vendió con éxito.


--------------------------------------------------------------------------

--------------------------------------------------------------------------

            ********** INFORMACIÓN DEL PRODUCTO **********

El producto  leche  del precio  1200  ha sido eliminado;
No hay más existencias.

--------------------------------------------------------------------------

--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Nombre: arroz  - Precio: ₡ 2500  - Cantidad: 11
    Nombre: frijoles  - Precio: ₡ 2000  - Cantidad: 4
    Nombre: cafe  - Precio: ₡ 1200  - Cantidad: 16
    Nombre: sal  - Precio: ₡ 700  - Cantidad: 1
    Nombre: azucar  - Precio: ₡ 1800  - Cantidad: 7
    Nombre: mantequilla  - Precio: ₡ 1000  - Cantidad: 20

--------------------------------------------------------------------------


--------------------------------------------------------------------------
                          LISTA DE PRODUCTOS MÍNIMOS

 [{frijoles 4 2000} {sal 1 700} {azucar 7 1800}]
--------------------------------------------------------------------------------

--------------------------------------------------------------------------

                SE ORDENÓ LA LISTA DE PRODUCTOS


--------------------------------------------------------------------------


--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Nombre: sal  - Precio: ₡ 700  - Cantidad: 1
    Nombre: frijoles  - Precio: ₡ 2000  - Cantidad: 4
    Nombre: azucar  - Precio: ₡ 1800  - Cantidad: 7
    Nombre: arroz  - Precio: ₡ 2500  - Cantidad: 11
    Nombre: cafe  - Precio: ₡ 1200  - Cantidad: 16
    Nombre: mantequilla  - Precio: ₡ 1000  - Cantidad: 20

--------------------------------------------------------------------------

--------------------------------------------------------------------------

                  SE AUMENTÓ EL  INVENTARIO DE MÍNIMAS EXISTECNIAS

--------------------------------------------------------------------------


--------------------------------------------------------------------------

                      LISTA DE PRODUCTOS

    Nombre: sal  - Precio: ₡ 700  - Cantidad: 10
    Nombre: frijoles  - Precio: ₡ 2000  - Cantidad: 10
    Nombre: azucar  - Precio: ₡ 1800  - Cantidad: 10
    Nombre: arroz  - Precio: ₡ 2500  - Cantidad: 11
    Nombre: cafe  - Precio: ₡ 1200  - Cantidad: 16
    Nombre: mantequilla  - Precio: ₡ 1000  - Cantidad: 20

--------------------------------------------------------------------------

Process finished with the exit code 0




*/
