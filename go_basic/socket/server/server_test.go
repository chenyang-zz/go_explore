package server_test

import (
	"testing"

	"github.com/chenyang-zz/go-learn/basic/socket/server"
)

func TestTcpServer(t *testing.T) {
	server.TcpServer()
}

func TestUdpServer(t *testing.T) {
	server.UdpServer()
}

func TestTcpLongConnection(t *testing.T) {
	server.TcpLongConnection()
}

func TestUdpLongConnection(t *testing.T) {
	server.UdpLongConnection()
}

func TestTcpStick(t *testing.T) {
	server.TcpStick()
}

func TestUdpRpcServer(t *testing.T) {
	server.UdpRpcServer()
}
