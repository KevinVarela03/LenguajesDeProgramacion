partir([], _, [], []).
partir([Cabeza|Cola], Umbral, Menores, Mayores) :-
    Cabeza < Umbral,
    partir(Cola, Umbral, NuevosMenores, Mayores),
    Menores = [Cabeza|NuevosMenores].

partir([Cabeza|Cola], Umbral, Menores, Mayores) :-
    Cabeza >= Umbral,
    partir(Cola, Umbral, Menores, NuevosMayores),
    Mayores = [Cabeza|NuevosMayores].

%Ejemplo
%partir([2,7,4,8,9,1,11,0], 6, Menores, Mayores).
