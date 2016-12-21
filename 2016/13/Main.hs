locationsWithin :: Int -> (Int, Int) -> Int -> [(Int, Int)]
locationsWithin num loc steps = snd (until done update ([[loc]], []))
    where done = (null . fst)
          update ((p:ps), seen) = (insertAll ps (map (\loc -> loc : p) (nextLocs num p)) seen, (head p) : seen)
          insertAll ps [] _ = ps
          insertAll ps (x:xs) seen = insertAll (insert ps x seen) xs seen
          insert [] x seen
            | (length x) - 1 > steps = []
            | elem (head x) seen = []
            | otherwise = [x]
          insert (p:ps) x seen
            | elem (head x) seen  = p : ps
            | (head x) == (head p) = p : ps
            | (length x) - 1 > steps = p : ps
            | length p <= length x = p : (insert ps x seen)
            | otherwise = x : p : ps

shortestPath :: Int -> (Int, Int) -> (Int, Int) -> [(Int, Int)]
shortestPath num loc target = head (until done update [[loc]])
    where done = (==target) . head . head
          update (p:ps) = insertAll ps (map (\loc -> loc : p) (nextLocs num p))
          insertAll ps [] = ps
          insertAll ps (x:xs) = insertAll (insert ps x) xs
          insert [] x = [x]
          insert (p:ps) x
            | (head x) == (head p) = p : ps
            | length p <= length x = p : (insert ps x)
            | otherwise = x : p : ps

nextLocs :: Int -> [(Int, Int)] -> [(Int, Int)]
nextLocs num path = ((filter (openSpace num)). (filter legalMove) . allNextLocs . head) path

allNextLocs :: (Int, Int) -> [(Int, Int)]
allNextLocs (x, y) = [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]

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
    print $
        length $
        locationsWithin favoriteNumber (1, 1) 50
