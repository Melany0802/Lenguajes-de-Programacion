%   -------------- EJERCICIO #4 --------------

% Modifique el predicado para averiguar las rutas en un grafo para
% que calcule la ruta más corta de todas en función de pesos que deben
% agregarse al grafo.


%
% ----- NUESTRO GRAFO -----
%           3       1
%        a ---- c ---- o
%   4  /   \   /
%     /	  5 \ / 2
%   i        X
%     \   7 / \ 9
%   10 \   /   \
%        b ---- d ---- f
%           6       15
%

conectar(a,i, 4).
conectar(a,c, 3).
conectar(a,x, 5).

conectar(i,b, 10).

conectar(b,d, 6).
conectar(b,x, 7).

conectar(d,f, 15).
conectar(d,x, 9).

conectar(c,o, 1).
conectar(c,x, 2).


ruta(Inicio,Fin,Ruta,Peso):- ruta2(Inicio,Fin,[],Ruta,Peso).

ruta2(Fin,Fin,_,[Fin],0).
ruta2(Inicio,Fin,RutaVisitada,[Inicio|Sobrante],Peso):-

    conectar(Inicio,Alguien,Peso2),
    not(member(Alguien,RutaVisitada)),
    ruta2(Alguien,Fin,[Inicio|RutaVisitada],Sobrante,P),
    Peso is P + Peso2.


listaRutasConPesos(Inicio,Fin,ListaRuta):-
    findall([Ruta,Peso],ruta(Inicio,Fin,Ruta,Peso),ListaRuta).

rutaCortaPesos(Inicio,Fin):-
    listaRutasConPesos(Inicio,Fin,ListaRuta),
    rutaCortaDeLosPesos(ListaRuta,[],100,RutaCorta,Peso),
    write('Ruta corta: '), write(RutaCorta), nl,
    write('Peso: '), write(Peso).

rutaCortaDeLosPesos([],RF,PF,RF,PF):-!.
rutaCortaDeLosPesos([[R,P]|T],_,Pa,RutaCorta,Peso):-
    P=<Pa,
    rutaCortaDeLosPesos(T,R,P,RutaCorta,Peso).

rutaCortaDeLosPesos([_|T],Ra,Pa,RutaCorta,Peso):-
    rutaCortaDeLosPesos(T,Ra,Pa,RutaCorta,Peso).
