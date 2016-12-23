import CPU
import Instruction

startCPU = (0, (0, 0, 0, 0)) :: CPU
startCPU2 = (0, (0, 0, 1, 0)) :: CPU

main = do
    input <- getContents
    let instructions = map read (lines input) :: [Instruction]
    print $
        (lookupReg (snd $ eval (instructions, startCPU)) "a")
    print $
        (lookupReg (snd $ eval (instructions, startCPU2)) "a")
