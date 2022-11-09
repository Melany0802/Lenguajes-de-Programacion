/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package editorgrafico;

public class Cuadrado implements Objeto_Representativo{
    private int lado;

    public Cuadrado(int lado) {
        super();
        this.lado = lado;
    }

    public int getLado() {
        return lado;
    }

    @Override
    public void calcularArea() {
        setArea(getLado()*getLado());
    }

    @Override
    public void dibujar() {
        System.out.println("Lado del cuadrado: " + getLado());
        System.out.println("Area del cuadrado: " + getArea());
    }

    @Override
    public String toString() {
        return "Cuadrado{" +
                "lado = " + lado +
                ", area = " + getArea() +
                '}';
    }
}