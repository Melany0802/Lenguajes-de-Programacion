/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package editorgrafico;

import java.util.LinkedList;

public class Documento {
    private String nombre;
    private LinkedList<Hoja> hojas;

    public Documento(String nombre) {
        this.nombre = nombre;
        this.hojas = new LinkedList<Hoja>();
    }

    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public LinkedList<Hoja> getHojas() {
        return hojas;
    }

    public void  addHojas(Hoja hoja) {
        this.hojas.add(hoja);
    }

    @Override
    public String toString() {
        return "Documento{" +
                "nombre='" + nombre + '\'' +
                ", hojas=" + hojas +
                '}';
    }
}