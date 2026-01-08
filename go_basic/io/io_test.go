package io_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/chenyang-zz/go-learn/basic/io"
)

func TestWriteFile(t *testing.T) {
	io.WriteFile()
}

func TestReadFile(t *testing.T) {
	io.ReadFile()
}

func TestWriteFileWithBuffer(t *testing.T) {
	io.WriteFileWithBuffer()
}

func TestReadFileWithBuffer(t *testing.T) {
	io.ReadFileWithBuffer()
}

func TestBufferedFileWriter(t *testing.T) {
	t1 := time.Now()
	io.WriteDirect("../data/no_buffer.txt")
	t2 := time.Now()
	io.WriteWithBuffer("../data/with_buffer.txt")
	t3 := time.Now()
	fmt.Printf("不用缓冲耗时%dms，用缓冲耗时%dms\n", t2.Sub(t1).Milliseconds(), t3.Sub(t2).Milliseconds())
}

func TestCreateFIle(t *testing.T) {
	io.CreateFile("../data/poem.txt")
}

func TestWalkDir(t *testing.T) {
	io.WalkDir("../data")
}

func TestSplitFile(t *testing.T) {
	imgFile := "../img/logo.png"
	io.SplitFile(imgFile, "../img/图像分割", 4)
}

func TestMergeFile(t *testing.T) {
	io.MergeFile("../img/图像分割", "../img/图像合并.png")
}

func TestLimitReader(t *testing.T) {
	io.LimitReader()
}

func TestMultiReader(t *testing.T) {
	io.MultiReader()
}

func TestMultiWriter(t *testing.T) {
	io.MultiWriter()
}

func TestTeeReader(t *testing.T) {
	io.TeeReader()
}

func TestPipe(t *testing.T) {
	io.PipeIO()
}

func TestCompress(t *testing.T) {
	io.Compress("../img/logo.png", "../img/logo.png.gzip", io.GZIP)
	io.Decompress("../img/logo.png.gzip", "../data/logo.png", io.GZIP)
}

func TestJson(t *testing.T) {
	io.JsonSerialize()
}

func TestLog(t *testing.T) {
	logger := io.NewLogger("../data/biz.log")
	io.Log(logger)
}

func TestSLog(t *testing.T) {
	logger := io.NewSLogger("../data/sbiz.log")
	io.SLog(logger)
}

func TestSysCall(t *testing.T) {
	io.SysCall()
}

func TestRegex(t *testing.T) {
	io.UseRegex()
}
