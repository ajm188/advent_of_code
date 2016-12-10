data Direction = U | D | L | R

charToDirection :: Char -> Direction
charToDirection 'U' = U
charToDirection 'D' = D
charToDirection 'L' = L
charToDirection 'R' = R

move :: Int -> Direction -> Int
move 1 d =
    case d of
        R -> 2
        D -> 4
        _ -> 1
move 2 d =
    case d of
        L -> 1
        R -> 3
        D -> 5
        _ -> 2
move 3 d =
    case d of
        L -> 2
        D -> 6
        _ -> 3
move 4 d =
    case d of
        U -> 1
        R -> 5
        D -> 7
        _ -> 4
move 5 d =
    case d of
        U -> 2
        L -> 4
        R -> 6
        D -> 8
move 6 d =
    case d of
        U -> 3
        L -> 5
        D -> 9
        _ -> 6
move 7 d =
    case d of
        U -> 4
        R -> 8
        _ -> 7
move 8 d =
    case d of
        U -> 5
        L -> 7
        R -> 9
        _ -> 8
move 9 d =
    case d of
        U -> 6
        L -> 8
        _ -> 9

nextButton :: Int -> String -> Int
nextButton button directions = foldl move button $ map charToDirection directions

main = do
    input <- getContents
    print $ concat $ (map show) $ tail $
        foldl (\a@(x:_) line -> a ++ [(nextButton x line)]) [5] (lines input)
