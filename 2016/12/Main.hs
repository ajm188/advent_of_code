import CPU
import Instruction

startCPU = (0, (0, 0, 0, 0)) :: CPU
startCPUPart2 = (0, (0, 0, 1, 0)) :: CPU

main = do
    input <- getContents
    print $
        ( lookupReg (eval (map read $ lines input) startCPU) "a"
        , lookupReg (eval (map read $ lines input) startCPUPart2) "a"
        )
