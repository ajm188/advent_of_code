data Disc = Disc Int Int

earliestTime :: [Disc] -> Int
earliestTime discs = snd (until goal rotateAll (discs, 0))
    where goal (discs, _) = ((all capsuleCanFall) . (map (\(d, n) -> rotateN d n))) (zip discs [1..(length discs)])
          rotateAll (discs, t) = (map rotate discs, t + 1)

capsuleCanFall :: Disc -> Bool
capsuleCanFall (Disc _ current) = current == 0

rotate :: Disc -> Disc
rotate (Disc total current) = Disc total (mod (current + 1) total)

rotateN :: Disc -> Int -> Disc
rotateN disc 0 = disc
rotateN disc n = rotateN (rotate disc) (n - 1)

newDisc :: String -> Disc
newDisc line = Disc total current
    where numbers = ((map read) . words . (filter spaceOrNum)) line
          spaceOrNum x = x == ' ' || elem x ['0'..'9']
          (total, current) = (numbers !! 1, numbers !! 3)

main = do
    input <- getContents
    print $
        earliestTime $
        map newDisc $
        lines input
    print $ -- part 2
        earliestTime $
        (map newDisc $
        lines input) ++ [Disc 11 0]
