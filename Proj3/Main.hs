{-
 - Author: Cael Shoop, cshoop2018@my.fit.edu
 - Course: CSE4250, Fall 2020
 - Project: Sat Solver
 - Language implementation: Glorious Glasgow Haskell Compilation System, version 8.4.3
 -}

main :: IO()

main = do
    _ <- getLine
    putStrLn "P"
    _ <- getLine
    putStrLn "((P & (Q v R)) => (P & Q))"
    _ <- getLine
    putStrLn "((P & (Q v R)) <=> ((P & Q) v (P & R)))"