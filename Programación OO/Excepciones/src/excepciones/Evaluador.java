
package excepciones;


import java.util.ArrayList;
import java.util.List;
import java.util.StringTokenizer;

public class Evaluador {
     public Evaluador() {
    }

    public List enlistar(String operacion){
        List lista = new ArrayList<>();
        StringTokenizer st = new StringTokenizer(operacion,"+/-* ", true);
        try {
            while (st.hasMoreTokens()) {
                String tk = st.nextToken();
                if (tk.matches("[0-9]*")){
                    Numero n = new Numero(Integer.valueOf(tk));
                    lista.add(n);
                }
                else if(tk.matches("[+/*[-]]+")){
                    Operador o = new Operador(tk);
                    lista.add(o);
                }
                else if(tk.equals(" ")){
                    continue;
                }
                else {
                    throw new Excepcion("Error, la operación contiene un carácter inválido");
                }
            }
       } catch (Excepcion e) {
            System.err.println(e.getMessage());
        }
        return lista;
    }

    public int operar(int n1, int n2, String operador){
        int resultado = 0;
        switch (operador){
            case "+":
                resultado = n1 + n2;
                return resultado;
            case "-":
                resultado = n1 - n2;
                return resultado;
            case "*":
                resultado = n1 * n2;
                return resultado;
            case "/":
                resultado = n1 / n2;
                return resultado;
        }
        return resultado;
    }

    public Numero resultado(List lista){
        Numero result = new Numero(-1);
        Object n1, n2, operador;
        Numero num1;
        Numero num2;
        Operador op;

        try {
            for(int i = 0; i < lista.size();){
                if (lista.size() < 3){
                    throw new Excepcion("Los elementos de la expresión no están ordenados en infijo");
                }
                n1 = lista.get(i);
                operador = lista.get(i+1);
                n2 = lista.get(i+2);
                if((n1.getClass().toString().equals("class Num")) && (operador.getClass().toString().equals("class Op")) && (n2.getClass().toString().equals("class Num"))){
                    num1 = (Numero) n1;
                    num2 = (Numero) n2;
                    op = (Operador) operador;

                    result.setNum(operar(num1.getNum(),num2.getNum(),op.getOp()));
                    lista.add(0, result);
                    lista.subList(1, 4).clear();

                }else {
                    throw new Excepcion("Error. Los elementos no estan ordenados en infijo");
                }
                if (lista.size() == 1) {
                    return result;
                }
            }

        }catch (Excepcion e) {
            System.err.println(e.getMessage());
        }

        return result;
    }
    
}
