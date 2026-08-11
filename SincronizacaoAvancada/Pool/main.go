package main

import (
	"bytes"
	"fmt"
	"sync"
)

var pool = sync.Pool{
	New: func() any {
		fmt.Println("Criando novo buffer...")
		return new(bytes.Buffer)
	},
}

func main() {

	buffer := pool.Get().(*bytes.Buffer)
	buffer.WriteString("Escrevendo no buffer!")
	fmt.Println(buffer.String())

	buffer.Reset()
	pool.Put(buffer)

	buffer2 := pool.Get().(*bytes.Buffer)
	buffer2.WriteString("Escrevendo no buffer 2!")
	fmt.Println(buffer2.String())
}
