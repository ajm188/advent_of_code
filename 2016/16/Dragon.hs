module Dragon
( dragon
) where

dragon :: String -> String
dragon a = a ++ ('0' : b)
    where b = (reverse . (map swap)) a
          swap '0' = '1'
          swap '1' = '0'
