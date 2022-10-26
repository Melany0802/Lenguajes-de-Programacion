%   -------------- EJERCICIO #2 --------------

% Construya una funci�n que se llame sub_conjunto. Esta funci�n recibe
% dos listas y debe retornar True cuando el primer argumento es
% subconjunto completo del segundo y #f en caso contrario.


subConjunto([],[_]).
subConjunto([],[]).
subConjunto([E|RL], [E2|RL2]) :-
	 E = E2 ->
        subConjunto(RL,RL2);
	subConjunto([E|RL],RL2),
	subConjunto([],[RL2]).
