
package agenda;

public class ContactoT1 extends Contacto{
    
    private String correo;

    public ContactoT1(Persona persona, String direccion, String telefono, String correo) {
        super(persona, direccion, telefono);
        this.correo = correo;
    }
    public void setCorreo(String correo) {
        this.correo = correo;
    }

    public ContactoT1(String nombre, String fechaNacimiento, String genero, String direccion, String telefono, String correo) {
        super(nombre, fechaNacimiento, genero, direccion, telefono);
        this.correo = correo;
    }


    @Override
    public void imprimir(){
        System.out.println("CONTACTO TIPO 1 : " + this.toString());
        ContactoT1Frame pantalla = new ContactoT1Frame();
        pantalla.nombre_txt.setText(this.getPersona().getNombre());
        pantalla.nacimiento_txt.setText(this.getPersona().getFechaNacimiento());
        pantalla.genero_txt.setText(this.getPersona().getGenero());
        pantalla.direccion_txt.setText(this.getDireccion());
        pantalla.telefono_txt.setText(this.getTelefono());
        pantalla.correo_txt.setText(correo);
        pantalla.setVisible(true);
    }

    @Override
    public String toString() {
        return "ContactoT1{" +
                super.toString() + '\'' +
                "correo='" + correo + '\'' +
                '}';
    }
    
   

    

    @Override
    public void modificarContacto(String fechaNacimiento, String genero, String telefono, String direccion, String correo) {
        this.getPersona().setFechaNacimiento(fechaNacimiento);
        this.getPersona().setGenero(genero);
        this.setTelefono(telefono);
        this.setDireccion(direccion);
        this.setCorreo(correo);
    }

    
}

