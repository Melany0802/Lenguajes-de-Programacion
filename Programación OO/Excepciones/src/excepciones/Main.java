
package excepciones;

import java.util.List;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        
        Scanner sc = new Scanner(System.in);
        System.out.println("Ingrese la operación: ");
        String operacion = sc.nextLine();

        Evaluador e = new Evaluador();
        List expresion = e.enlistar(operacion);
        System.out.println("El resultado total es: "+e.resultado(expresion).getNum());
    }
}


/*RESPUESTA DEL CÓDIGO

=>CONSULTA #1
Ingrese la operación: 
512 * 30 / 2
El resultado total es: 7680

=>CONSULTA #2
Ingrese la operación: 
512 * 30 / 2a
Error, la operación contiene un carácter inválido
Error. Los elementos no estan ordenados en infijo
El resultado total es: -1

=>CONSULTA #3
Ingrese la operación: 
512 * 30 / - 2 
Error. Los elementos no estan ordenados en infijo
El resultado total es: -1


*/
