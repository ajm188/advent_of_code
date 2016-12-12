module Instruction
( Instruction(..)
) where

data Instruction = Cpy String String
                   | Inc String
                   | Dec String
                   | Jnz String String
                   deriving (Show)

instance Read Instruction where

    readsPrec _ ('c':'p':'y':rs) = [(Cpy r1 r2, "")]
        where (r1:r2:_) = words rs
    readsPrec _ ('i':'n':'c':rs) = [(Inc r1, "")]
        where (r1:_) = words rs
    readsPrec _ ('d':'e':'c':rs) = [(Dec r1, "")]
        where (r1:_) = words rs
    readsPrec _ ('j':'n':'z':rs) = [(Jnz r1 r2, "")]
        where (r1:r2:_) = words rs
