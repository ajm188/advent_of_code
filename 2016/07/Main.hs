hasTLS x = hasTLS' x 0 False

hasTLS' :: String -> Int -> Bool -> Bool
hasTLS' (a:b:b':[]) bracketDepth foundAbba = foundAbba
hasTLS' (a:b:b':a':xs) bracketDepth foundAbba
    | a == '[' = hasTLS' rest (bracketDepth + 1) foundAbba
    | a == ']' = hasTLS' rest (bracketDepth - 1) foundAbba
    | a == a' && b == b' && a /= b = if bracketDepth > 0 then False else hasTLS' rest bracketDepth True
    | otherwise = hasTLS' rest bracketDepth foundAbba
    where rest = (b:b':a':xs)

main = do
    input <- getContents
    print $ length $ filter hasTLS $ lines input
