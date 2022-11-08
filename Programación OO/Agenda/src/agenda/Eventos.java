
package agenda;


public abstract class Eventos extends Object{
    
    private String asunto;
    private String fecha;
    private String jornada;
    private String lugar;
    
    
    
    public Eventos(String asunto, String fecha, String jornada, String lugar){
        this.asunto = asunto;
        this.fecha = fecha;
        this.jornada = jornada;
        this.lugar = lugar;
       
    }

    
    public String getFecha() {
        return fecha;
    }

    public String getJornada() {
        return jornada;
    }

    public String getLugar() {
        return lugar;
    }

    public String getAsunto() {
        return asunto;
    }
    
    public void setAsunto(String asunto) {
        this.asunto = asunto;
    }

  
    public void setFecha(String fecha) {
        this.fecha = fecha;
    }

    public void setJornada(String jornada) {
        this.jornada = jornada;
    }

    public void setLugar(String lugar) {
        this.lugar = lugar;
    }

    
    
    
    public abstract void imprimir();
    public abstract void modificarEvento(String asunto, String fecha, String jornada, String lugar, String nivelImportante);


    @Override
    public String toString() {
        return "Evento{" +
                "fecha=" + this.getFecha() +
                ", jornada ='" + this.getJornada() + '\'' +
                ", lugar='" + this.getLugar() + '\'' +
                ", asunto='" + this.getAsunto() + '\'' +
                '}';
    }


}
