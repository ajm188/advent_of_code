type Point = (Int, Int)

data Direction = North | South | East | West

origin :: Point
origin = (0, 0)

displacement :: Point -> Point -> Int
displacement (x1, y1) (x2, y2) = xdiff + ydiff
    where xdiff = abs (x1 - x2)
          ydiff = abs (y1 - y2)

move :: Point -> Direction -> Int -> Point
move (x, y) North d = (x, y + d)
move (x, y) South d = (x, y - d)
move (x, y) East d = (x + d, y)
move (x, y) West d = (x - d, y)

step :: (Point, Direction) -> String -> (Point, Direction)
step (p, d) (t:amt) = ((move p d' amtAsInt), d')
    where d' = turn d t
          amtAsInt = read amt :: Int

turn :: Direction -> Char -> Direction
turn North 'L' = West
turn North 'R' = East
turn South 'L' = East
turn South 'R' = West
turn East 'L' = North
turn East 'R' = South
turn West 'L' = South
turn West 'R' = North

main = do
    input <- getContents
    print $ displacement origin $ fst $ foldl step (origin, North) (lines input)
