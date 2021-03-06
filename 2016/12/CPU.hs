module CPU
( CPU
, eval
, exec
, lookupReg
) where

import Instruction

type ProgramCounter = Int
type Registers = (Int, Int, Int, Int)
type CPU = (ProgramCounter, Registers)

registerNames = ["a", "b", "c", "d"]

eval :: [Instruction] -> CPU -> CPU
eval prog cpu@(pc, _)
    | pc < length prog = eval prog $ exec (prog !! pc) cpu
    | otherwise = cpu

exec :: Instruction -> CPU -> CPU
exec (Cpy r1 r2) cpu@(pc, rs@(a, b, c, d))
    | elem r1 registerNames = exec (Cpy (show $ lookupReg cpu r1) r2) cpu
    | otherwise = (pc + 1, rs')
        where rs' = case r2 of "a" -> (lit, b, c, d)
                               "b" -> (a, lit, c, d)
                               "c" -> (a, b, lit, d)
                               "d" -> (a, b, c, lit)
                               _ -> rs
              lit = read r1 :: Int
exec (Inc r) cpu = exec (Cpy (show $ (lookupReg cpu r) + 1) r) cpu
exec (Dec r) cpu = exec (Cpy (show $ (lookupReg cpu r) - 1) r) cpu
exec (Jnz r1 r2) cpu@(pc, rs@(a, b, c, d))
    | elem r1 registerNames = exec (Jnz (show $ lookupReg cpu r1) r2) cpu
    | elem r2 registerNames = exec (Jnz r1 (show $ lookupReg cpu r2)) cpu
    | otherwise =
        case lit1 of 0 -> (pc + 1, rs)
                     _ -> (pc + lit2, rs)
        where lit1 = read r1 :: Int
              lit2 = read r2 :: Int

lookupReg :: CPU -> String -> Int
lookupReg (_, (a, b, c, d)) r
    | r == "a" = a
    | r == "b" = b
    | r == "c" = c
    | r == "d" = d
    | otherwise = 0
