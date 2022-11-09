
package editorgrafico;

public class Main{
    public static void main (String[] args) {
        System.out.println("-----------Círculo-----------");
        Circulo circulo = new Circulo(6);
        circulo.calcularArea();
        circulo.dibujar();

        System.out.println("-----------Cuadrado-----------");
        Cuadrado cuadrado = new Cuadrado(10);
        cuadrado.calcularArea();
        cuadrado.dibujar();

        System.out.println("-----------Texto-----------");
        Texto texto = new Texto("Hola mundo soy Julieth");
        texto.dibujar();

        System.out.println("-------------------Grupo-------------------");
        Grupo grupo = new Grupo();
        grupo.addObjetos(circulo);
        grupo.addObjetos(cuadrado);
        grupo.dibujar();

        System.out.println("-------------------Hoja-------------------");
        Hoja hoja = new Hoja("21cm x 29cm");
        hoja.addObjetos(grupo);
        System.out.println(hoja);

        System.out.println("-------------------Documento-------------------");
        Documento doc = new Documento("Documento de prueba");
        doc.addHojas(hoja);
        System.out.println(doc);
    }
}