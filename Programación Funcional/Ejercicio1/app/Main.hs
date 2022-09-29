--                ___________________________________________________________________________________
--               |                                   EJERCICIO #1                                    |
--               |     Haciendo uso de la función filter, implemente una función que a partir        |
--               |     de una lista de cadenas de parámetro, filtre aquellas que contengan una       |
--               |     subcadena que el usuario indique en otro argumento. Ej                        |
--               |                                                                                   |
--               |      sub_cadenas “la” [“la casa, “el perro”, “pintando la cerca”]                 |
--               |                                                                                   |
--               |      [“la casa, “pintando la cerca”]                                              |
--               |                                                                                   |
--               |                                                                                   |
--               |___________________________________________________________________________________|



module Main (main) where

import Data.List
import Lib
import Data.List

sub_cadenas :: String -> [String] -> [String]
sub_cadenas str list = 
    filter (isInfixOf str) list


main :: IO ()
main = do
    
    print("")
    print(sub_cadenas "la" ["la casa", "el perro", "pintando la cerca"])
    print("")

    
  
{-       RESPUESTA DEL CÓDIGO
Windows PowerShell
Copyright (C) Microsoft Corporation. All rights reserved.

Install the latest PowerShell for new features and improvements! https://aka.ms/PSWindows

PS C:\Users\Melany\Desktop\Haskell\Repositorio\Ejercicio1> stack run
""
["la casa","pintando la cerca"]
""-}  



