import qualified Data.Hash.MD5 as MD5

hash :: String -> Int -> String
hash salt x = (MD5.md5s . MD5.Str) (salt ++ (show x))

index :: String -> Int -> Int
index salt keys = (fst (until ((==keys) . (snd)) (next) (0, 0))) - 1
    where next (i, keys') = if (key (hash' i) (i + 1)) then (i + 1, keys' + 1) else (i + 1, keys')
          hash' = (hash salt)
          key blob i = case maybeTriple blob of Just triple -> (any ((hasConsecutive 5 triple) . hash')) [i..(i + 1000)]
                                                Nothing     -> False

maybeTriple :: Eq a => [a] -> Maybe a
maybeTriple (x:xs)
    | length xs < 2 = Nothing
    | ((all (==x)) . (take 2)) xs = Just x
    | otherwise = maybeTriple xs

hasConsecutive :: Eq a => Int -> a -> [a] -> Bool
hasConsecutive n t xs
    | length xs < n = False
    | (all (==t)) (take n xs) = True
    | otherwise = hasConsecutive n t (tail xs)

main = do
    let (salt, numKeys) = ("yjdafjpo", 64)
    print $ index salt numKeys
