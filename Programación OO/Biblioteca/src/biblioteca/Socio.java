
package biblioteca;

import java.util.LinkedList;

public class Socio {
 private int id;
    private String nombre;
    private String direccion;
    private LinkedList<Libro> libros;

    public Socio(int id, String nombre, String direccion) {
        this.id = id;
        this.nombre = nombre;
        this.direccion = direccion;
        this.libros = new LinkedList<Libro>();
    }

    public int getId() {
        return id;
    }

    public String getNombre() {
        return nombre;
    }

    public String getDireccion() {
        return direccion;
    }

    public LinkedList<Libro> getLibros() {
        return libros;
    }

    public int cantidadDeLibros(){
        if(libros.size() == 3){
            return 1;
        }
        return 0;
    }
    public void librosPrestados(Libro libro){
        libros.add(libro);
    }

    @Override
    public String toString() {
        final StringBuilder sb = new StringBuilder("Socio{");
        sb.append("Id del Socio= ").append(id);
        sb.append(", Nombre= '").append(nombre).append('\'');
        sb.append(", Direccion= '").append(direccion).append('\'');
        sb.append(", Libros= ").append(libros);
        sb.append('}');
        return sb.toString();
    }
}