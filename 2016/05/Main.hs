import qualified Data.Hash.MD5 as MD5

doorId = "ojvtpuvg"

hash :: String -> Int -> String
hash door x = (MD5.md5s . MD5.Str) (door ++ (show x))

password door len = (fst . (until ((\(p, _) -> length p == len)) next)) ("", 0)
    where next (p, i) = if (isInPassword . hash') i then (p ++ [((!! 5) . hash') i], i + 1) else (p, i + 1)
          isInPassword = ((=="00000") . (take 5))
          hash' = (hash door)

main = do
    print $ password doorId 8
