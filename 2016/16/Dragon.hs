module Dragon
( dragon
) where

dragon :: String -> String
dragon a = a ++ ('0' : b)
    where b = map (\x -> if x == '0' then '1' else '0') $ reverse a
