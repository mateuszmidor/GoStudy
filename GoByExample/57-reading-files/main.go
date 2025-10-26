package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

const filename = "/proc/version"

func main() {
	// read all data
	dat, err := ioutil.ReadFile(filename)
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	// read only some bytes from beginning
	b1 := make([]byte, 6)
	n1, err := f.Read(b1)
	check(err)
	fmt.Println(string(b1[:n1])) // Linux

	// seek from current(1) position and read
	_, err = f.Seek(8, 1)
	check(err)
	b2 := make([]byte, 8)
	n2, err := f.Read(b2)
	check(err)
	fmt.Println(string(b2[:n2])) // 5.6.15-1

	_, err = f.Seek(1, 1)
	r := bufio.NewReader(f)
	b3, err := r.ReadString(' ')
	check(err)
	fmt.Println(string(b3)) // MANJARO
}
