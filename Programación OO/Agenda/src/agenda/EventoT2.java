
package agenda;

public class EventoT2 extends Eventos{
      private String telefonoOfinina;
      private String nombrePaginaWeb;
      
    
    public EventoT2(String asunto, String fecha, String jornada, String lugar,String telefonoOfinina, String nombrePaginaWeb){
        super(asunto,fecha, jornada, lugar);
        this.telefonoOfinina = telefonoOfinina;
        this.nombrePaginaWeb = nombrePaginaWeb;
        
    }
    public void setTelefonoOfinina(String telefonoOfinina) {
        this.telefonoOfinina = telefonoOfinina;
    }

    public void setNombrePaginaWeb(String nombrePaginaWeb) {
        this.nombrePaginaWeb = nombrePaginaWeb;
    }

    @Override
    public void imprimir(){
        System.out.println("Evento TIPO 2 : " + this.toString());
        EventoT2Frame pantalla = new EventoT2Frame();
        pantalla.asunto_txt.setText(this.getAsunto());
        pantalla.fecha_txt.setText(this.getFecha());
        pantalla.jornada_txt2.setText(this.getJornada());
        pantalla.lugar_txt.setText(this.getLugar());
        pantalla.pagina_txt.setText(this.nombrePaginaWeb);
        pantalla.oficina_txt.setText(telefonoOfinina);
        
       
        
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
                "nombrePaginaWeb='" + this.nombrePaginaWeb + '\'' +
                "telefonoOfinina='" + this.telefonoOfinina + '\'' +
                '}';
    }

    @Override
    public void modificarEvento(String asunto, String fecha, String jornada, String lugar, String nivelImportante) {
       this.setAsunto(asunto);
       this.setFecha(fecha);
       this.setJornada(jornada);
       this.setLugar(lugar);
       this.setNombrePaginaWeb(nombrePaginaWeb);
       this.setTelefonoOfinina(telefonoOfinina);
    }
               
}

    
