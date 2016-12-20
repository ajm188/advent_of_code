allIPs = [0..4294967295]

inRange :: Ord a => a -> (a, a) -> Bool
inRange x (y, z) = x >= y && x <= z

lowestAllowedIP :: [Int] -> [String] -> Int
lowestAllowedIP allIPs lines = (head . (filter (\x -> all (not . (inRange x)) blocked))) allIPs
    where blocked = blockedIPs lines

blockedIPs :: [String] -> [(Int, Int)]
blockedIPs [] = []
blockedIPs (ip:ips) = (low, high) : (blockedIPs ips)
    where (low, high) = blockedRange ip

blockedRange :: String -> (Int, Int)
blockedRange line = (read f, (read . tail) s)
    where (f, s) = break (=='-') line

main = do
    input <- getContents
    print $
        lowestAllowedIP allIPs $
        lines input
