package io

import "os"

const (
	logText = "黄梅时节家家雨，青草池塘处处蛙。有约不来过夜半，闲敲棋子落灯花。\n"
)

type BufferedFileWriter struct {
	fout           *os.File
	buffer         []byte
	bufferEndIndex int
}

func NewBufferedFileWriter(fout *os.File, bufferSize int) *BufferedFileWriter {
	return &BufferedFileWriter{
		fout:           fout,
		buffer:         make([]byte, bufferSize),
		bufferEndIndex: 0,
	}
}

func (w *BufferedFileWriter) Flush() {
	w.fout.Write(w.buffer[0:w.bufferEndIndex])
	w.bufferEndIndex = 0
}

func (w *BufferedFileWriter) Write(cont []byte) {
	if len(cont) >= len(w.buffer) {
		w.Flush()
		w.fout.Write(cont)
	} else {
		if len(cont)+w.bufferEndIndex > len(w.buffer) {
			w.Flush()
		}
		copy(w.buffer[w.bufferEndIndex:], cont)
		w.bufferEndIndex += len(cont)
	}
}

func (w *BufferedFileWriter) WriteString(cont string) {
	w.Write([]byte(cont))
}

// 直接写文件
func WriteDirect(outFile string) {
	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	for i := 0; i < 100000; i++ {
		fout.WriteString(logText)
	}
}

// 带缓冲写文件
func WriteWithBuffer(outFile string) {
	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	writer := NewBufferedFileWriter(fout, 4096)
	defer writer.Flush() //最后，务必把缓冲里残留的内容写入磁盘
	for i := 0; i < 100000; i++ {
		writer.WriteString(logText)
	}
}
