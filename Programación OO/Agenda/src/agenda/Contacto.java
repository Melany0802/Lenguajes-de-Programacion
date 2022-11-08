
package agenda;

public abstract class Contacto extends Object {
  
    private Persona persona;
    private String direccion;
    private String telefono;

    public Contacto(Persona persona, String direccion, String telefono) {
        this.persona = persona;
        this.direccion = direccion;
        this.telefono = telefono;
    }

    public Contacto(String nombre, String fechaNacimiento, String genero, String direccion, String telefono) {
        this.persona = new Persona(nombre,fechaNacimiento,genero);
        this.direccion = direccion;
        this.telefono = telefono;
    }

    public Persona getPersona() {
        return persona;
    }

    public abstract void imprimir();
    public abstract void modificarContacto(String fechaNacimiento, String genero, String telefono, String direccion, String correo);

    @Override
    public String toString() {
        return "Contacto{" +
                "persona=" + persona +
                ", direccion='" + getDireccion() + '\'' +
                ", telefono='" + getTelefono() + '\'' +
                '}';
    }

    public String getDireccion() {
        return direccion;
    }

    public String getTelefono() {
        return telefono;
    }

    public void setPersona(Persona persona) {
        this.persona = persona;
    }

    public void setDireccion(String direccion) {
        this.direccion = direccion;
    }

    public void setTelefono(String telefono) {
        this.telefono = telefono;
    }
    
    

}
