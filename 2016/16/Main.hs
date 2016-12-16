import Dragon (dragon)

fillDisk :: String -> Int -> String
fillDisk state diskLen
    | length state >= diskLen = take diskLen state
    | otherwise = fillDisk (dragon state) diskLen

checksum :: String -> String
checksum state = checksum' $ map (\(a, b) -> if a == b then '1' else '0') $ pairs state

checksum' :: String -> String
checksum' state
    | mod (length state) 2 == 0 = checksum state
    | otherwise = state

pairs :: [a] -> [(a, a)]
pairs xs
    | length xs < 2 = []
    | otherwise = ((a, b) : pairs rest)
    where a = head xs
          b = (head . tail) xs
          rest = (tail . tail) xs

initialState = "10010000000110000"

main = do
    print $
        checksum $
        fillDisk initialState 272
