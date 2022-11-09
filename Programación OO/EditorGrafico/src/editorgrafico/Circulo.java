
package editorgrafico;

public class Circulo implements Objeto_Representativo {
    private int radio;

    public Circulo(int radio) {
        super();
        this.radio = radio;
    }

    public int getRadio() {
        return radio;
    }

   
    @Override
    public void calcularArea() {
        double area = Math.PI * Math.pow(getRadio(), 2);
        setArea((int)area);
    }

    @Override
    public void dibujar() {
        System.out.println("Radio del círculo: " + getRadio());
        System.out.println("Area del circulo: " + getArea());

    }

    @Override
    public String toString() {
        return "Circulo{" +
                "radio = " + radio +
                ", area = " + getArea() +
                '}';
    }

    
}