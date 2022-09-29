--                ___________________________________________________________________________________
--               |                                   EJERCICIO #3                                    |
--               |     Implemente la función aplanar. Esta función recibe una lista con múltiples    |
--               |     listas anidadas como elementos y devuelve una lista con los mismos elementos  |
--               |     de manera lineal (sin listas). Ej:                                            |
--               |                                                                                   |
--               |      aplanar ‘(1 2 (3 (4 5) (6 7))))                                              |
--               |                                                                                   |
--               |      (1 2 3 4 5 6 7)                                                              |
--               |                                                                                   |
--               |                                                                                   |
--               |___________________________________________________________________________________|


module Main (main) where

aplanar :: [[a]] -> [a]
aplanar [] = []
aplanar xs = foldr (++) [] xs


main :: IO ()
main = do

  print("")
  print(aplanar [[1, 2], [3, 4, 5, 6]])
  print("")



{-          RESPUESTA DEL CÓDIGO  

Windows PowerShell
Copyright (C) Microsoft Corporation. All rights reserved.

Install the latest PowerShell for new features and improvements! https://aka.ms/PSWindows

PS C:\Users\Melany\Desktop\Haskell\Repositorio\Ejercicio3> stack run
""
[1,2,3,4,5,6]
""-}


