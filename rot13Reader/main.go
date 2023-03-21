package main

import (
	"io"
	"os"
	"strings"
)

const FIRST = int('A')
const LAST = int('z')
const NUM_LETTERS = LAST - FIRST + 1

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	buff := make([]byte, 1)
	_, err := r.r.Read(buff)
	if err == io.EOF {
		return 0, io.EOF
	}

	c := int(buff[0])
	c -= FIRST
	c = (c + 13) % NUM_LETTERS
	c += FIRST
	b[0] = byte(c)
	return 1, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
