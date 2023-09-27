package gofillslicewithrand

import (
	"crypto/rand"
	"encoding/binary"
	"sync"
)

type Reader64 struct {
	lck      sync.Mutex
	rndBytes [64]byte
	nextByte int
}

func (seq *Reader64) prepareRandomBytes() (p []byte) {
	if seq.nextByte == 0 {
		if _, err := rand.Read(seq.rndBytes[:]); nil != err {
			mathRandRead(seq.rndBytes[:])
		}
	}
	p = seq.rndBytes[seq.nextByte:]
	seq.nextByte = (seq.nextByte + 8) & 63
	return
}

func (seq *Reader64) nolockInt64() (v int64) {
	p := seq.prepareRandomBytes()
	v = int64(binary.NativeEndian.Uint64(p) & ((1 << 63) - 1))
	return
}

func (seq *Reader64) Int64() int64 {
	seq.lck.Lock()
	defer seq.lck.Unlock()
	return seq.nolockInt64()
}

func (seq *Reader64) nolockUint64() (v uint64) {
	p := seq.prepareRandomBytes()
	v = binary.NativeEndian.Uint64(p)
	return
}

func (seq *Reader64) Uint64() uint64 {
	seq.lck.Lock()
	defer seq.lck.Unlock()
	return seq.nolockUint64()
}

type ReaderInt64Masked struct {
	Reader64
	Mask int64
}

func (seq *ReaderInt64Masked) Int64Masked() int64 {
	v := int64(seq.Uint64())
	return (v & seq.Mask)
}

type ReaderInt64N struct {
	Reader64
	n   int64
	max int64
}

func NewReaderInt64N(n int64) ReaderInt64N {
	return ReaderInt64N{
		n:   n,
		max: int64(((1 << 63) - 1) - (1<<63)%uint64(n)),
	}
}

func (seq *ReaderInt64N) Int64N() int64 {
	seq.lck.Lock()
	defer seq.lck.Unlock()
	v := seq.nolockInt64()
	for v > seq.max {
		v = seq.nolockInt64()
	}
	return (v % seq.n)
}
