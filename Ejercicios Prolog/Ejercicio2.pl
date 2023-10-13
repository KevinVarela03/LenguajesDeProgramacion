subconj([], []).
subconj([H|T], [H|T1]) :- subconj(T, T1).
subconj([_|T], T1) :- subconj(T, T1).

%Ejemplo
%subconj([a, b, c], [a]).
%subconj([a, b, c], [a,c]).
