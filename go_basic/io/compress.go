package io

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

const (
	_ = iota
	GZIP
	ZLIB
)

func Compress(inFile, outFile string, compressAlgo int) {
	fin, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fin.Close()

	stat, _ := fin.Stat()
	fmt.Printf("压缩前文件大小 %dB\n", stat.Size())

	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()

	var writer io.WriteCloser
	switch compressAlgo {
	case GZIP:
		writer = gzip.NewWriter(fout)
	case ZLIB:
		writer = zlib.NewWriter(fout)
	}
	defer writer.Close()

	io.Copy(writer, fin)
}

func Decompress(inFile, outFile string, compressAlgo int) {
	fin, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fin.Close()

	stat, _ := fin.Stat()
	fmt.Printf("压缩后文件大小 %dB\n", stat.Size())

	var reader io.ReadCloser
	switch compressAlgo {
	case GZIP:
		reader, err = gzip.NewReader(fin)
	case ZLIB:
		reader, err = zlib.NewReader(fin)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reader.Close()

	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()

	io.Copy(fout, reader)
}
