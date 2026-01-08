package encryption_test

import (
	"bytes"
	"testing"

	"github.com/chenyang-zz/go-learn/basic/encryption"
)

func TestMyEncryption(t *testing.T) {
	key := [8]byte{34, 65, 12, 125, 65, 70, 54, 27}

	algo := encryption.NewMyEncryption(key, encryption.None)
	plain := []byte("明月多情应笑我")
	cypher := algo.Encrypt(plain)
	t.Log(cypher)
	plain2, err := algo.Decrypt(cypher)
	if err != nil {
		t.Error(err)
	} else {
		if !bytes.Equal(plain, plain2) { // 比较两个byte切片里的元素是否完全相等
			t.Log(len(plain2), string(plain2))
			t.Fail()
		}
	}

	algo = encryption.NewMyEncryption(key, encryption.CBC)
	cypher = algo.Encrypt(plain)
	t.Log(cypher)
	plain2, err = algo.Decrypt(cypher)
	if err != nil {
		t.Error(err)
	} else {
		if !bytes.Equal(plain, plain2) { // 比较两个byte切片里的元素是否完全相等
			t.Log(len(plain2), string(plain2))
			t.Fail()
		}
	}
}
