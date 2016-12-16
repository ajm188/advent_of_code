data Rep = Repeated Int String
           | Raw Char
data Sequence = Sequence [Rep]

reps :: Sequence -> [Rep]
reps (Sequence reps') = reps'

parse :: String -> Sequence
parse [] = Sequence []
parse ('(':xs) = Sequence ((Repeated repCount content) : (reps $ parse rest))
    where contentLen = read (takeWhile (/= 'x') xs) :: Int
          afterContentLen = drop 1 $ dropWhile (/= 'x') xs
          repCount = read (takeWhile (/= ')') afterContentLen) :: Int
          afterRepCount = drop 1 $ dropWhile (/= ')') afterContentLen
          content = take contentLen afterRepCount
          rest = drop contentLen afterRepCount
parse (x:xs) = Sequence ((Raw x) : (reps $ parse xs))

instance Show Rep where
    show (Repeated repCount content) = foldl (++) "" [x | (x, _) <- zip (repeat content) [1..repCount]]
    show (Raw x) = x : ""

instance Show Sequence where
    show (Sequence seq) = foldl (++) "" $ map show seq

main = do
    input <- getContents
    print $
        sum $
        map (length . show . parse) $
        lines input
