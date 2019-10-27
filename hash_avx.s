#include "textflag.h"

#define mod(n, op) \
    MOVUPD +n(AX), X0 \
    MOVUPD +n(BX), X1 \
    op X0, X1 \
    MOVUPD X1, +n(BX)

#define modV(n, op) \
    MOVUPD +n(AX), X0 \
    MOVUPD +n(BX), X1 \
    op X0, X1, X1 \
    MOVUPD X1, +n(BX)

#define apply(modify, op) \
    modify(0, op) \
    modify(16, op) \
    modify(32, op) \
    modify(48, op) \
    modify(64, op) \
    modify(80, op) \
    modify(96, op) \
    modify(112, op) \
    modify(128, op) \
    modify(144, op) \
    modify(160, op) \
    modify(176, op) \
    modify(192, op) \
    modify(208, op) \
    modify(224, op) \
    modify(240, op) \
    modify(256, op) \
    modify(272, op) \
    modify(288, op) \
    modify(304, op) \
    modify(320, op) \
    modify(336, op) \
    modify(352, op) \
    modify(368, op) \
    modify(384, op) \
    modify(400, op) \
    modify(416, op) \
    modify(432, op) \
    modify(448, op) \
    modify(464, op) \
    modify(480, op) \
    modify(496, op) \
    modify(512, op) \
    modify(528, op) \
    modify(544, op) \
    modify(560, op) \
    modify(576, op) \
    modify(592, op) \
    modify(608, op) \
    modify(624, op) \
    modify(640, op) \
    modify(656, op) \
    modify(672, op) \
    modify(688, op) \
    modify(704, op) \
    modify(720, op) \
    modify(736, op) \
    modify(752, op) \
    modify(768, op) \
    modify(784, op) \
    modify(800, op) \
    modify(816, op) \
    modify(832, op) \
    modify(848, op) \
    modify(864, op) \
    modify(880, op) \
    modify(896, op) \
    modify(912, op) \
    modify(928, op) \
    modify(944, op) \
    modify(960, op) \
    modify(976, op) \
    modify(992, op) \
    modify(1008, op) \
    modify(1024, op) \
    modify(1040, op) \
    modify(1056, op) \
    modify(1072, op) \
    modify(1088, op) \
    modify(1104, op) \
    modify(1120, op) \
    modify(1136, op) \
    modify(1152, op) \
    modify(1168, op) \
    modify(1184, op) \
    modify(1200, op) \
    modify(1216, op) \
    modify(1232, op) \
    modify(1248, op) \
    modify(1264, op) \
    modify(1280, op) \
    modify(1296, op) \
    modify(1312, op) \
    modify(1328, op) \
    modify(1344, op) \
    modify(1360, op) \
    modify(1376, op) \
    modify(1392, op) \
    modify(1408, op) \
    modify(1424, op) \
    modify(1440, op) \
    modify(1456, op) \
    modify(1472, op) \
    modify(1488, op) \
    modify(1504, op) \
    modify(1520, op) \
    modify(1536, op) \
    modify(1552, op) \
    modify(1568, op) \
    modify(1584, op) \
    modify(1600, op) \
    modify(1616, op) \
    modify(1632, op) \
    modify(1648, op) \
    modify(1664, op) \
    modify(1680, op) \
    modify(1696, op) \
    modify(1712, op) \
    modify(1728, op) \
    modify(1744, op) \
    modify(1760, op) \
    modify(1776, op) \
    modify(1792, op) \
    modify(1808, op) \
    modify(1824, op) \
    modify(1840, op) \
    modify(1856, op) \
    modify(1872, op) \
    modify(1888, op) \
    modify(1904, op) \
    modify(1920, op) \
    modify(1936, op) \
    modify(1952, op) \
    modify(1968, op) \
    modify(1984, op) \
    modify(2000, op) \
    modify(2016, op) \
    modify(2032, op)

// func add16AVX(a, b *[2048]byte)
TEXT ·add16AVX(SB),NOSPLIT,$0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX

    apply(mod, PADDW)

    RET

// func sub16AVX(a, b *[2048]byte)
TEXT ·sub16AVX(SB),NOSPLIT,$0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX

    apply(mod, PSUBW)

    RET


// func add32AVX(a, b *[2048]byte)
TEXT ·add32AVX(SB),NOSPLIT,$0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX

    apply(modV, VPADDD)

    RET

// func sub32AVX(a, b *[2048]byte)
TEXT ·sub32AVX(SB),NOSPLIT,$0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX

    apply(modV, VPSUBD)

    RET
