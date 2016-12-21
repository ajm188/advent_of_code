module Operation
( Operation(..)
, operate
) where

import Data.Ix

data Operation = SwapIndices Int Int
               | SwapLetters Char Char
               | RotateLeft Int
               | RotateRight Int
               | RotateAroundLetter Char
               | Reverse Int Int
               | Move Int Int
               deriving (Show)

operate :: String -> Operation -> String
operate str (SwapIndices x y) = map (\(c', c) -> if c == x then y' else if c == y then x' else c') (zip str [0..])
    where (x', y') = (str !! x, str !! y)
operate str (SwapLetters x y) = map (\c -> if c == x then y else if c == y then x else c) str
operate str (RotateLeft steps)
    | steps == 0 = str
    | otherwise = operate str' (RotateLeft (steps - 1))
    where str' = (tail str) ++ [head str]
operate str (RotateRight steps)
    | steps == 0 = str
    | otherwise = operate str' (RotateRight (steps - 1))
    where str' = (last str) : (take (n - 1) str)
          n = length str
operate str (RotateAroundLetter x) = operate str rotator
    where xIndex = (snd . head) (filter (\(c, i) -> c == x) (zip str [0..]))
          rotator = RotateRight (1 + xIndex + maybeExtra)
          maybeExtra = if xIndex >= 4 then 1 else 0
operate str (Reverse x y) = front ++ (reverse middle) ++ back
    where front = ((map fst) . (filter (\(a, b) -> b < x))) (zip str [0..])
          middle = ((map fst) . (filter (\(a, b) -> inRange (x, y) b))) (zip str [0..])
          back = ((map fst) . (filter (\(a, b) -> b > y))) (zip str [0..])
operate str (Move x y) = str''
    where chr = str !! x
          str' = (take x str) ++ (drop (x + 1) str)
          str'' = (take y str') ++ [chr] ++ (drop y str')

instance Read Operation where

    readsPrec _ ('s':'w':'a':'p':rs)
        | r1 == "position" = [(SwapIndices ((read p1) :: Int) ((read p2) :: Int), "")]
        | otherwise = [(SwapLetters (head p1) (head p2), "")]
        where (r1:p1:_:_:p2:_) = words rs
    readsPrec _ ('r':'o':'t':'a':'t':'e':rs)
        | dir == "left" = [(RotateLeft n, "")]
        | dir == "right" = [(RotateRight n, "")]
        | otherwise = [(RotateAroundLetter ((head . last) rs'), "")]
        where rs'@(dir:_) = words rs
              n = read (rs' !! 1) :: Int
              c = (head . last) rs'
    readsPrec _ ('r':'e':'v':'e':'r':'s':'e':rs) = [(Reverse ((read p1) :: Int) ((read p2) :: Int), "")]
        where (_:p1:_:p2:_) = words rs
    readsPrec _ ('m':'o':'v':'e':rs) = [(Move ((read p1) :: Int) ((read p2) :: Int), "")]
        where (_:p1:_:_:p2:_) = words rs
