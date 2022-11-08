
package agenda;
        import java.util.LinkedList;

public class Agenda {
    private LinkedList<Contacto> contactos;
    private LinkedList<Eventos> eventos;

    public Agenda() {
        this.contactos = new LinkedList<Contacto>();
        this.eventos = new LinkedList<Eventos>();
    }

    //Agregar Contacto
    public void agregarContacto(Contacto c){
        this.contactos.add(c);
    }
    //Eliminar Contacto
    public void eliminarContacto(int index){
        this.contactos.remove(index);
    }
    public void eliminarContacto(String nombre){
        for(Contacto c : this.contactos){
            if (c.getPersona().getNombre().equals(nombre))
                this.contactos.remove(c);
        }
    }
    //Modificar Contacto
    public void modificarContacto(String nombre, String fechaNacimiento, String genero, String telefono, String direccion, String correo){
        for(Contacto c : this.contactos){
            if (c.getPersona().getNombre().equals(nombre)){
                c.modificarContacto(fechaNacimiento,genero,telefono, direccion, correo);
                break;
            }
        }

    }
    
    //Agregar Evento
    public void agregarEvento(Eventos e){
        this.eventos.add(e);
    }
    //Eliminar Evento
    public void eliminarEventos(int index){
        this.contactos.remove(index);
    }
    public void eliminarEventos(String asunto){
        for(Eventos e : this.eventos){
            if (e.getAsunto().equals(asunto))
                this.eventos.remove(e);
        }
    }
    //Modificar Evento
    public void modificarEvento(String asunto, String fecha, String jornada, String lugar, String nivelImportante){
        for(Eventos e : this.eventos){
            if (e.getAsunto().equals(asunto)){
                e.modificarEvento(asunto, fecha, jornada, lugar, nivelImportante);
                break;
            }
        }

    }

    //Imprimir Contactos
    public void imprimirContactos(){
        for(Contacto c : this.contactos)
            c.imprimir();
    }
    
    //Imprimir Eventos
    public void imprimirEventos(){
        for(Eventos e : this.eventos)
            e.imprimir();
    }
}
