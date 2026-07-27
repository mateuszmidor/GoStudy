package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include <bzlib.h>
int bz2compress(bz_stream *s, int action, char *in, unsigned *inlen, char *out, unsigned *outlen);
*/
import "C"
import (
	"io"
	"unsafe"
)

type writer struct {
	w      io.Writer
	stream *C.bz_stream
	outbuf [64 * 1024]byte
}

// NewWriter creates a writing interface for bzip2 compressed sreams
func NewWriter(out io.Writer) io.WriteCloser {
	const (
		blockSize  = 9
		verbosity  = 0
		workFactor = 30
	)
	_ = new(int)
	w := &writer{w: out, stream: new(C.bz_stream)}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	if w.stream == nil {
		panic("writing stream is closed")
	}

	var total int // non-compressed bytes count that has been written

	for len(data) > 0 {
		inlen, outlen := C.uint(len(data)), C.uint(cap(w.outbuf))
		C.bz2compress(w.stream, C.BZ_RUN, (*C.char)(unsafe.Pointer(&data[0])), &inlen, (*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		total += int(inlen)
		data = data[inlen:]
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return total, err
		}
	}
	return total, nil
}

// Close flushes compressed data and closes the stream
func (w *writer) Close() error {
	if w.stream == nil {
		panic("writing stream is closed")
	}

	defer func() {
		C.BZ2_bzCompressEnd(w.stream)
		w.stream = nil
	}()

	for {
		inlen, outlen := C.uint(0), C.uint(cap(w.outbuf))
		outbuf := (*C.char)(unsafe.Pointer(&w.outbuf))
		r := directBZ2Compress(w.stream, C.BZ_FINISH, nil, &inlen, outbuf, &outlen)
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return err
		}
		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}

// Call the custom wrapper from bzip2.c
func wrappedBZ2Compress(stream *C.bz_stream, action C.int, inbuf *C.char, inlen *C.uint, outbuf *C.char, outlen *C.uint) C.int {
	return C.bz2compress(stream, action, inbuf, inlen, outbuf, outlen)
}

// Call the BZ2_bzCompress directly from bzlib.h
func directBZ2Compress(s *C.bz_stream, action C.int, inbuf *C.char, inlen *C.uint, outbuf *C.char, outlen *C.uint) C.int {
	s.next_in = inbuf
	s.avail_in = *inlen
	s.next_out = outbuf
	s.avail_out = *outlen
	r := C.BZ2_bzCompress(s, action)
	*inlen -= s.avail_in
	*outlen -= s.avail_out
	return r
}
