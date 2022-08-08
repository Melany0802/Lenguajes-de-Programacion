package main

import (
	"fmt"
)

type persona struct {
	nombre    string
	apellido  string
	cromosoma []byte
}

var lista_personas_array [5]persona
var lista_personas_slice []persona

func crearPersona(n string, a string, c []byte) persona {
	p := persona{nombre: n, apellido: a, cromosoma: c}
	return p
}

func agregarPersona(p persona) {
	lista_personas_slice = append(lista_personas_slice, p) //El slice sí me permite agregar
}

func llenar1() { //Se crean las personas, nos permite meter los datos del arreglo
	lista_personas_array[0].nombre = "Oscar1"
	lista_personas_array[0].apellido = "Viquez1"
	lista_personas_array[0].cromosoma = []byte{'1', '0', '1', '0', '0', '1', '1', '1', '0', '1'}

	lista_personas_array[1].nombre = "Oscar2"
	lista_personas_array[1].apellido = "Viquez2"
	lista_personas_array[1].cromosoma = []byte{'1', '1', '1', '0', '1', '0', '1', '0', '0', '1'}

	lista_personas_array[2].nombre = "Oscar3"
	lista_personas_array[2].apellido = "Viquez3"
	lista_personas_array[2].cromosoma = []byte{'0', '1', '1', '0', '0', '1', '0', '1', '1', '0'}

	lista_personas_array[3].nombre = "Oscar4"
	lista_personas_array[3].apellido = "Viquez4"
	lista_personas_array[3].cromosoma = []byte{'0', '0', '1', '1', '0', '0', '0', '1', '0', '0'}

	lista_personas_array[4].nombre = "Oscar5"
	lista_personas_array[4].apellido = "Viquez5"
	lista_personas_array[4].cromosoma = []byte{'1', '1', '0', '1', '1', '1', '0', '0', '0', '1'}
}

func llenar2() { //Se crean las personas, nos permite meter los datos del slice
	//Solo se metio los datos de una persona
	agregarPersona(crearPersona("Oscar1", "Viquez1", []byte{'1', '0', '1', '0', '0', '1', '1', '1', '0', '1'}))
}

// PRIMERA IMPLEMENTACIÓN
// recorrido tradicional con range... pruebe hacerlo utilizando la función len(arreglo)
func masParecido(lista [5]persona, muestra [10]byte) persona { //El resultado más parecido lo que retorna es un número
	//  |=Se recibe la lista de 5 personas
	cont, max := 0, 0 //El "max" tiene la cantidad de genes iguales
	masparecido := 0
	for i, individuo := range lista { //Recorre persona por persona en el rango de la lista
		cont = 0
		for j, gen := range individuo.cromosoma { //Para cada persona en el rango del cromosoma
			if gen == muestra[j] {
				cont++
			}
		}
		if cont > max {
			masparecido = i
			max = cont
		}

	}
	return
}

// SEGUNDA IMPLEMENTACIÓN
// usando punteros a arreglos... un poco más difícil de leer
func masParecido2(lista *[5]persona, muestra [10]byte) persona {
	//    |=Se recibe el puntero de 5 personas
	cont, max := 0, 0
	masparecido := 0
	for i, individuo := range *lista {
		cont = 0
		for j, gen := range individuo.cromosoma {
			if gen == muestra[j] {
				cont++
			}
		}
		if cont > max {
			masparecido = i
			max = cont
		}

	}
	return lista[masparecido] //Saca al más parecido dentro de la lista
	//Ya que arriba, la lista tiene persona
}

// TERCERA IMPLEMENTACIÓN
// USANDO PUNTEROS A LOS REGISTROS DE CADA UNO DE LOS ELEMENTOS DEL ARREGLO
func (person *persona) verificar(muestra [10]byte) int { //Se verifica el resultado más parecido del individuo
	cont := 0
	for j, gen := range person.cromosoma {
		if gen == muestra[j] {
			cont++
		}
	}
	return cont
}

func masParecido3(lista [5]persona, muestra [10]byte) persona { //Tanto la "muestra" como "persona" son arreglos
	cont, max := 0, 0
	masparecido := 0
	for i, individuo := range lista {
		cont = individuo.verificar(muestra) //Envia la "muestra" pero se hace un enlace al puntero
		if cont > max {
			masparecido = i
			max = cont
		}

	}
	return lista[masparecido] //Saca al más parecido dentro de la lista
	//Ya que arriba, la lista tiene persona
}

// CUARTA IMPLEMENTACIÓN
// Usando Slices... Es mejor en Go utilizar Slice en vez de punteros y hay más flexibilidad de trabajar con el arreglo
func masParecido4(lista []persona, muestra []byte) persona { //Tanto la "muestra" como persona son slice
	cont, max := 0, 0
	masparecido := 0
	for i, individuo := range lista {
		cont = 0
		for j, gen := range individuo.cromosoma {
			if gen == muestra[j] {
				cont++
			}
		}
		if cont > max {
			masparecido = i
			max = cont
		}

	}
	return lista[masparecido] //Saca al más parecido dentro de la lista
	//Ya que arriba, la lista tiene persona
}

func main() {
	llenar1()
	llenar2()

	fmt.Println("\n")
	fmt.Printf("Cromosoma Persona #2 ARRAY: %s\n", string(lista_personas_array[1].cromosoma))

	fmt.Println("Más parecido: ", masParecido(lista_personas_array, [10]byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))
	fmt.Println("Más parecido: ", masParecido2(&lista_personas_array, [10]byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))
	fmt.Println("Más parecido: ", masParecido3(lista_personas_array, [10]byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))

	fmt.Println("\n\n")
	fmt.Printf("Cromosoma Persona #1 SLICE: %s\n", string(lista_personas_array[1].cromosoma))
	fmt.Println("Más parecido: ", masParecido4(lista_personas_slice, []byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))
	agregarPersona(crearPersona("Oscar2", "Viquez2", []byte{'1', '1', '0', '1', '0', '0', '1', '1', '0', '1'}))
	fmt.Println("Más parecido: ", masParecido4(lista_personas_slice, []byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))
}
