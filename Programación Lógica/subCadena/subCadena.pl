%   -------------- EJERCICIO #1 --------------

% Implemente un predicado que, a partir de una lista de cadenas
% de parámetro, filtre aquellas que contengan una subcadena que el
% usuario indique en otro argumento.



subCadena(FiltrarPalabra,Palabra) :-
        sub_string(Palabra, _, _, _, FiltrarPalabra),!.
filtrarSubCadena(FiltrarPalabra, X, Y) :-
        include(subCadena(FiltrarPalabra), X, Y).
