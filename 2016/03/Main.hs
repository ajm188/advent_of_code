isTriangle :: (Int, Int, Int) -> Bool
isTriangle (a, b, c) = a + b > c && b + c > a && a + c > b

toInts :: String -> [Int]
toInts line = map (\x -> read x :: Int) $ words line

asTriplet :: [Int] -> (Int, Int, Int)
asTriplet (a:b:c:_) = (a, b, c)

main = do
    input <- getContents
    print $ length $ filter isTriangle $
        (map (asTriplet . toInts) $ lines input)
