import qualified Data.Hash.MD5 as MD5

doorId = "ojvtpuvg"

hash :: String -> Int -> String
hash door x = (MD5.md5s . MD5.Str) (door ++ (show x))

password door len = (fst . (until ((\(p, _) -> length p == len)) next)) ("", 0)
    where next (p, i) = if (isInPassword . hash') i then (p ++ [((!! 5) . hash') i], i + 1) else (p, i + 1)
          isInPassword = ((=="00000") . (take 5))
          hash' = (hash door)

password2 door = ((map fst) . fst . (until (\(pass', _) -> foundPassword2 pass') next)) (zip ['0'..'7'] (repeat False), 0)
    where next (pass', i) = if (isInPassword . hash') i then ((updatePassword2 pass') (hash' i), i + 1) else (pass', i + 1)
          isInPassword = ((=="00000") . (take 5))
          hash' = (hash door)

updatePassword2 :: [(Char, Bool)] -> String -> [(Char, Bool)]
updatePassword2 pass' h
    | validPosition = updatePassword2'
    | otherwise = pass'
    where posChar = h !! 5
          pos = (read . (:"")) posChar
          validPosition = elem posChar ['0'..'7'] && (not . snd) (pass' !! pos)
          updatePassword2' = take pos pass' ++ [(h !! 6, True)] ++ drop (pos + 1) pass'

foundPassword2 :: ([(Char, Bool)] -> Bool)
foundPassword2 = (all id) . (map snd)

main = do
    print $ password doorId 8
    print $ password2 doorId
