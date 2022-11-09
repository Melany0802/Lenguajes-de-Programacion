
package biblioteca;

import org.w3c.dom.ls.LSOutput;

public class Main {

    public static void main(String[] args) {
        // Socios
        Socio socio1 = new Socio(1111, "Camila", "Fortuna");
        Socio socio2 = new Socio(2222, "Samuel", "Ciudad Quesada");
        Socio socio3 = new Socio(3333, "Daniel", "Cedral");
        Socio socio4 = new Socio(4444, "Cristina", "Aguas Zarcas");

        // Libros
        Libro libro1 = new Libro(2020, "A traves de mi ventana", "Ariana Godoy", true);
        Libro libro2 = new Libro(3030, "After: Aquí empieza todo", "Anna Todd", true);
        Libro libro3 = new Libro(4040, "Cazadores de sombras", "Cassandra Clare", true);
        Libro libro4 = new Libro(5050, "Heartstopper", "Alice Oseman", true);
        Libro libro5 = new Libro(6060, "Bajo la misma estrella", "John Green", true);

        // Biblioteca
        Biblioteca biblioteca = new Biblioteca("FantasiaCris");
        biblioteca.agregarLibros(libro1);
        biblioteca.agregarLibros(libro2);
        biblioteca.agregarLibros(libro3);
        biblioteca.agregarLibros(libro4);
        biblioteca.agregarLibros(libro5);

        biblioteca.agregarSocios(socio1);
        biblioteca.agregarSocios(socio2);
        biblioteca.agregarSocios(socio3);
        biblioteca.agregarSocios(socio4);

        biblioteca.agregarPrestamo(new Prestamo(1001, 1111, 4040, "15/01/2017"));
        biblioteca.agregarPrestamo(new Prestamo(1001, 3333, 2020, "07/08/2019"));
        biblioteca.agregarPrestamo(new Prestamo(1001, 4444, 3030, "08/03/2021"));
        biblioteca.agregarPrestamo(new Prestamo(1001, 2222, 6060, "07/05/2002"));

        System.out.println(biblioteca.cantidadDeSociosLibros());
        System.out.println(biblioteca.toString());

    }
}

/*RESPUESTA DEL CODIGO

Socio{Id del Socio= 1111, Nombre= 'Camila', Direccion= 'Fortuna', Libros= [Libro {Código del libro = 4040, Titulo = Cazadores de sombras', Autor = Cassandra Clare', Disponibilidad = false}]}
Socio{Id del Socio= 2222, Nombre= 'Samuel', Direccion= 'Ciudad Quesada', Libros= [Libro {Código del libro = 6060, Titulo = Bajo la misma estrella', Autor = John Green', Disponibilidad = false}]}
Socio{Id del Socio= 3333, Nombre= 'Daniel', Direccion= 'Cedral', Libros= [Libro {Código del libro = 2020, Titulo = A traves de mi ventana', Autor = Ariana Godoy', Disponibilidad = false}]}
Socio{Id del Socio= 4444, Nombre= 'Cristina', Direccion= 'Aguas Zarcas', Libros= [Libro {Código del libro = 3030, Titulo = After: Aquí empieza todo', Autor = Anna Todd', Disponibilidad = false}]}

Biblioteca{Nombre de la Biblioteca= 'FantasiaCris', Nombre del Socios= [Socio{Id del Socio= 1111, Nombre= 'Camila', Direccion= 'Fortuna', Libros= [Libro {Código del libro = 4040, Titulo = Cazadores de sombras', Autor = Cassandra Clare', Disponibilidad = false}]}, Socio{Id del Socio= 2222, Nombre= 'Samuel', Direccion= 'Ciudad Quesada', Libros= [Libro {Código del libro = 6060, Titulo = Bajo la misma estrella', Autor = John Green', Disponibilidad = false}]}, Socio{Id del Socio= 3333, Nombre= 'Daniel', Direccion= 'Cedral', Libros= [Libro {Código del libro = 2020, Titulo = A traves de mi ventana', Autor = Ariana Godoy', Disponibilidad = false}]}, Socio{Id del Socio= 4444, Nombre= 'Cristina', Direccion= 'Aguas Zarcas', Libros= [Libro {Código del libro = 3030, Titulo = After: Aquí empieza todo', Autor = Anna Todd', Disponibilidad = false}]}], Nombre del Libros= [Libro {Código del libro = 2020, Titulo = A traves de mi ventana', Autor = Ariana Godoy', Disponibilidad = false}, Libro {Código del libro = 3030, Titulo = After: Aquí empieza todo', Autor = Anna Todd', Disponibilidad = false}, Libro {Código del libro = 4040, Titulo = Cazadores de sombras', Autor = Cassandra Clare', Disponibilidad = false}, Libro {Código del libro = 5050, Titulo = Heartstopper', Autor = Alice Oseman', Disponibilidad = true}, Libro {Código del libro = 6060, Titulo = Bajo la misma estrella', Autor = John Green', Disponibilidad = false}], Nombre del Prestamos= [Prestamo {Id del Prestamo = 1001, Id del Socio = 1111, Código del Libro = 4040, Fecha = 15/01/2017'}, Prestamo {Id del Prestamo = 1001, Id del Socio = 3333, Código del Libro = 2020, Fecha = 07/08/2019'}, Prestamo {Id del Prestamo = 1001, Id del Socio = 4444, Código del Libro = 3030, Fecha = 08/03/2021'}, Prestamo {Id del Prestamo = 1001, Id del Socio = 2222, Código del Libro = 6060, Fecha = 07/05/2002'}]}
BUILD SUCCESSFUL (total time: 0 seconds)

*/