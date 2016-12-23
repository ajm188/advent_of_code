module Direction
( Direction(..)
, charToDirection
) where

data Direction = U | D | L | R

charToDirection :: Char -> Direction
charToDirection 'U' = U
charToDirection 'D' = D
charToDirection 'L' = L
charToDirection 'R' = R
