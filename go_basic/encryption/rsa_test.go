package encryption_test

import (
	"testing"

	"github.com/chenyang-zz/go-learn/basic/encryption"
)

func TestRSA(t *testing.T) {
	encryption.ReadRSAKey("../data/rsa_public_key.pem", "../data/rsa_private_key.pem")

	plain := "因为我们没有什么不同"
	cipher, err := encryption.RsaEncrypt([]byte(plain))
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("密文：%v\n", (cipher))
		bPlain, err := encryption.RsaDecrypt(cipher)
		if err != nil {
			t.Log(err)
		} else {
			t.Logf("明文：%s\n", string(bPlain))
		}
	}
}
