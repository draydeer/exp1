package cryptors

type Cryptor interface {
	Decrypt(value string, decryptionKey string) string
}

type CryptorInstance struct {

}
