package encryption_test

import (
	"testing"

	"github.com/chenyang-zz/go-learn/basic/encryption"
)

func TestECC(t *testing.T) {
	prvKey, err := encryption.GenPrivateKey()
	if err != nil {
		t.Fatalf("genPrivateKey fail: %s\n", err)
	}
	pubKey := prvKey.PublicKey
	plain := "因为我们没有什么不同"
	cipher, err := encryption.ECCEncrypt(plain, pubKey)
	if err != nil {
		t.Fatalf("ECCEncrypt fail: %s\n", err)
	}
	plain, err = encryption.ECCDecrypt(cipher, prvKey)
	if err != nil {
		t.Fatalf("ECCDecrypt fail: %s\n", err)
	}
	t.Logf("明文：%s\n", plain)
}
