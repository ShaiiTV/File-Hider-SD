package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("FileConfig.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fmt.Println("Read", n, "bytes:", string(data[:n]))
	}
}
