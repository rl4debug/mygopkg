package main

import (
	"flag"
	"fmt"
)

// Reader is an
type Reader interface {
	Read(p []byte) (n int, err error)
}

//AlertCounter is an unexported type that
//contains an integer counter for alerts.
type AlertCounter int

// Writer is and
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
	Reader
	Writer
}

//MyType This is
//my type
type MyType struct {
	Name  string /*name of type*/
	Value string // struct of type
}

func (mt MyType) f1() {

}

func test(rw ReadWriter) {
	rw.Read(nil)
}

//This is used for ...
func main() {
	mt := new(MyType)
	mt.f1()

	f1 := flag.String("flat1", "123", "this is flag 1")
	flag.Parse()
	fmt.Printf("%s", *f1)
}
