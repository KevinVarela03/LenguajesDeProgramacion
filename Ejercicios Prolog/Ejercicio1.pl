sumlist([],0).
sumlist([X | Resto], Suma) :-
    sumlist(Resto, SumaResto),
    Suma is X + SumaResto.

%Ejemplo
%sumlist([1,2,3,4],10).
