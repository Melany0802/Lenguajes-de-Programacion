%   -------------- EJERCICIO #3 --------------

% Implemente la función aplanar. Esta función recibe una lista con
% múltiples listas anidadas como elementos y devuelve una lista con los
% mismos elementos de manera lineal (sin listas).


aplanar([], []) :- !.
aplanar([L|Ls], FlatL) :-
    !,
    aplanar(L, NewL),
    aplanar(Ls, NewLs),
    append(NewL, NewLs, FlatL).
aplanar(L, [L]).
