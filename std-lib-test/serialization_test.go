package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGob(t *testing.T) {

	//var buf bytes.Buffer
	//enc := gob.NewEncoder(&buf)
	//enc.Encode("111")

	var buf1 bytes.Buffer
	p := [6]byte{65, 66, 67, 226, 130, 172}
	buf1.Write(p[:])
	fmt.Println(buf1.String())

}
