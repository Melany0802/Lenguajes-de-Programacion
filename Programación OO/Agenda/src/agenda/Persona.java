
package agenda;


public class Persona {
    private String nombre;
    private String fechaNacimiento;
    private String genero; 

    public Persona(String nombre, String fechaNacimiento, String genero) {
        this.nombre = nombre;
        this.fechaNacimiento = fechaNacimiento;
        this.genero = genero;
    }

    @Override
    public String toString() {
        return "Persona{" +
                "nombre='" + getNombre() + '\'' +
                ", fecha de nacimiento=" + getFechaNacimiento() +
                ", genero=" + getGenero() +
                '}';
    }

    public String getNombre() {
        return nombre;
    }

 
    public String getFechaNacimiento() {
        return fechaNacimiento;
    }

    public String getGenero() {
        return genero;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }


    public void setFechaNacimiento(String fechaNacimiento) {
        this.fechaNacimiento = fechaNacimiento;
    }

    public void setGenero(String genero) {
        this.genero = genero;
    }

    
}
