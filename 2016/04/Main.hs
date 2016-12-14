isDigit :: (Char -> Bool)
isDigit = (\x -> elem x ['0'..'9'])

splitUp :: [String] -> [String]
splitUp (x:xs) = (fst partition):(snd partition):xs
    where partition = break isDigit x

asTrip :: [String] -> (String, Int, String)
asTrip (a:b:c:_) = (a, read b :: Int, c)

countDupes :: Eq a => [a] -> [(a, Int)]
countDupes items = foldl counter [] items
    where counter [] item = [(item, 1)]
          counter (t@(i, c):xs) item = case item == i of True -> (i, c + 1) : xs
                                                         _    -> t : (counter xs item)
mostFrequent :: Ord a => [a] -> a
mostFrequent items = fst $ maxBy freqComparator $ countDupes items

freqComparator :: (Ord a, Ord b) => (a, b) -> (a, b) -> Ordering
freqComparator (x1, y1) (x2, y2) = case result of EQ -> compare x2 x1
                                                  _  -> result
    where result = compare y1 y2

maxBy :: (a -> a -> Ordering) -> [a] -> a
maxBy f (x:xs) = foldl orderer x xs
    where orderer x y = case f x y of LT -> y
                                      _  -> x

sortBy :: Ord a => (a -> a -> Ordering) -> [a] -> [a]
sortBy _ [] = []
sortBy f l = reverse $ fst $ foldl func ([], l) [1..(length l)]
    where func (sorted, rest) _ = ((maxBy f rest) : sorted, (filter (/= (maxBy f rest)) rest))

isRealRoom :: (String, a, String) -> Bool
isRealRoom (name, _, checksum) = checksum == top5Freq
    where top5Freq = take 5 $ map fst $ sortBy freqComparator $ countDupes name

main = do
    input <- getContents
    print $
        sum $
        map roomId $
        filter isRealRoom $
        map (asTrip . splitUp . words) $
        lines input
    where roomId (_, id', _) = id'
