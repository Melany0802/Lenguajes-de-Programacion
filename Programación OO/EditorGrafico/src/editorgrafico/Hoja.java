/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package editorgrafico;
import java.util.LinkedList;

public class Hoja {
    private String dimensiones;
    private LinkedList<Objeto_Representativo> objetos;

    public Hoja(String dimensiones) {
        this.dimensiones = dimensiones;
        this.objetos = new LinkedList<Objeto_Representativo>();
    }

    public String getDimensiones() {
        return dimensiones;
    }

    public void setDimensiones(String dimensiones) {
        this.dimensiones = dimensiones;
    }

    public void addObjetos(Objeto_Representativo objeto) {
        this.objetos.add(objeto);
    }

    @Override
    public String toString() {
        return "Hoja{" +
                "dimensiones='" + dimensiones + '\'' +
                ", objetos=" + objetos +
                '}';
    }

}