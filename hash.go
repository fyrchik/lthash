package lthash

// HomoHash interface defines a set of operations
// for any homomorphic hash.
type HomoHash interface {
	// Add adds p to the multiset.
	Add(p []byte)

	// Remove removes p from the multiset.
	Remove(p []byte)

	// Sum appends the hash of the current multiset to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// SetState initializes current state with s.
	SetState(s []byte)
}

// Size is a length of a checksum.
const Size = 2048
