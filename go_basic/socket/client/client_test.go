package client_test

import (
	"testing"

	"github.com/chenyang-zz/go-learn/basic/socket/client"
)

func TestTcpClient(t *testing.T) {
	client.TcpClient()
}

func TestUdpClient(t *testing.T) {
	client.UdpClient()
}

func TestTcpLongConnection(t *testing.T) {
	client.TcpLongConnection()
}

func TestUdpLongConnection(t *testing.T) {
	client.UdpLongConnection()
}

func TestTcpStick(t *testing.T) {
	client.TcpStick()
}

func TestUdpRpcClient(t *testing.T) {
	client.UdpRpcClient()
}
