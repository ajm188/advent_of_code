next :: [a] -> [a]
next [] = []
next (x:[]) = [x]
next xs
    | (even . length) xs = ((map fst) . (filter (\(a, b) -> even b))) (zip xs [0..])
    | otherwise = (last xs) : (next ((reverse . tail . reverse) xs))

next2 :: [a] -> [a]
next2 [] = []
next2 (x:[]) = [x]
next2 xs = (drop numRemoved front) ++ back' ++ (take numRemoved front)
    where half = (div (length xs) 2)
          (front, back) = splitAt half xs
          back'
            | length back < 3 = tail back
            | otherwise = ((map fst) . (filter (\(_, b) -> (mod b 3) == 0))) (zip back [indexStart..])
          indexStart = if (even . length) xs then 1 else 2
          numRemoved = (length back) - (length back')

winner elves = until (done) (next) elves
    where done = (==1) . length

winner2 elves = until done next2 elves
    where done = (==1) . length

main = do
    let num = 3014603
    print $
        head $
        winner [1..num]
    print $
        head $
        winner2 [1..num]
