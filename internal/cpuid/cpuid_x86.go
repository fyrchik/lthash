// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
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

// +build amd64

package cpuid

const (
	bitOSXSAVE = 1 << 27
	bitAVX     = 1 << 28
	bitAVX2    = 1 << 5
)

var (
	hasAVX  bool
	hasAVX2 bool
)

func init() {
	maxID, _, _, _ := cpuid(0, 0)
	if maxID < 1 {
		return
	}

	_, _, ecx1, _ := cpuid(1, 0)
	hasOSXSAVE := isSet(ecx1, bitOSXSAVE)

	osSupportsAVX := false
	if hasOSXSAVE {
		eax, _ := xgetbv()
		osSupportsAVX = isSet(eax, 1<<1) && isSet(eax, 1<<2)
	}

	hasAVX = isSet(ecx1, bitAVX) && osSupportsAVX

	if maxID < 7 {
		return
	}

	_, ebx7, _, _ := cpuid(7, 0)
	hasAVX2 = isSet(ebx7, bitAVX2) && osSupportsAVX
}

// HasAVX returns true iff CPU supports AVX instruction set.
func HasAVX() bool { return hasAVX }

// HasAVX2 returns true iff CPU supports AVX2 instruction set.
func HasAVX2() bool { return hasAVX2 }

func isSet(hwc uint32, value uint32) bool {
	return hwc&value != 0
}

func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)
func xgetbv() (eax, edx uint32)
