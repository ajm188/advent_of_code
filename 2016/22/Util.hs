module Util
( isDigit
, words'
) where

words' :: Eq a => [a] -> a -> [[a]]
words' [] splitter = []
words' xs@(x:_) splitter
    | x == splitter = words' (tail xs) splitter
    | otherwise = (takeWhile (/=splitter) xs) : (words' (dropWhile (/=splitter) xs) splitter)

isDigit :: Char -> Bool
isDigit x = elem x ['0'..'9']
