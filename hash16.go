package lthash

import (
	"encoding/binary"
	"io"

	"golang.org/x/crypto/blake2b"
)

const (
	mask16A = 0xFFFF0000FFFF0000
	mask16B = 0x0000FFFF0000FFFF
)

type hash16 struct {
	d   blake2b.XOF
	sum [Size]byte
	cur [Size]byte
}

// New16 returns lthash which works with 16-bit numbers.
func New16() HomoHash {
	d, err := blake2b.NewXOF(Size, nil)
	if err != nil {
		panic(err)
	}

	return &hash16{
		d: d,
	}
}

// Add implements HomoHash interface.
func (h *hash16) Add(p []byte) {
	h.d.Reset()
	h.d.Write(p)
	io.ReadFull(h.d, h.cur[:])
	add16(&h.cur, &h.sum)
}

// Sum implements HomoHash interface.
func (h *hash16) Sum(b []byte) (cs []byte) {
	return append(b, h.sum[:]...)
}

// Remove implements HomoHash interface.
func (h *hash16) Remove(p []byte) {
	h.d.Reset()
	h.d.Write(p)
	io.ReadFull(h.d, h.cur[:])
	sub16(&h.cur, &h.sum)
}

// SetState implements HomoHash interface.
func (h *hash16) SetState(s []byte) {
	for i := range s {
		h.sum[i] = 0
	}

	copy(h.sum[:], s)
}

// add16Generic computes sum of successive 16-byte
// little-endian unsigned integers of a and b and saves them into b.
// For the sake of performance we add 4 numbers at once:
//   uint64 = [uint16, 0, uint16, 0]
//   uint64 = [0, uint16, 0, uint16]
func add16Generic(a, b *[Size]byte) {
	for i := 0; i < Size; i += 8 {
		v1 := binary.LittleEndian.Uint64((*b)[i:])
		v2 := binary.LittleEndian.Uint64((*a)[i:])

		v1a := v1 & mask16A
		v1b := v1 & mask16B
		v2a := v2 & mask16A
		v2b := v2 & mask16B
		v1 = (v1a + v2a) & mask16A
		v2 = (v1b + v2b) & mask16B

		binary.LittleEndian.PutUint64((*b)[i:], v1|v2)
	}
}

// sub16Generic computes diff of successive 16-byte
// little-endian unsigned integers of a and b and saves them into b.
func sub16Generic(a, b *[Size]byte) {
	for i := 0; i < Size; i += 8 {
		v1 := binary.LittleEndian.Uint64((*b)[i:])
		v2 := binary.LittleEndian.Uint64((*a)[i:])

		v1a := v1 & mask16A
		v1b := v1 & mask16B
		v2a := v2 & mask16A
		v2b := v2 & mask16B
		v1 = (v1a + (mask16B - v2a)) & mask16A
		v2 = (v1b + (mask16A - v2b)) & mask16B

		binary.LittleEndian.PutUint64((*b)[i:], v1|v2)
	}
}
