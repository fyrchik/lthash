# LtHash

LtHash is a family of functions defined by Bellare and Micciancio in
[[1](https://cseweb.ucsd.edu/~daniele/papers/IncHash.pdf)] and implemented by engineers from Facebook in
[[2](https://eprint.iacr.org/2019/227.pdf),
[3](https://github.com/facebook/folly/tree/master/folly/experimental/crypto)].

This algorithm calculates hash of a set. _Homomorphic_ hashing means that hash of a whole set can be
calculated based on hashes of it's subsets, e.g. `h({a,b}) = h({a})*h({b})`.

In essense this algorithm calculates `blake2b` hash of every element and then adds vectors as if
they were a collection of `uint16` (`uint32`) values. On the byte level LittleEndian representation is used.


# Usage

```golang
h := lthash.New16()

// Add 3 objects
h.Add([]byte("John"))
h.Add([]byte("Maria"))

// remember the hash of the set {"John", "Maria"}
sum := h.Checksum()

// after we add Anna to our party, hash should change.
h.Add([]byte("Anna"))

// after removing an element we have a set {"John", "Maria"}
// once again so a hash should be the same.
h.Remove([]byte("Anna"))
if !bytes.Equal(sum, h.Checksum()) {
    panic("unexpected")
}
```

# Notes
1. Actually this algorithm calculates hash of a [multiset](https://en.wikipedia.org/wiki/Multiset),
not a set. This means that any element can be added multiple times
and every time checksum will be changed.
2. Technically there is nothing wrong in removing element which you have not previously added.
You will get no error in this case, so be careful.

# Benchmarks
TODO

# Other implementations
- [Facebook implementation](https://github.com/facebook/folly/tree/master/folly/experimental/crypto)
- [Another golang implementation](https://github.com/lukechampine/lthash) (it is slightly slower)
