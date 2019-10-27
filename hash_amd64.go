// +build amd64

package lthash

import "github.com/fyrchik/lthash/internal/cpuid"

var (
	useAVX  = cpuid.HasAVX()
	useAVX2 = cpuid.HasAVX() && cpuid.HasAVX2()
)

//go:noescape
func add16AVX(a, b *[Size]byte)

//go:noescape
func sub16AVX(a, b *[Size]byte)

//go:noescape
func add16AVX2(a, b *[Size]byte)

//go:noescape
func sub16AVX2(a, b *[Size]byte)

//go:noescape
func add32AVX(a, b *[Size]byte)

//go:noescape
func sub32AVX(a, b *[Size]byte)

//go:noescape
func add32AVX2(a, b *[Size]byte)

//go:noescape
func sub32AVX2(a, b *[Size]byte)

func add16(a, b *[Size]byte) {
	switch {
	case useAVX2:
		add16AVX2(a, b)
	case useAVX:
		add16AVX(a, b)
	default:
		add16Generic(a, b)
	}
}

func sub16(a, b *[Size]byte) {
	switch {
	case useAVX2:
		sub16AVX2(a, b)
	case useAVX:
		sub16AVX(a, b)
	default:
		sub16Generic(a, b)
	}
}

func add32(a, b *[Size]byte) {
	switch {
	case useAVX2:
		add32AVX2(a, b)
	case useAVX:
		add32AVX(a, b)
	default:
		add32Generic(a, b)
	}
}

func sub32(a, b *[Size]byte) {
	switch {
	case useAVX2:
		sub32AVX2(a, b)
	case useAVX:
		sub32AVX(a, b)
	default:
		sub32Generic(a, b)
	}
}
