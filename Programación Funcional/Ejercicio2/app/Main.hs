--                ___________________________________________________________________________________
--               |                                   EJERCICIO #2                                    |
--               |     Construya una función que se llame sub_conjunto. Esta función recibe          |
--               |     dos listas y debe retornar True cuando el primer argumento es subconjunto     |
--               |     completo del segundo y #f en caso contrario. Por ejemplo:                     |
--               |                                                                                   |
--               |      sub_conjunto [] [1,2,3,4,5]                                                  |
--               |        => True                                                                    |
--               |                                                                                   |
--               |      sub_conjunto [1,2,3] [1,2,3,4,5]                                             |
--               |        => True                                                                    |
--               |                                                                                   |
--               |      sub_conjunto [1,2,6] [1,2,3,4,5]                                             |
--               |        => False                                                                   |
--               |___________________________________________________________________________________|


module Main (main) where

import Lib

sub_conjunto :: (Eq a) => [a] -> [a] -> Bool

sub_conjunto (m:ms) (w:ws)

    | m == w = sub_conjunto ms ws
    | otherwise = sub_conjunto (m:ms) ws

sub_conjunto [] [] = True
sub_conjunto [] _ = True
sub_conjunto _ [] = False


main :: IO ()
main = do
    
    print("")
    print(sub_conjunto [][1,2,3,4,5])
    print(sub_conjunto [1,2,3][1,2,3,4,5])
    print(sub_conjunto [1,2,6][1,2,3,4,5])
    print("")



{-      RESPUESTA DEL CÓDIGO

Windows PowerShell
Copyright (C) Microsoft Corporation. All rights reserved.

Install the latest PowerShell for new features and improvements! https://aka.ms/PSWindows

PS C:\Users\Melany\Desktop\Haskell\Repositorio\Ejercicio2> stack run  
""
True
True
False
""-}