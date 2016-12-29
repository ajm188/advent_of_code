hasTLS x = hasTLS' x 0 False

hasTLS' :: String -> Int -> Bool -> Bool
hasTLS' (a:b:b':[]) bracketDepth foundAbba = foundAbba
hasTLS' (a:b:b':a':xs) bracketDepth foundAbba
    | a == '[' = hasTLS' rest (bracketDepth + 1) foundAbba
    | a == ']' = hasTLS' rest (bracketDepth - 1) foundAbba
    | a == a' && b == b' && a /= b = if bracketDepth > 0 then False else hasTLS' rest bracketDepth True
    | otherwise = hasTLS' rest bracketDepth foundAbba
    where rest = (b:b':a':xs)

hasSSL :: String -> Bool
hasSSL ip = any hasBAB hypernets
    where (supernets, hypernets) = decompose ip
          abas = foldl (++) [] (map findABAs supernets)
          hasBAB hypernet = any (\(a:b:_) -> containsSequence [b, a, b] hypernet) abas

decompose :: String -> ([String], [String])
decompose [] = ([], [])
decompose ('[':xs) = (supernets, inside : hypernets)
    where inside = takeWhile (/=']') xs
          rest = drop ((length inside) + 1) xs
          (supernets, hypernets) = decompose rest
decompose ip@(x:xs) = (supernet : supernets, hypernets)
    where supernet = takeWhile (/='[') ip
          rest = drop (length supernet) ip
          (supernets, hypernets) = decompose rest

findABAs :: Eq a => [a] -> [[a]]
findABAs xs
    | length xs < 3 = []
    | (head xs) == (head (drop 2 xs)) && (head xs) /= (head (tail xs)) = (take 3 xs) : (findABAs (tail xs))
    | otherwise = findABAs (tail xs)

containsSequence :: Eq a => [a] -> [a] -> Bool
containsSequence seq xs
    | length xs < length seq = False
    | (take (length seq) xs) == seq = True
    | otherwise = containsSequence seq (tail xs)

main = do
    input <- getContents
    print $ length $ filter hasTLS $ lines input
    print $ length $ filter hasSSL $ lines input
