package util

import (
	"io"
	"net"
)

type IBytes interface {
	Len() int
	Bytes() []byte
	Reuse() bool
}

type Buffer []byte

func (Buffer) Reuse() bool {
	return false
}

func (b Buffer) Len() int {
	return len(b)
}

func (b Buffer) Cap() int {
	return cap(b)
}

func (b Buffer) SubBuf(start int, length int) Buffer {
	return b[start : start+length]
}

func (b *Buffer) Reset() {
	*b = b.SubBuf(0, 0)
}

func (b *Buffer) Malloc(count int) Buffer {
	l := b.Len()
	newL := l + count
	if newL > b.Cap() {
		n := make(Buffer, newL)
		copy(n, *b)
		*b = n
	} else {
		*b = b.SubBuf(0, newL)
	}
	return b.SubBuf(l, count)
}

func (b *Buffer) Relloc(count int) {
	b.Reset()
	b.Malloc(count)
}

func (b *Buffer) Split(n int) (result net.Buffers) {
	origin := *b
	for {
		if b.CanReadN(n) {
			result = append(result, b.ReadN(n))
		} else {
			result = append(result, *b)
			*b = origin
			return
		}
	}
}

func (b Buffer) CanReadN(n int) bool {
	return b.Len() >= n
}

func (b Buffer) CanRead() bool {
	return b.CanReadN(1)
}

func (b *Buffer) ReadN(n int) Buffer {
	l := len(*b)
	r := (*b)[:n]
	*b = (*b)[n:l]
	return r
}

func (b *Buffer) Read(buf []byte) (n int, err error) {
	if !b.CanReadN(len(buf)) {
		copy(buf, *b)
		return b.Len(), io.EOF
	}
	ret := b.ReadN(len(buf))
	copy(buf, ret)
	return len(ret), err
}

func (b *Buffer) Write(a []byte) (n int, err error) {
	l := b.Len()
	newL := l + len(a)
	if newL > b.Cap() {
		*b = append(*b, a...)
	} else {
		*b = b.SubBuf(0, newL)
		copy((*b)[l:], a)
	}
	return len(a), nil
}

func SizeOfBuffers[T ~[]byte](buf []T) (size int) {
	for _, b := range buf {
		size += len(b)
	}
	return
}

func ConcatBuffers[T ~[]byte](input []T) (out []byte) {
	for _, v := range input {
		out = append(out, v...)
	}
	return
}
