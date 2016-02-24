// func FastRand(x int) int
TEXT ·FastRand(SB),$0
    MOVQ x+0(FP), BX
    SUBQ $1, BX
    RDTSC
    ANDQ BX, AX
    MOVQ AX, ret+8(FP)
    RET
