module Node
( Node(..)
, newNode
, viablePair
) where

import Util

data Node = Node (Int, Int) Int Int
          deriving (Show)

newNode :: String -> Int -> Int -> Node
newNode name used avail = Node (read (dropWhile (not . isDigit) x), read (dropWhile (not . isDigit) y)) used avail
    where (x:y:_) = ((drop 1) . (words' name)) '-'

viablePair :: Node -> Node -> Bool
viablePair (Node locA usedA availA) (Node locB usedB availB)
    | usedA == 0 = False
    | locA == locB = False
    | otherwise = usedA < availB
