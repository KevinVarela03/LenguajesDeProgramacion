open System

let filtrarCadenasConSubcadena subcadena listaCadenas =
    let contieneSubcadena cadena =
        (string cadena).Contains(subcadena)

    List.filter contieneSubcadena listaCadenas

[<EntryPoint>]
let main argv =
    let subcadena = "la"
    let listaCadenas = ["la casa"; "el perro"; "pintando la cerca"; "Kevin tiene una nueva camisa, la condenada es linda"]

    let resultado = filtrarCadenasConSubcadena subcadena listaCadenas

    printfn "Cadenas que contienen la subcadena \"%s\":" subcadena
    List.iter (fun cadena -> printfn "- %s" cadena) resultado

    0
