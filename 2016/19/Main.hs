next :: [a] -> [a]
next [] = []
next (x:[]) = [x]
next xs
    | (even . length) xs = ((map fst) . (filter (\(a, b) -> even b))) (zip xs [0..])
    | otherwise = (last xs) : (next ((reverse . tail . reverse) xs))

winner elves = until (done) (next) elves
    where done = (==1) . length

main = do
    let num = 3014603
    print $
        head $
        winner [1..num]
