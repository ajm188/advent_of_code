import Operation

scramble :: (String -> [Operation] -> String)
scramble = foldl operate

unscramble :: (String -> [Operation] -> String)
unscramble = foldl undo

main = do
    input <- getContents
    print $
        scramble "abcdefgh" $
        map read $
        lines input
    print $
        unscramble "fbgdceah" $
        reverse $
        map read $
        lines input
