/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package editorgrafico;


public class Texto implements Objeto_Representativo{
    private String texto;

    public Texto(String texto) {
        this.texto = texto;
    }

    public String getTexto() {
        return texto;
    }

    @Override
    public void dibujar() {
        System.out.println("Texto escrito: " + getTexto());
    }

    @Override
    public String toString() {
        return "Texto{" +
                "texto='" + texto + '\'' +
                '}';
    }
}