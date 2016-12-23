module Move2
( move
) where

import Direction

move :: Char -> Direction -> Char
move '1' d =
    case d of
        D -> '3'
        _ -> '1'
move '2' d =
    case d of
        R -> '3'
        D -> '6'
        _ -> '2'
move '3' d =
    case d of
        U -> '1'
        L -> '2'
        R -> '4'
        D -> '7'
move '4' d =
    case d of
        L -> '3'
        D -> '8'
        _ -> '4'
move '5' d =
    case d of
        R -> '6'
        _ -> '5'
move '6' d =
    case d of
        U -> '2'
        L -> '5'
        R -> '7'
        D -> 'A'
move '7' d =
    case d of
        U -> '3'
        L -> '6'
        R -> '8'
        D -> 'B'
move '8' d =
    case d of
        U -> '4'
        L -> '7'
        R -> '9'
        D -> 'C'
move '9' d =
    case d of
        L -> '8'
        _ -> '9'
move 'A' d =
    case d of
        U -> '6'
        R -> 'B'
        _ -> 'A'
move 'B' d =
    case d of
        U -> '7'
        L -> 'A'
        R -> 'C'
        D -> 'D' -- lol
move 'C' d =
    case d of
        U -> '8'
        L -> 'B'
        _ -> 'C'
move 'D' d =
    case d of
        U -> 'B'
        _ -> 'D'
