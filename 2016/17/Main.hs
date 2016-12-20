import qualified Data.Hash.MD5 as MD5

data State = State String (Int, Int) String

shortestPath :: String -> String
shortestPath passcode = (path . head) (until done update [start])
    where done = (isGoal . head)
          update xs = update' (tail xs) ((next . head) xs)
          update' [] ys = ys
          update' xs [] = xs
          update' (x:xs) (y:ys)
            | (length . path) x < (length . path) y = x : (update' xs (y:ys))
            | otherwise = y : (update' (x:xs) ys)
          start = State passcode (0, 3) ""

path :: State -> String
path (State _ _ p) = p

isGoal :: State -> Bool
isGoal (State _ location _) = location == (3, 0)

next :: State -> [State]
next (State passcode loc path) = map (\m -> State passcode (move loc m) (path ++ m)) (moves passcode loc path)

moves :: String -> (Int, Int) -> String -> [String]
moves passcode location path = (filter (legal . (move location))) (doors passcode path)

doors passcode path = ((map snd) . (filter doorIsOpen)) (zip hash ["U", "D", "L", "R"])
    where hash = (MD5.md5s . MD5.Str) (passcode ++ path)
          doorIsOpen (h, _) = elem h "abcdef"

move :: (Int, Int) -> String -> (Int, Int)
move (x, y) "U" = (x, y + 1)
move (x, y) "D" = (x, y - 1)
move (x, y) "L" = (x - 1, y)
move (x, y) "R" = (x + 1, y)

legal :: (Int, Int) -> Bool
legal (x, y) = x >= 0 && x < 4 && y >= 0 && y < 4

main = do
    print $ shortestPath "edjrjqaa"
