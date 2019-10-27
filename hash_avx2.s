#include "textflag.h"

#define modify(n, op) \
    VMOVUPS +n(AX), Y0 \
    VMOVUPS +n(BX), Y1 \
    op Y0, Y1, Y1 \
    VMOVUPS Y1, +n(BX)

#define apply(op) \
    modify(0, op) \
    modify(32, op) \
    modify(64, op) \
    modify(96, op) \
    modify(128, op) \
    modify(160, op) \
    modify(192, op) \
    modify(224, op) \
    modify(256, op) \
    modify(288, op) \
    modify(320, op) \
    modify(352, op) \
    modify(384, op) \
    modify(416, op) \
    modify(448, op) \
    modify(480, op) \
    modify(512, op) \
    modify(544, op) \
    modify(576, op) \
    modify(608, op) \
    modify(640, op) \
    modify(672, op) \
    modify(704, op) \
    modify(736, op) \
    modify(768, op) \
    modify(800, op) \
    modify(832, op) \
    modify(864, op) \
    modify(896, op) \
    modify(928, op) \
    modify(960, op) \
    modify(992, op) \
    modify(1024, op) \
    modify(1056, op) \
    modify(1088, op) \
    modify(1120, op) \
    modify(1152, op) \
    modify(1184, op) \
    modify(1216, op) \
    modify(1248, op) \
    modify(1280, op) \
    modify(1312, op) \
    modify(1344, op) \
    modify(1376, op) \
    modify(1408, op) \
    modify(1440, op) \
    modify(1472, op) \
    modify(1504, op) \
    modify(1536, op) \
    modify(1568, op) \
    modify(1600, op) \
    modify(1632, op) \
    modify(1664, op) \
    modify(1696, op) \
    modify(1728, op) \
    modify(1760, op) \
    modify(1792, op) \
    modify(1824, op) \
    modify(1856, op) \
    modify(1888, op) \
    modify(1920, op) \
    modify(1952, op) \
    modify(1984, op) \
    modify(2016, op)

// func add16AVX2(a, b *[2048]byte)
TEXT ·add16AVX2(SB),NOSPLIT,$0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX

    apply(VPADDW)

    RET

// func sub16AVX2(a, b *[2048]byte)
TEXT ·sub16AVX2(SB),NOSPLIT,$0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX

    apply(VPSUBW)

    RET

// func add16AVX2(a, b *[2048]byte)
TEXT ·add32AVX2(SB),NOSPLIT,$0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX

    apply(VPADDD)

    RET

// func sub16AVX2(a, b *[2048]byte)
TEXT ·sub32AVX2(SB),NOSPLIT,$0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX

    apply(VPSUBD)

    RET
