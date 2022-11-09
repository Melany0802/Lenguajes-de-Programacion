/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package editorgrafico;

import java.util.LinkedList;

public class Grupo implements Objeto_Representativo {
    private LinkedList<Objeto_Representativo> objetos;

    public Grupo() {
        super();
        this.objetos = new LinkedList<Objeto_Representativo>();
    }

    public LinkedList<Objeto_Representativo> getObjetos() {
        return objetos;
    }

    public void addObjetos(Objeto_Representativo objeto) {
        this.objetos.add(objeto);
    }

    @Override
    public void dibujar() {
        System.out.println("Este grupo de objetos representables tiene:");
        System.out.println(getObjetos());
        
    }
    @Override
    public String toString() {
        return "Grupo{" +
                "objetos=" + objetos +
                '}';
    }
    
}