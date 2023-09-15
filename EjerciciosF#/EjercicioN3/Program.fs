open System

let n_esimo n lista =
    let seleccionarElemento indice elemento =
        if indice = n then Some elemento
        else None

    let elementosSeleccionados = List.mapi seleccionarElemento lista

    match List.tryFind (fun x -> Option.isSome x) elementosSeleccionados with
    | Some (Some (valor)) -> valor
    | _ -> 0 // Valor predeterminado si no se encuentra ningún elemento

[<EntryPoint>]
let main argv =
    let lista = [1; 2; 3; 4; 5]

    let resultado1 = n_esimo 2 lista
    let resultado2 = n_esimo 3 lista
    let resultado3 = n_esimo 6 lista

    printfn "n_esimo 2 [1,2,3,4,5] = %A" resultado1
    printfn "n_esimo 3 [1,2,3,4,5] = %A" resultado2
    printfn "n_esimo 6 [1,2,3,4,5] = %A" resultado3

    0
