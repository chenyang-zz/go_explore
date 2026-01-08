package io

import (
	"io"
	"log"
	"os"
	"path"
	"strconv"
)

func SplitFile(inFile string, outDir string, n int) {
	fin, err := os.Open(inFile)
	if err != nil {
		log.Panic(err)
	}
	defer fin.Close()

	stat, err := fin.Stat()
	if err != nil {
		log.Panic(err)
	}

	fileSize := stat.Size()
	chunk := fileSize / int64(n)
	if chunk <= 0 {
		panic("file is too small or n is too large")
	}

	for i := 0; i < n; i++ {
		fout, err := os.OpenFile(path.Join(outDir, strconv.Itoa(i)+"_"+path.Base(inFile)), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}

		need := int(chunk)
		if i == n-1 {
			need = int(fileSize) - (n-1)*int(chunk)
		}
		buffer := make([]byte, need)
		_, err = fin.Read(buffer)
		if err != nil {
			log.Panic(err)
		}

		_, err = fout.Write(buffer)
		if err != nil {
			log.Panic(err)
		}

		fout.Close()
	}
}

func MergeFile(dir string, mergedFile string) {
	fout, err := os.OpenFile(mergedFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	defer fout.Close()

	if fileInfos, err := os.ReadDir(dir); err != nil {
		log.Panic(err)
	} else {
		for _, fileInfo := range fileInfos {
			if fileInfo.Type().IsRegular() {
				infile := path.Join(dir, fileInfo.Name())
				AppendFile(fout, infile)
			}
		}
	}

}

func AppendFile(fout *os.File, inFile string) {
	fin, err := os.Open(inFile)
	if err != nil {
		log.Panic(err)
	}
	defer fin.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := fin.Read(buffer)
		if err != nil {
			if err == io.EOF {
				if n > 0 {
					fout.Write(buffer[:n])
				}

			} else {
				log.Println(err)
			}
			break
		} else {
			fout.Write(buffer[:n])
		}
	}
}
