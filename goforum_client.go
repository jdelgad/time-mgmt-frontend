// +build main2

package main

import (
	"github.com/jdelgad/goforum/transport"
)

func main() {
	var s transport.Socketer
	s = transport.NewClientSocket()
	s.Open()
	s.Connect("tcp://127.0.0.1:4000")

	s.Send([]byte("request"))

	s.Close()
}