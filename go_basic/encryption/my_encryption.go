package encryption

import "errors"

const (
	None = iota
	CBC
)

// XOR 异或运算，要求plain和key的长度相同
//
// 单看一个比特，任何数(0或1)跟0异或还是自己、跟1异或是其反面，所以任何数跟0两次异或还是自己、任何数跟1两次异或还是自己
type MyEncryption struct {
	Key       [8]byte
	BlockSize int
	BlockMode int
}

func NewMyEncryption(key [8]byte, blockMode int) *MyEncryption {
	return &MyEncryption{
		Key:       key,
		BlockSize: len(key),
		BlockMode: blockMode,
	}
}

func (en *MyEncryption) encryptBlock(plain, cypher []byte) {
	for i := 0; i < len(plain); i++ {
		cypher[i] = plain[i] ^ en.Key[i]
	}
}

func (en *MyEncryption) decryptBlock(plain, cypher []byte) {
	for i := 0; i < len(cypher); i++ {
		plain[i] = cypher[i] ^ en.Key[i]
	}
}

func (en *MyEncryption) confuseBlock(plain, prevCypher []byte) {
	switch en.BlockMode {
	case CBC:
		for i := 0; i < len(plain); i++ {
			plain[i] = plain[i] ^ prevCypher[i]
		}
	default:
	}
}

func (en *MyEncryption) deconfuseBlock(plain, prevCypher []byte) {
	en.confuseBlock(plain, prevCypher)
}

func (en *MyEncryption) Encrypt(plain []byte) []byte {
	plainPadding := PKCS7.Padding(plain, en.BlockSize) //明文末尾填充字节，长度成为BlockSize的整倍数
	cypher := make([]byte, len(plainPadding))
	prevCypher := make([]byte, en.BlockSize) //全0，任何数跟0异或还是自身
	for i := 0; i < len(plainPadding); i += en.BlockSize {
		begin := i
		end := i + en.BlockSize
		en.confuseBlock(plainPadding[begin:end], prevCypher)
		en.encryptBlock(plainPadding[begin:end], cypher[begin:end])
		copy(prevCypher, cypher[begin:end])
	}
	return cypher
}

func (en *MyEncryption) Decrypt(cypher []byte) ([]byte, error) {
	if len(cypher)%en.BlockSize != 0 {
		return nil, errors.New("密文长度不合法")
	}
	if len(cypher) == 0 {
		return []byte{}, nil
	}
	plainPadding := make([]byte, len(cypher))
	blockNum := len(cypher) / en.BlockSize
	// 倒着解密，先解密最后一个分组（也可以正着解密，因为跟顺序没关系，甚至可以并行解密）
	for i := blockNum - 1; i >= 0; i-- {
		begin := i * en.BlockSize
		end := begin + en.BlockSize
		var prevCypher []byte //前一组密文
		if i == 0 {
			prevCypher = make([]byte, en.BlockSize) //全0，任何数跟0异或还是自身
		} else {
			prevCypher = cypher[begin-en.BlockSize : end-en.BlockSize]
		}
		en.decryptBlock(plainPadding[begin:end], cypher[begin:end])
		en.deconfuseBlock(plainPadding[begin:end], prevCypher)
	}
	return PKCS7.Unpadding(plainPadding, en.BlockSize)
}
