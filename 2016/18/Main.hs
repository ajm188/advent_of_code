isTrap :: (Bool, Bool, Bool) -> Bool
isTrap (left, center, right)
    | left && center && not right = True
    | center && right && not left = True
    | left && (not center && not right) = True
    | right && (not center && not left) = True
    | otherwise = False

toTrap :: Char -> Bool
toTrap '.' = False
toTrap '^' = True

nextRow :: [Bool] -> [Bool]
nextRow row = map isTrap (zip3 (False : (take ((length row) - 1) row)) row (tail (row ++ [False])))

room :: Int -> [[Bool]] -> [[Bool]]
room x rows
    | length rows == x = rows
    | otherwise = room x (((nextRow . head) rows) : rows)

main = do
    input <- getContents
    let startRow = map toTrap ((head . lines) input)
    let numRows1 = ((read . head . tail . lines) input) :: Int
    let numRows2 = ((read . head . tail . tail . lines) input) :: Int
    print $
        map (\x -> (sum . (map (length . (filter not))) . (room x)) [startRow]) [numRows1, numRows2]
