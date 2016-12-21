shortestPath :: Int -> (Int, Int) -> (Int, Int) -> [(Int, Int)]
shortestPath num loc target = head (until done update [[loc]])
    where done = (==target) . head . head
          update (p:ps) = insertAll ps (map (\loc -> loc : p) (nextLocs p))
          nextLocs p = ((filter (openSpace num)) . (filter legalMove) . allNextLocs . head) p
          allNextLocs (x, y) = [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]
          insertAll ps [] = ps
          insertAll ps (x:xs) = insertAll (insert ps x) xs
          insert [] x = [x]
          insert (p:ps) x
            | (head x) == (head p) = p : ps
            | length p <= length x = p : (insert ps x)
            | otherwise = x : p : ps

legalMove :: (Int, Int) -> Bool
legalMove (x, y) = all (>=0) [x, y]

openSpace :: Int -> (Int, Int) -> Bool
openSpace num (x, y) = (even . length . (filter (=='1')) . toBinary) (t1 + num)
    where t1 = (x * x) + (3 * x) + (2 * x * y) + y + (y * y)

toBinary :: Int -> String
toBinary 0 = "0"
toBinary 1 = "1"
toBinary n = (toBinary n') ++ (show r')
    where (n', r') = divMod n 2

main = do
    let (favoriteNumber, target) = (1358, (31, 39))

    print $
        (\path -> (length path) - 1) $
        shortestPath favoriteNumber (1, 1) target
