%   -------------- EJERCICIO #3 --------------

% Implemente la funci�n aplanar. Esta funci�n recibe una lista con
% m�ltiples listas anidadas como elementos y devuelve una lista con los
% mismos elementos de manera lineal (sin listas).


aplanar([], []) :- !.
aplanar([L|Ls], FlatL) :-
    !,
    aplanar(L, NewL),
    aplanar(Ls, NewLs),
    append(NewL, NewLs, FlatL).
aplanar(L, [L]).
