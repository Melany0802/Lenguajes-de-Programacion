
package agenda;


public class Main {

   
    public static void main(String[] args) {
        
        Agenda agendaPersonal = new Agenda();
        System.out.println("---------- AGREGANDO CONTACTOS ----------");
        
        agendaPersonal.agregarContacto(new ContactoT1("Mariana","07 Enero, 2000","Femenino","Cartago","86354896","manuela@suCorreo.com" ));
        agendaPersonal.agregarContacto(new ContactoT1("Samuel","02 Noviembre, 2017","Masculino","Cedral","84217639","samuel@suCorreo.com" ));
        
        agendaPersonal.agregarContacto(new ContactoT2("Santiago","25 Febrero, 1995", "Prefiero no decirlo", "Fortuna", "84269514", "Santiago Gonzales", "santi25", "sanGon25"));
        agendaPersonal.agregarContacto(new ContactoT2("Valencia","10 Septiembre, 2014", "Femenino", "San José", "84026315", "Valencia Torres", "val10", "vanTor10"));
        
        
        //agendaPersonal.eliminarContacto("Valencia");
        //agendaPersonal.modificarContacto("Mariana", "10 Enero, 2001", "Femenino", "85743659", "Cartago", "manuela@suCorreo.com");
        
        System.out.println("***IMPRIMIENDO CONTACTOS***");
        agendaPersonal.imprimirContactos();
        
        
        
//---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------        
//---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------        
       System.out.println("\n\n\n---------- AGREGANDO EVENTOS ----------\n");
        
        agendaPersonal.agregarEvento(new EventoT1("Reunión Contable", "25 de Noviembre, 2022", "Diurna", "Ofinica Contable", "Importante"));
        agendaPersonal.agregarEvento(new EventoT1("Reunión Ejecutiva", "13 de Diciembre, 2022", "Nocturna", "Ofinica Secretaria", "Muy Importante"));
        
        agendaPersonal.agregarEvento(new EventoT2("Compra Muebles","05 Diciembre, 2022", "Diurna", "San José", "20149685", "InOfis"));
        agendaPersonal.agregarEvento(new EventoT2("Ajustes de Pago","28 Noviembre, 2022", "Diurna", "Alajuela", "23697450", "EjecutivoFull"));
        
        //agendaPersonal.eliminarEventos("Compra Muebles");
        //agendaPersonal.modificarEvento("Reunión Ejecutiva", "15 de Noviembre, 2022", "Nocturna", "Cartago", "Muy Importante");
        
        System.out.println("***IMPRIMIENDO EVENTOS***");
        agendaPersonal.imprimirEventos();
    }
   
    
 /*RESPUESTA DEL CÓDIGO
   ==> IMPRIMIENDO CONTACTOS
    ---------- AGREGANDO CONTACTOS ----------
***IMPRIMIENDO CONTACTOS***
CONTACTO TIPO 1 : ContactoT1{Contacto{persona=Persona{nombre='Mariana', fecha de nacimiento=07 Enero, 2000, genero=Femenino}, direccion='Cartago', telefono='86354896'}'correo='manuela@suCorreo.com'}
CONTACTO TIPO 1 : ContactoT1{Contacto{persona=Persona{nombre='Samuel', fecha de nacimiento=02 Noviembre, 2017, genero=Masculino}, direccion='Cedral', telefono='84217639'}'correo='samuel@suCorreo.com'}
CONTACTO TIPO 2 : ContactoT2{facebook='Santiago Gonzales', instagram='santi25', twitter='sanGon25'} Contacto{persona=Persona{nombre='Santiago', fecha de nacimiento=25 Febrero, 1995, genero=Prefiero no decirlo}, direccion='Fortuna', telefono='84269514'}
CONTACTO TIPO 2 : ContactoT2{facebook='Valencia Torres', instagram='val10', twitter='vanTor10'} Contacto{persona=Persona{nombre='Valencia', fecha de nacimiento=10 Septiembre, 2014, genero=Femenino}, direccion='San José', telefono='84026315'}
BUILD SUCCESSFUL (total time: 2 seconds)

  
   ==> ELIMINANDO CONTACTOS
    
---------- AGREGANDO CONTACTOS ----------
***IMPRIMIENDO CONTACTOS***
CONTACTO TIPO 1 : ContactoT1{Contacto{persona=Persona{nombre='Mariana', fecha de nacimiento=07 Enero, 2000, genero=Femenino}, direccion='Cartago', telefono='86354896'}'correo='manuela@suCorreo.com'}
CONTACTO TIPO 1 : ContactoT1{Contacto{persona=Persona{nombre='Samuel', fecha de nacimiento=02 Noviembre, 2017, genero=Masculino}, direccion='Cedral', telefono='84217639'}'correo='samuel@suCorreo.com'}
CONTACTO TIPO 2 : ContactoT2{facebook='Santiago Gonzales', instagram='santi25', twitter='sanGon25'} Contacto{persona=Persona{nombre='Santiago', fecha de nacimiento=25 Febrero, 1995, genero=Prefiero no decirlo}, direccion='Fortuna', telefono='84269514'}
BUILD SUCCESSFUL (total time: 2 seconds)
    
    
  
   ==> MODIFICANDO CONTACTOS
    
    ---------- AGREGANDO CONTACTOS ----------
***IMPRIMIENDO CONTACTOS***
CONTACTO TIPO 1 : ContactoT1{Contacto{persona=Persona{nombre='Mariana', fecha de nacimiento=10 Enero, 2001, genero=Femenino}, direccion='Cartago', telefono='85743659'}'correo='manuela@suCorreo.com'}
CONTACTO TIPO 1 : ContactoT1{Contacto{persona=Persona{nombre='Samuel', fecha de nacimiento=02 Noviembre, 2017, genero=Masculino}, direccion='Cedral', telefono='84217639'}'correo='samuel@suCorreo.com'}
CONTACTO TIPO 2 : ContactoT2{facebook='Santiago Gonzales', instagram='santi25', twitter='sanGon25'} Contacto{persona=Persona{nombre='Santiago', fecha de nacimiento=25 Febrero, 1995, genero=Prefiero no decirlo}, direccion='Fortuna', telefono='84269514'}
CONTACTO TIPO 2 : ContactoT2{facebook='Valencia Torres', instagram='val10', twitter='vanTor10'} Contacto{persona=Persona{nombre='Valencia', fecha de nacimiento=10 Septiembre, 2014, genero=Femenino}, direccion='San José', telefono='84026315'}
BUILD SUCCESSFUL (total time: 4 seconds)


    */
    
    

/*RESPUESTA DEL CÓDIGO
  ==>IMPRIMIENDO EVENTOS
    ---------- AGREGANDO EVENTOS ----------

***IMPRIMIENDO EVENTOS***
Evento TIPO 1 : EventoT1{Evento{fecha=25 de Noviembre, 2022, jornada ='Diurna', lugar='Ofinica Contable', asunto='Reunión Contable'}'asunto='Reunión Contable'fecha='25 de Noviembre, 2022'jornada='Diurna'lugar='Ofinica Contable'nivelImportancia='Importante'}
Evento TIPO 1 : EventoT1{Evento{fecha=13 de Diciembre, 2022, jornada ='Nocturna', lugar='Ofinica Secretaria', asunto='Reunión Ejecutiva'}'asunto='Reunión Ejecutiva'fecha='13 de Diciembre, 2022'jornada='Nocturna'lugar='Ofinica Secretaria'nivelImportancia='Muy Importante'}
Evento TIPO 2 : EventoT1{Evento{fecha=05 Diciembre, 2022, jornada ='Diurna', lugar='San José', asunto='Compra Muebles'}'asunto='Compra Muebles'fecha='05 Diciembre, 2022'jornada='Diurna'lugar='San José'nombrePaginaWeb='InOfis'telefonoOfinina='20149685'}
Evento TIPO 2 : EventoT1{Evento{fecha=28 Noviembre, 2022, jornada ='Diurna', lugar='Alajuela', asunto='Ajustes de Pago'}'asunto='Ajustes de Pago'fecha='28 Noviembre, 2022'jornada='Diurna'lugar='Alajuela'nombrePaginaWeb='EjecutivoFull'telefonoOfinina='23697450'}
BUILD SUCCESSFUL (total time: 2 seconds)

    
    
  ==>ELIMINANDO EVENTOS
    ---------- AGREGANDO EVENTOS ----------

***IMPRIMIENDO EVENTOS***
Evento TIPO 1 : EventoT1{Evento{fecha=25 de Noviembre, 2022, jornada ='Diurna', lugar='Ofinica Contable', asunto='Reunión Contable'}'asunto='Reunión Contable'fecha='25 de Noviembre, 2022'jornada='Diurna'lugar='Ofinica Contable'nivelImportancia='Importante'}
Evento TIPO 1 : EventoT1{Evento{fecha=13 de Diciembre, 2022, jornada ='Nocturna', lugar='Ofinica Secretaria', asunto='Reunión Ejecutiva'}'asunto='Reunión Ejecutiva'fecha='13 de Diciembre, 2022'jornada='Nocturna'lugar='Ofinica Secretaria'nivelImportancia='Muy Importante'}
Evento TIPO 2 : EventoT1{Evento{fecha=28 Noviembre, 2022, jornada ='Diurna', lugar='Alajuela', asunto='Ajustes de Pago'}'asunto='Ajustes de Pago'fecha='28 Noviembre, 2022'jornada='Diurna'lugar='Alajuela'nombrePaginaWeb='EjecutivoFull'telefonoOfinina='23697450'}
BUILD SUCCESSFUL (total time: 2 seconds)

    
    
   ==>MODIFICANDO EVENTOS 
    ---------- AGREGANDO EVENTOS ----------

***IMPRIMIENDO EVENTOS***
Evento TIPO 1 : EventoT1{Evento{fecha=25 de Noviembre, 2022, jornada ='Diurna', lugar='Ofinica Contable', asunto='Reunión Contable'}'asunto='Reunión Contable'fecha='25 de Noviembre, 2022'jornada='Diurna'lugar='Ofinica Contable'nivelImportancia='Importante'}
Evento TIPO 1 : EventoT1{Evento{fecha=15 de Noviembre, 2022, jornada ='Nocturna', lugar='Cartago', asunto='Reunión Ejecutiva'}'asunto='Reunión Ejecutiva'fecha='15 de Noviembre, 2022'jornada='Nocturna'lugar='Cartago'nivelImportancia='Muy Importante'}
Evento TIPO 2 : EventoT1{Evento{fecha=05 Diciembre, 2022, jornada ='Diurna', lugar='San José', asunto='Compra Muebles'}'asunto='Compra Muebles'fecha='05 Diciembre, 2022'jornada='Diurna'lugar='San José'nombrePaginaWeb='InOfis'telefonoOfinina='20149685'}
Evento TIPO 2 : EventoT1{Evento{fecha=28 Noviembre, 2022, jornada ='Diurna', lugar='Alajuela', asunto='Ajustes de Pago'}'asunto='Ajustes de Pago'fecha='28 Noviembre, 2022'jornada='Diurna'lugar='Alajuela'nombrePaginaWeb='EjecutivoFull'telefonoOfinina='23697450'}
BUILD SUCCESSFUL (total time: 9 seconds)

    
    
    */    
    
    
}
