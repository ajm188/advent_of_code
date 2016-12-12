import CPU
import Instruction

startCPU = (0, (0, 0, 0, 0)) :: CPU

main = do
    input <- getContents
    print $
        lookupReg (eval (map read $ lines input) startCPU) "a"
