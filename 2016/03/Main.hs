import Data.List (transpose)

isTriangle :: (Int, Int, Int) -> Bool
isTriangle (a, b, c) = a + b > c && b + c > a && a + c > b

toInts :: String -> [Int]
toInts line = map (\x -> read x :: Int) $ words line

asTriplet :: [Int] -> (Int, Int, Int)
asTriplet (a:b:c:_) = (a, b, c)

groupBy :: Int -> [String] -> [[String]]
groupBy x [] = []
groupBy x items = (transpose $ map words $ take x $ items) ++ groupBy x (drop x items)

main = do
    input <- getContents
    print $ length $ filter isTriangle $
        (map (asTriplet . toInts) $ lines input)
    print $
        length $ filter isTriangle $
        map (asTriplet . toInts . unwords) $
        groupBy 3 $ lines input
