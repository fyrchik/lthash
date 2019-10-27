// +build !amd64

package lthash

func add16(a, b *[Size]byte) {
	add16Generic(a, b)
}

func sub16(a, b *[Size]byte) {
	sub16Generic(a, b)
}

func add32(a, b *[Size]byte) {
	add32Generic(a, b)
}

func sub32(a, b *[Size]byte) {
	sub32Generic(a, b)
}
