package encryption_test

import (
	"testing"
	"time"

	"github.com/chenyang-zz/go-learn/basic/encryption"
)

var (
	BigFile = "/Users/sheepzhao/Downloads/node-v22.5.1-linux-x64.tar.xz"
)

func TestHash(t *testing.T) {
	data := "123456"
	hs := encryption.Sha1(data)
	t.Log("SHA-1", hs, len(hs))
	hm := encryption.Md5(data)
	t.Log("MD5", hm, len(hm))
}

func TestCreateSha256OfSmallFile(t *testing.T) {
	begin := time.Now()
	hash, err := encryption.CreateSha256OfSmallFile(BigFile)
	if err != nil {
		t.Error(err)
	}
	t.Log("CreateSha256OfSmallFile", hash, "use time", time.Since(begin).Milliseconds())
}

func TestCreateSha256OfBigFile(t *testing.T) {
	begin := time.Now()
	hash, err := encryption.CreateSha256OfBigFile(BigFile, 10<<20)
	if err != nil {
		t.Error(err)
	}
	t.Log("CreateSha256OfBigFile", hash, "use time", time.Since(begin).Milliseconds())
}
