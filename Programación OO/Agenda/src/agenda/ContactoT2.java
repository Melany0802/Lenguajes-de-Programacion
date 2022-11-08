
package agenda;


public class ContactoT2 extends Contacto{
    private String facebook;
    private String instagram;
    private String twitter;

    public ContactoT2(Persona persona, String direccion, String telefono, String facebook, String instagram, String twitter) {
        super(persona, direccion, telefono);
        this.facebook = facebook;
        this.instagram = instagram;
        this.twitter = twitter;
    }

    public ContactoT2(String nombre, String fechaNacimiento, String genero, String direccion, String telefono, String facebook, String instagram, String twitter) {
        super(nombre, fechaNacimiento, genero, direccion, telefono);
        this.facebook = facebook;
        this.instagram = instagram;
        this.twitter = twitter;
    }
    
     public void setFacebook(String facebook) {
        this.facebook = facebook;
    }

    
    public void setInstagram(String instagram) {
        this.instagram = instagram;
    }

   
    public void setTwitter(String twitter) {
        this.twitter = twitter;
    }
    

    @Override
    public void imprimir(){
        System.out.println("CONTACTO TIPO 2 : " + this.toString());
        ContactoT2Frame pantalla = new ContactoT2Frame();
        pantalla.nombre_txt.setText(this.getPersona().getNombre());
        pantalla.nacimiento_txt.setText(this.getPersona().getFechaNacimiento());
        pantalla.genero_txt.setText(this.getPersona().getGenero());
        pantalla.direccion_txt.setText(this.getDireccion());
        pantalla.telefono_txt.setText(this.getTelefono());
        pantalla.face_txt.setText(this.facebook);
        pantalla.insta_txt.setText(instagram);
        pantalla.twi_txt.setText(twitter);
        
        pantalla.setVisible(true);
    }

    @Override
    public String toString() {
        return "ContactoT2{" +
                "facebook='" + facebook + '\'' +
                ", instagram='" + instagram + '\'' +
                ", twitter='" + twitter + '\'' +
                "} " + super.toString();
    }

    @Override
    public void modificarContacto(String fechaNacimiento, String genero, String telefono, String direccion, String correo) {
        this.getPersona().setFechaNacimiento(fechaNacimiento);
        this.getPersona().setGenero(genero);
        this.setTelefono(telefono);
        this.setDireccion(direccion);
        this.setFacebook(facebook);
        this.setInstagram(instagram);
        this.setTwitter(twitter);
      
    }

   

   
}

   
   
