import Data.Ix

allIPs = [0..4294967295]

rangeOverlaps :: Ix a => (a, a) -> (a, a) -> Bool
rangeOverlaps r1@(low, high) r2@(low', high') = (any (inRange r2) [low, high]) || (any (inRange r1) [low', high'])

mergeRanges :: Ix a => [(a, a)] -> [(a, a)]
mergeRanges rs = fst (until done next (rs, []))
    where done (a, b) = a == b
          next ([], _) = ([], [])
          next (r'@(r:rs), _) = ((foldl next' r (filter (rangeOverlaps r) rs)) : (mergeRanges (filter (not . (rangeOverlaps r)) rs)), r')
          next' (low, high) (low', high') = (minimum [low, low'], maximum [high, high'])

numAllowedIPs :: [Int] -> [String] -> Int
numAllowedIPs allIPs lines = maxIP - totalBlocked + 1
    where blocked = (mergeRanges . blockedIPs) lines
          totalBlocked = (sum . (map (\(low, high) -> (high - low + 1)))) blocked

lowestAllowedIP :: [Int] -> [String] -> Int
lowestAllowedIP allIPs lines = (head . (filter (\x -> all (not . (inRange' x)) blocked))) allIPs
    where blocked = blockedIPs lines
          inRange' x r = inRange r x

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
    print $
        numAllowedIPs allIPs $
        lines input
