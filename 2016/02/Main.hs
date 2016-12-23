import Direction
import Move

passcode :: ((Char -> String -> Char) -> [String] -> String)
passcode buttonFunc = (tail . reverse . (foldl (\code line -> (buttonFunc (head code) line) : code) ['5']))

nextButton :: Char -> String -> Char
nextButton button directions = foldl Move.move button $ map charToDirection directions

main = do
    input <- getContents
    print $
        passcode nextButton $
        lines input
