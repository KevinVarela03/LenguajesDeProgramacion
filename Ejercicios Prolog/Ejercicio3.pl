aplanar([],[]).

aplanar([Cabeza|Cola], Resultado) :-
   aplanar(Cabeza, Sublista),
   aplanar(Cola, RestoResultado),
   append(Sublista, RestoResultado, Resultado).

aplanar([Cabeza|Cola], [Cabeza|Resultado]) :-
   aplanar(Cola, Resultado).

%Ejemplo
%aplanar([[1,2],[3,4],5,10],L).
