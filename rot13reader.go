package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (this rot13Reader) Read(p []byte) (n int, err error) {
	n, err = this.r.Read(p)
	if err != nil {
		return n, err
	}
  
  Rot13 := func(l, base byte) byte {
    return base + (l-base+13)%26
  }

	for i := 0; i < n; i++ {
		switch {
      case p[i] >= 'A' && p[i] <= 'Z':
      p[i] = Rot13(p[i], 'A')
      
      case p[i] >= 'a' && p[i] <= 'z':
      p[i] = Rot13(p[i], 'a')
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
