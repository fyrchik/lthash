package lthash

import (
	"encoding/binary"
	"io"

	"golang.org/x/crypto/blake2b"
)

const (
	mask32A = 0xFFFFFFFF00000000
	mask32B = 0x00000000FFFFFFFF
)

type hash32 struct {
	d   blake2b.XOF
	sum [Size]byte
	cur [Size]byte
}

// New32 returns lthash which works with 32-bit numbers.
func New32() HomoHash {
	d, err := blake2b.NewXOF(Size, nil)
	if err != nil {
		panic(err)
	}

	return &hash32{
		d: d,
	}
}

// Add implements HomoHash interface.
func (h *hash32) Add(p []byte) {
	h.d.Reset()
	h.d.Write(p)
	io.ReadFull(h.d, h.cur[:])
	add32(&h.cur, &h.sum)
}

// Remove implements HomoHash interface.
func (h *hash32) Remove(p []byte) {
	h.d.Reset()
	h.d.Write(p)
	io.ReadFull(h.d, h.cur[:])
	sub32(&h.cur, &h.sum)
}

// Sum implements HomoHash interface.
func (h *hash32) Sum(b []byte) (cs []byte) {
	return append(b, h.sum[:]...)
}

// SetState implements HomoHash interface.
func (h *hash32) SetState(s []byte) {
	for i := range s {
		h.sum[i] = 0
	}
	copy(h.sum[:], s)
}

// add32Generic computes sum of successive 32-byte
// little-endian unsigned integers of a and b and saves them into b.
// For the sake of performance we add low and high halves simultaneously.
func add32Generic(a, b *[Size]byte) {
	for i := 0; i < Size; i += 8 {
		v1 := binary.LittleEndian.Uint64((*b)[i:])
		v2 := binary.LittleEndian.Uint64((*a)[i:])

		v1a := v1 & mask32A
		v1b := v1 & mask32B
		v2a := v2 & mask32A
		v2b := v2 & mask32B
		v1 = (v1a + v2a) & mask32A
		v2 = (v1b + v2b) & mask32B

		binary.LittleEndian.PutUint64((*b)[i:], v1|v2)
	}
}

// sub32Generic computes sum of successive 32-byte
// little-endian unsigned integers of a and b and saves them into b.
func sub32Generic(a, b *[Size]byte) {
	for i := 0; i < Size; i += 8 {
		v1 := binary.LittleEndian.Uint64((*b)[i:])
		v2 := binary.LittleEndian.Uint64((*a)[i:])

		v1a := v1 & mask32A
		v1b := v1 & mask32B
		v2a := v2 & mask32A
		v2b := v2 & mask32B
		v1 = (v1a + (mask32B - v2a)) & mask32A
		v2 = (v1b + (mask32A - v2b)) & mask32B

		binary.LittleEndian.PutUint64((*b)[i:], v1|v2)
	}
}
