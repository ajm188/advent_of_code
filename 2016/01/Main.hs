type Point = (Int, Int)

data Direction = North | South | East | West

origin :: Point
origin = (0, 0)

displacement :: Point -> Point -> Int
displacement (x1, y1) (x2, y2) = xdiff + ydiff
    where xdiff = abs (x1 - x2)
          ydiff = abs (y1 - y2)

move :: Point -> Direction -> Int -> [Point]
move _ _ 0 = []
move (x, y) dir dist = (move p' dir (dist - 1)) ++ [p']
    where p' = case dir of North -> (x, y + 1)
                           South -> (x, y - 1)
                           East -> (x + 1, y)
                           West -> (x - 1, y)

step :: [(Point, Direction)] -> String -> [(Point, Direction)]
step [] _ = [(origin, North)]
step l@((p, d):_) (t:a) = points ++ l
    where points = map (\point -> (point, d')) $ move p d' a'
          d' = turn d t
          a' = read a :: Int

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
    let (endpoint:_) = foldl step [(origin, North)] (lines input) in
        print $ displacement origin $ fst $ endpoint
