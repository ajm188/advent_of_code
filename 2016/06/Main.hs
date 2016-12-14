import Data.List

decode :: [String] -> String
decode columns = foldl (++) "" $ map (\x -> [mostFrequent x]) columns

decode2 :: [String] -> String
decode2 columns = foldl (++) "" $ map (\x -> [leastFrequent x]) columns

mostFrequent :: Eq a => [a] -> a
mostFrequent items = fst $ maxBy (\(_, x) (_, y) -> compare x y) $ countDupes items

leastFrequent :: Eq a => [a] -> a
leastFrequent items = fst $ minBy (\(_, x) (_, y) -> compare x y) $ countDupes items

countDupes :: Eq a => [a] -> [(a, Int)]
countDupes items = foldl counter [] items
    where counter [] item = [(item, 1)]
          counter (t@(i, c):xs) item = case item == i of True -> (i, c + 1) : xs
                                                         _    -> t : (counter xs item)

maxBy :: (a -> a -> Ordering) -> [a] -> a
maxBy f (x:xs) = foldl orderer x xs
    where orderer x y = case f x y of LT -> y
                                      _  -> x

minBy :: (a -> a -> Ordering) -> [a] -> a
minBy f (x:xs) = foldl orderer x xs
    where orderer x y = case f x y of GT -> y
                                      _  -> x

main = do
    input <- getContents
    print $
        decode $ transpose $ lines input
    print $
        decode2 $ transpose $ lines input
