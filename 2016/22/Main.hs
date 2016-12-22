import Node
import Util

pathPrefix :: String
pathPrefix = "/dev/grid/"

nodeName :: (String -> String)
nodeName = (drop (length pathPrefix))

viablePairs :: [Node] -> [(Node, Node)]
viablePairs nodes = concat (map (\(n, i) -> viablePairs' n ((take (i - 1) nodes) ++ (drop i nodes))) (zip nodes [1..]))
    where viablePairs' node = ((map (\node' -> (node, node'))) . (filter (viablePair node)))

main = do
    input <- getContents
    let lines' = ((drop 2) . lines) input
    print $
        length $
        viablePairs $
        map (\(name:_:used:avail:_) -> newNode (nodeName name) (read (takeWhile isDigit used)) (read (takeWhile isDigit avail))) $
        map words $
        lines'
