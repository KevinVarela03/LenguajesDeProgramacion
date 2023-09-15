open System

let funcionShift (lista : int list) numDesp direccion =

    let nuevaLista =
        let largo = List.length lista

        if numDesp >= largo then
            printfn "Numero de desplazamientos: %A" numDesp
            printfn "La lista vieja es: %A" lista
            List.replicate largo 0 
        else
            let n = numDesp % largo

            match direccion with
            | "izq" ->
                printfn "Numero de desplazamientos hacia la izquierda: %A" numDesp
                printfn "La lista vieja es: %A" lista
                List.append (List.skip n lista) (List.replicate n 0)
            | "der" ->
                printfn "Numero de desplazamientos hacia la derecha: %A" numDesp
                printfn "La lista vieja es: %A" lista
                List.append (List.replicate n 0) (List.take (largo - n) lista)
            | _ ->
                lista

    printfn "La nueva lista es: %A\n" nuevaLista

[<EntryPoint>]
let main argv =
    let mutable lista : int list = []
    lista <- [1; 2; 3; 4; 5]

    funcionShift lista 3 "izq"
    funcionShift lista 2 "der"
    funcionShift lista 6 "izq" 
    0
