/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package editorgrafico;

public abstract class Objeto_Geometrico implements Objeto_Representativo{
    private int area;

    public Objeto_Geometrico(int area){
        this.area = area;
    }

    public Objeto_Geometrico() {
    }

    public int getArea() {
        return area;
    }

    public void setArea(int area) {
        this.area = area;
    }

    public abstract void calcularArea();


    @Override
    public abstract void dibujar();
}