eliminarElemento(_, [], []).

eliminarElemento(Y, [Y|Xs], Zs):-
          eliminarElemento(Y, Xs, Zs), !.

eliminarElemento(X, [Y|Xs], [Y|Zs]):-
          eliminarElemento(X, Xs, Zs).
