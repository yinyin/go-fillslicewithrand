package gofillslicewithrand

import (
	"crypto/rand"
	"encoding/binary"
	"sync"
)

type Reader32 struct {
	lck      sync.Mutex
	rndBytes [32]byte
	nextByte int
}

func (seq *Reader32) prepareRandomBytes() (p []byte) {
	if seq.nextByte == 0 {
		if _, err := rand.Read(seq.rndBytes[:]); nil != err {
			mathRandRead(seq.rndBytes[:])
		}
	}
	p = seq.rndBytes[seq.nextByte:]
	seq.nextByte = (seq.nextByte + 4) & 31
	return
}

func (seq *Reader32) nolockInt32() (v int32) {
	p := seq.prepareRandomBytes()
	v = int32(binary.NativeEndian.Uint32(p) & ((1 << 31) - 1))
	return
}

func (seq *Reader32) Int32() int32 {
	seq.lck.Lock()
	defer seq.lck.Unlock()
	return seq.nolockInt32()
}

func (seq *Reader32) nolockUint32() (v uint32) {
	p := seq.prepareRandomBytes()
	v = binary.NativeEndian.Uint32(p)
	return
}

func (seq *Reader32) Uint32() uint32 {
	seq.lck.Lock()
	defer seq.lck.Unlock()
	return seq.nolockUint32()
}

type ReaderInt32Masked struct {
	Reader32
	Mask int32
}

func (seq *ReaderInt32Masked) Int32Masked() int32 {
	v := int32(seq.Uint32())
	return (v & seq.Mask)
}

type ReaderInt32N struct {
	Reader32
	n   int32
	max int32
}

func NewReaderInt32N(n int32) ReaderInt32N {
	return ReaderInt32N{
		n:   n,
		max: int32(((1 << 31) - 1) - (1<<31)%uint32(n)),
	}
}

func (seq *ReaderInt32N) Int32N() int32 {
	seq.lck.Lock()
	defer seq.lck.Unlock()
	v := seq.nolockInt32()
	for v > seq.max {
		v = seq.nolockInt32()
	}
	return (v % seq.n)
}
