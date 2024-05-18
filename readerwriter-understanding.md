```go

package main

import (
	"compress/gzip"
	"fmt"
)


type Reader interface {
	Read([]byte) (int, error)
}
type Writer interface {
	Write([]byte) (int, error)
}

type WriterReader interface {
	Reader
	Writer
}

type myBuff struct {
	buff []byte
}

func NewBuff() *myBuff {
	buff := make([]byte, 0)
	return &myBuff{
		buff: buff,
	}
}

func (w *myBuff) Write(p []byte) (int, error) {
	w.buff = append(w.buff, p...)
	return len(p), nil
}

func (w *myBuff) Read(p []byte) (int, error) {
	copy(p, w.buff)
	return len(p), nil
}



func main(){

	buff := new(myBuff)
	wr := gzip.NewWriter(buff)
	wr.Write([]byte("Hello, Gophers!"))
	wr.Close()
	fmt.Println(buff.buff)
	
}
```