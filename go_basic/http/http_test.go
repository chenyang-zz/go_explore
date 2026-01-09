package http_test

import (
	"testing"

	"github.com/chenyang-zz/go-learn/basic/http"
)

func TestHttpServer(t *testing.T) {
	http.StartServer()
}

func TestHttpServerStream(t *testing.T) {
	http.StartHttpServerStream()
}

func TestHttpClient(t *testing.T) {
	// http.HttpObservationClient()
	// http.GetClient()
	http.HugeBodyClient()
}
