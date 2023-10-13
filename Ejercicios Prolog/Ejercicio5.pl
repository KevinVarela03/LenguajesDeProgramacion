sub_cadenas(_, [], []).

sub_cadenas(Subcadena, [Cadena|Resto], [Cadena|Filtradas]) :-
    subcadena_presente(Subcadena, Cadena),
    sub_cadenas(Subcadena, Resto, Filtradas).

sub_cadenas(Subcadena, [Cadena|Resto], Filtradas) :-
    \+ subcadena_presente(Subcadena, Cadena),
    sub_cadenas(Subcadena, Resto, Filtradas).

subcadena_presente(Subcadena, Cadena) :-
    sub_string(Cadena, _, _, _, Subcadena).

%Ejemplo
%sub_cadenas("la", ["la casa", "el perro", "pintando la cerca"], Filtradas).
