package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	checkErr(err)
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(bs)
	checkErr(err)
	log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString("Banik vole pyčo!")
	checkErr(err)

	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	bufferedWriter.Flush()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
