//实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。

package main

import "fmt"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader := MyReader{}
	buffer := make([]byte, 5)
	n, err := reader.Read(buffer)
	fmt.Printf("Read %d bytes: %s, error: %v\n", n, buffer, err)
}
