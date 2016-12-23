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

eval :: ([Instruction], CPU) -> ([Instruction], CPU)
eval state@(prog, cpu@(pc, _))
    | pc < length prog = eval $ exec (prog !! pc) (prog, cpu)
    | otherwise = state

exec :: Instruction -> ([Instruction], CPU) -> ([Instruction], CPU)
exec (Cpy r1 r2) state@(prog, cpu@(pc, rs@(a, b, c, d)))
    | notElem r2 registerNames = state -- invalid! skipping
    | elem r1 registerNames = exec (Cpy (show $ lookupReg cpu r1) r2) state
    | otherwise = (prog, (pc + 1, rs'))
        where rs' = case r2 of "a" -> (lit, b, c, d)
                               "b" -> (a, lit, c, d)
                               "c" -> (a, b, lit, d)
                               "d" -> (a, b, c, lit)
                               _ -> rs
              lit = read r1 :: Int
exec (Inc r) state@(_, cpu) = exec (Cpy (show $ (lookupReg cpu r) + 1) r) state
exec (Dec r) state@(_, cpu) = exec (Cpy (show $ (lookupReg cpu r) - 1) r) state
exec (Jnz r1 r2) state@(prog, cpu@(pc, rs@(a, b, c, d)))
    | elem r1 registerNames = exec (Jnz (show $ lookupReg cpu r1) r2) state
    | elem r2 registerNames = exec (Jnz r1 (show $ lookupReg cpu r2)) state
    | otherwise =
        case lit1 of 0 -> (prog, (pc + 1, rs))
                     _ -> (prog, (pc + lit2, rs))
        where lit1 = read r1 :: Int
              lit2 = read r2 :: Int
exec (Tgl r) state@(prog, cpu@(pc, rs@(a, b, c, d)))
    | elem r registerNames = ((toggle (index (lookupReg cpu r)) prog), (pc + 1, rs))
    | otherwise = ((toggle (index lit) prog), (pc + 1, rs))
        where lit = read r :: Int
              index offset = pc + offset
              toggle i prog = if elem i [0..((length prog) - 1)] then (take i prog) ++ [toggle' (prog !! i)] ++ (drop (i + 1) prog) else prog
              toggle' (Inc r) = Dec r
              toggle' (Dec r) = Inc r
              toggle' (Tgl r) = Inc r
              toggle' (Jnz r1 r2) = Cpy r1 r2
              toggle' (Cpy r1 r2) = Jnz r1 r2

lookupReg :: CPU -> String -> Int
lookupReg (_, (a, b, c, d)) r
    | r == "a" = a
    | r == "b" = b
    | r == "c" = c
    | r == "d" = d
    | otherwise = 0
