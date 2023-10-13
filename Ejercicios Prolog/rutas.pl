conectado(i,a).
conectado(i,b).
conectado(a,b).
conectado(a,c).
conectado(b,f).
conectado(b,c).
conectado(c,f).

conectados(X,Y):- conectado(X,Y).
conectados(X,Y):- conectado(Y,X).


ruta(Ini,Fin,Ruta):- ruta2(Ini,Fin,[],Ruta).

ruta2(Fin,Fin,_,[Fin]).
ruta2(Ini,Fin,Visitados,[Ini|Resto]):-
    conectados(Ini,Alguien),
    %%conectado(Ini,Alguien),
    not(member(Alguien,Visitados)),
    ruta2(Alguien,Fin,[Ini|Visitados],Resto).

encuentra_ruta_corta(Origen, Destino, RutaCorta) :-
    findall(Ruta, ruta(Origen, Destino, Ruta), Rutas),
    minima(Rutas, RutaCorta).

% encuentra la ruta más corta de una lista
minima([Ruta], Ruta).
minima([Ruta1,Ruta2|Cola], RutaCorta) :-
    length(Ruta1, Largo1),
    length(Ruta2, Largo2),
    (   Largo1 < Largo2 ->
        minima([Ruta1|Cola], RutaCorta)
    ;   Largo1 >= Largo2 ->
        minima([Ruta2|Cola], RutaCorta)
    ).




