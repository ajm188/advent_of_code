import Operation

scramble :: (String -> [Operation] -> String)
scramble = foldl operate

main = do
    input <- getContents
    print $
        scramble "abcdefgh" $
        map read $
        lines input
