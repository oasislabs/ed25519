// Copyright (c) 2019 Oasis Labs Inc.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//   * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Oasis Labs Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

#include "textflag.h"

// Upstream: `ed25519-donna-64bit-x86.h`

// func scalarmultBaseChooseNielsAMD64(u uint64, table *byte, t *ge25519niels, sign uint64)
TEXT ·scalarmultBaseChooseNielsAMD64(SB),NOSPLIT,$0-32
	MOVQ table+8(FP), R14
	MOVQ t+16(FP), R15

	// ysubx+xaddy+t2d
	MOVQ u+0(FP), AX
	MOVD AX, X14
	PSHUFD $0x00, X14, X14
	PXOR X0, X0
	PXOR X1, X1
	PXOR X2, X2
	PXOR X3, X3
	PXOR X4, X4
	PXOR X5, X5

	// 0
	MOVQ $0, AX
	MOVD AX, X15
	PSHUFD $0x00, X15, X15
	PCMPEQL X14, X15
	MOVQ $1, AX
	MOVD AX, X6
	PXOR X7, X7
	PAND X15, X6
	PAND X15, X7
	POR X6, X0
	POR X7, X1
	POR X6, X2
	POR X7, X3

	// 1
	MOVQ $1, AX
	MOVD AX, X15
	PSHUFD $0x00, X15, X15
	PCMPEQL X14, X15
	MOVOU 0(R14), X6
	MOVOU 16(R14), X7
	MOVOU 32(R14), X8
	MOVOU 48(R14), X9
	MOVOU 64(R14), X10
	MOVOU 80(R14), X11
	PAND X15, X6
	PAND X15, X7
	PAND X15, X8
	PAND X15, X9
	PAND X15, X10
	PAND X15, X11
	POR X6, X0
	POR X7, X1
	POR X8, X2
	POR X9, X3
	POR X10, X4
	POR X11, X5

	// 2
	MOVQ $2, AX
	MOVD AX, X15
	PSHUFD $0x00, X15, X15
	PCMPEQL X14, X15
	MOVOU 96(R14), X6
	MOVOU 112(R14), X7
	MOVOU 128(R14), X8
	MOVOU 144(R14), X9
	MOVOU 160(R14), X10
	MOVOU 176(R14), X11
	PAND X15, X6
	PAND X15, X7
	PAND X15, X8
	PAND X15, X9
	PAND X15, X10
	PAND X15, X11
	POR X6, X0
	POR X7, X1
	POR X8, X2
	POR X9, X3
	POR X10, X4
	POR X11, X5

	// 3
	MOVQ $3, AX
	MOVD AX, X15
	PSHUFD $0x00, X15, X15
	PCMPEQL X14, X15
	MOVOU 192(R14), X6
	MOVOU 208(R14), X7
	MOVOU 224(R14), X8
	MOVOU 240(R14), X9
	MOVOU 256(R14), X10
	MOVOU 272(R14), X11
	PAND X15, X6
	PAND X15, X7
	PAND X15, X8
	PAND X15, X9
	PAND X15, X10
	PAND X15, X11
	POR X6, X0
	POR X7, X1
	POR X8, X2
	POR X9, X3
	POR X10, X4
	POR X11, X5

	// 4
	MOVQ $4, AX
	MOVD AX, X15
	PSHUFD $0x00, X15, X15
	PCMPEQL X14, X15
	MOVOU 288(R14), X6
	MOVOU 304(R14), X7
	MOVOU 320(R14), X8
	MOVOU 336(R14), X9
	MOVOU 352(R14), X10
	MOVOU 368(R14), X11
	PAND X15, X6
	PAND X15, X7
	PAND X15, X8
	PAND X15, X9
	PAND X15, X10
	PAND X15, X11
	POR X6, X0
	POR X7, X1
	POR X8, X2
	POR X9, X3
	POR X10, X4
	POR X11, X5

	// 5
	MOVQ $5, AX
	MOVD AX, X15
	PSHUFD $0x00, X15, X15
	PCMPEQL X14, X15
	MOVOU 384(R14), X6
	MOVOU 400(R14), X7
	MOVOU 416(R14), X8
	MOVOU 432(R14), X9
	MOVOU 448(R14), X10
	MOVOU 464(R14), X11
	PAND X15, X6
	PAND X15, X7
	PAND X15, X8
	PAND X15, X9
	PAND X15, X10
	PAND X15, X11
	POR X6, X0
	POR X7, X1
	POR X8, X2
	POR X9, X3
	POR X10, X4
	POR X11, X5

	// 6
	MOVQ $6, AX
	MOVD AX, X15
	PSHUFD $0x00, X15, X15
	PCMPEQL X14, X15
	MOVOU 480(R14), X6
	MOVOU 496(R14), X7
	MOVOU 512(R14), X8
	MOVOU 528(R14), X9
	MOVOU 544(R14), X10
	MOVOU 560(R14), X11
	PAND X15, X6
	PAND X15, X7
	PAND X15, X8
	PAND X15, X9
	PAND X15, X10
	PAND X15, X11
	POR X6, X0
	POR X7, X1
	POR X8, X2
	POR X9, X3
	POR X10, X4
	POR X11, X5

	// 7
	MOVQ $7, AX
	MOVD AX, X15
	PSHUFD $0x00, X15, X15
	PCMPEQL X14, X15
	MOVOU 576(R14), X6
	MOVOU 592(R14), X7
	MOVOU 608(R14), X8
	MOVOU 624(R14), X9
	MOVOU 640(R14), X10
	MOVOU 656(R14), X11
	PAND X15, X6
	PAND X15, X7
	PAND X15, X8
	PAND X15, X9
	PAND X15, X10
	PAND X15, X11
	POR X6, X0
	POR X7, X1
	POR X8, X2
	POR X9, X3
	POR X10, X4
	POR X11, X5

	// 8
	MOVQ $8, AX
	MOVD AX, X15
	PSHUFD $0x00, X15, X15
	PCMPEQL X14, X15
	MOVOU 672(R14), X6
	MOVOU 688(R14), X7
	MOVOU 704(R14), X8
	MOVOU 720(R14), X9
	MOVOU 736(R14), X10
	MOVOU 752(R14), X11
	PAND X15, X6
	PAND X15, X7
	PAND X15, X8
	PAND X15, X9
	PAND X15, X10
	PAND X15, X11
	POR X6, X0
	POR X7, X1
	POR X8, X2
	POR X9, X3
	POR X10, X4
	POR X11, X5

	// conditionally swap ysubx and xaddy
	MOVQ sign+24(FP), AX
	XORQ $1, AX
	MOVD AX, X14
	PXOR X15, X15
	PSHUFD $0x00, X14, X14
	PXOR X0, X2
	PXOR X1, X3
	PCMPEQL X14, X15
	MOVOU X2, X6
	MOVOU X3, X7
	PAND X15, X6
	PAND X15, X7
	PXOR X6, X0
	PXOR X7, X1
	PXOR X0, X2
	PXOR X1, X3

	// store ysubx
	MOVQ $0x7ffffffffffff, AX
	MOVD X0, CX
	MOVD X0, R8
	MOVD X1, SI
	PSHUFD $0xee, X0, X0
	PSHUFD $0xee, X1, X1
	MOVD X0, DX
	MOVD X1, DI
	SHRQ $51, DX, R8
	SHRQ $38, SI, DX
	SHRQ $25, DI, SI
	SHRQ $12, DI
	ANDQ AX, CX
	ANDQ AX, R8
	ANDQ AX, DX
	ANDQ AX, SI
	ANDQ AX, DI
	MOVQ CX, 0(R15)
	MOVQ R8, 8(R15)
	MOVQ DX, 16(R15)
	MOVQ SI, 24(R15)
	MOVQ DI, 32(R15)

	// store xaddy
	MOVQ $0x7ffffffffffff, AX
	MOVD X2, CX
	MOVD X2, R8
	MOVD X3, SI
	PSHUFD $0xee, X2, X2
	PSHUFD $0xee, X3, X3
	MOVD X2, DX
	MOVD X3, DI
	SHRQ $51, DX, R8
	SHRQ $38, SI, DX
	SHRQ $25, DI, SI
	SHRQ $12, DI
	ANDQ AX, CX
	ANDQ AX, R8
	ANDQ AX, DX
	ANDQ AX, SI
	ANDQ AX, DI
	MOVQ CX, 40(R15)
	MOVQ R8, 48(R15)
	MOVQ DX, 56(R15)
	MOVQ SI, 64(R15)
	MOVQ DI, 72(R15)

	// extract t2d
	MOVQ $0x7ffffffffffff, AX
	MOVD X4, CX
	MOVD X4, R8
	MOVD X5, SI
	PSHUFD $0xee, X4, X4
	PSHUFD $0xee, X5, X5
	MOVD X4, DX
	MOVD X5, DI
	SHRQ $51, DX, R8
	SHRQ $38, SI, DX
	SHRQ $25, DI, SI
	SHRQ $12, DI
	ANDQ AX, CX
	ANDQ AX, R8
	ANDQ AX, DX
	ANDQ AX, SI
	ANDQ AX, DI

	// conditionally negate t2d
	MOVQ sign+24(FP), AX
	MOVQ $0xfffffffffffda, R9
	MOVQ $0xffffffffffffe, R10
	MOVQ R10, R11
	MOVQ R10, R12
	MOVQ R10, R13
	SUBQ CX, R9
	SUBQ R8, R10
	SUBQ DX, R11
	SUBQ SI, R12
	SUBQ DI, R13
	CMPQ AX, $1
	CMOVQEQ R9, CX
	CMOVQEQ R10, R8
	CMOVQEQ R11, DX
	CMOVQEQ R12, SI
	CMOVQEQ R13, DI

	// store t2d
	MOVQ CX, 80(R15)
	MOVQ R8, 88(R15)
	MOVQ DX, 96(R15)
	MOVQ SI, 104(R15)
	MOVQ DI, 112(R15)

	RET