--                ___________________________________________________________________________________
--               |                                   EJERCICIO #4                                    |
--               |       Implemente el equivalente para la función anterior, pero utilizando la      |
--               |       función map para dicho fin.                                                 |
--               |                                                                                   |
--               |___________________________________________________________________________________|


module Main (main) where

import Data.List

aplanarMap :: [[a]] -> [a]
aplanarMap = concatMap id

main :: IO ()
main = do

  print("")
  print(aplanarMap [[1, 2], [3, 4, 5, 6, 7], [8,9,10]])
  print("")

{-        RESPUESTA DEL CÓDIGO

Windows PowerShell
Copyright (C) Microsoft Corporation. All rights reserved.

Install the latest PowerShell for new features and improvements! https://aka.ms/PSWindows

PS C:\Users\Melany\Desktop\Haskell\Repositorio\Ejercicio4> stack run
""
[1,2,3,4,5,6,7,8,9,10]
""
}  
