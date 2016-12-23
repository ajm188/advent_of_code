import Direction
import Move

nextButton :: Char -> String -> Char
nextButton button directions = foldl Move.move button $ map charToDirection directions

main = do
    input <- getContents
    print $
        tail $
        reverse $
        foldl (\code line -> (nextButton (head code) line) : code) ['5'] (lines input)
