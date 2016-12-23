import CPU
import Instruction

startCPU = (0, (7, 0, 0, 0)) :: CPU

main = do
    input <- getContents
    let instructions = map read (lines input) :: [Instruction]
    print $
        (lookupReg (snd $ eval (instructions, startCPU)) "a")
