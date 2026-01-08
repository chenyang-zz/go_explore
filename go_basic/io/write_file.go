package io

import (
	"bufio"
	"fmt"
	"os"
)

func WriteFile() {
	if fout, err := os.OpenFile("../data/verse.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o666); err != nil {
		fmt.Printf("open file failed: %s", err)
	} else {
		defer fout.Close()
		fout.WriteString("的兰溪的\n")
		fout.WriteString("发了多少\n")
		fout.WriteString("发了多少发了多少")
	}
}

func WriteFileWithBuffer() {
	if fout, err := os.OpenFile("../data/verse.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o666); err != nil {
		fmt.Printf("open file failed: %s", err)
	} else {
		defer fout.Close()
		writer := bufio.NewWriter(fout)
		writer.WriteString("的兰溪的\n")
		writer.WriteString("发了多少\n")
		writer.WriteString("发了多少发了多少")
		writer.Flush()
	}
}
