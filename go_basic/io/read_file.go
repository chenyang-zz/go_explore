package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile() {
	if fin, err := os.Open("../data/verse.txt"); err != nil {
		fmt.Printf("open file failed: %s", err)
	} else {
		defer fin.Close()
		buf := make([]byte, 1024)
		fin.Read(buf)
		fmt.Println(string(buf))

		fin.Seek(0, 0)
		fin.Read(buf)
		fmt.Println(string(buf))

		fin.Seek(0, 0)
		const BATCH = 10
		buffer := make([]byte, BATCH)
		for {
			n, err := fin.Read(buffer)
			if n > 0 {
				fmt.Println(buffer[:n])
			}
			if err == io.EOF {
				break
			}
			// fin.Seek(0, 1)
		}
	}
}

func ReadFileWithBuffer() {
	if fin, err := os.Open("../data/verse.txt"); err != nil {
		fmt.Printf("open file failed: %s", err)
	} else {
		defer fin.Close()
		reader := bufio.NewReader(fin)
		for {
			line, err := reader.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\n")
				fmt.Println(line)
			}
			if err == io.EOF {
				break
			}
		}
	}
}
