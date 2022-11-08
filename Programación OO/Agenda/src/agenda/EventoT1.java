
package agenda;


public class EventoT1 extends Eventos{
    
    private String nivelImportante;
    
    public EventoT1(String asunto, String fecha, String jornada, String lugar,String nivelImportante){
        super(asunto,fecha, jornada, lugar);
        this.nivelImportante = nivelImportante;
        
    }
    public String getNivelImportante() {
        return nivelImportante;
    }

    public void setNivelImportante(String nivelImportante) {
        this.nivelImportante = nivelImportante;
    }
    

    @Override
    public void imprimir(){
        System.out.println("Evento TIPO 1 : " + this.toString());
        EventoT1Frame pantalla = new EventoT1Frame();
        pantalla.asunto_txt.setText(this.getAsunto());
        pantalla.fecha_txt.setText(this.getFecha());
        pantalla.jornada_txt2.setText(this.getJornada());
        pantalla.lugar_txt.setText(this.getLugar());
        pantalla.importancia_txt.setText(this.getNivelImportante());
        
        pantalla.setVisible(true);
    }

    @Override
    public String toString() {
        return "EventoT1{" +
                super.toString() + '\'' +
                "asunto='" + this.getAsunto() + '\'' +
                "fecha='" + this.getFecha() + '\'' +
                "jornada='" + this.getJornada() + '\'' +
                "lugar='" + this.getLugar() + '\'' +
                "nivelImportancia='" + this.getNivelImportante() + '\'' +
                '}';
    }

    @Override
    public void modificarEvento(String asunto, String fecha, String jornada, String lugar, String nivelImportante) {
        this.setAsunto(asunto);
        this.setFecha(fecha);
        this.setJornada(jornada);
        this.setLugar(lugar);
        this.setNivelImportante(nivelImportante);
    }

    
}
    

