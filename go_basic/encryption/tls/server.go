package tls

import (
	"log"
	"net"
	"time"

	"github.com/chenyang-zz/go-learn/basic/encryption"
	transport "github.com/chenyang-zz/go-learn/basic/socket"
)

func init() {
	encryption.ReadRSAKey("../../data/rsa_public_key.pem", "../../data/rsa_private_key.pem")
}

func StartTLSServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:5678")
	transport.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	transport.CheckError(err)
	log.Println("waiting for client connection ......")
	conn, err := listener.Accept()
	transport.CheckError(err)
	log.Printf("establish connection to client %s\n", conn.RemoteAddr().String())
	conn.SetReadDeadline(time.Now().Add(60 * time.Second)) //设置一个读的期限，超过这个期限再调Read()就会发生error。默认是60s内可Read()。
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer) //读取RSA加密之后的客户端的AES key
	transport.CheckError(err)
	aesKey, err := encryption.RsaDecrypt(buffer[:n]) //RSA解密
	transport.CheckError(err)
	conn.Write([]byte("I receive aes key"))

	key := [16]byte{}
	if len(aesKey) != 16 {
		panic(len(aesKey))
	}
	for i := 0; i < 16; i++ {
		key[i] = aesKey[i]
	}

	n, err = conn.Read(buffer) //读取AES加密之后的机密数据
	transport.CheckError(err)
	plain, err := encryption.AesDecrypt(string(buffer[:n]), key)
	transport.CheckError(err)
	log.Println(plain)
}
