package tls_test

import (
	"testing"
	"time"

	"github.com/chenyang-zz/go-learn/basic/encryption/tls"
)

func TestTLS(t *testing.T) {
	go tls.StartTLSServer()
	time.Sleep(time.Second)
	tls.StartTLSClient()
}
