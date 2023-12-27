package tapngo

import "crypto/rsa"

type Payment struct {
	AppId      string
	ApiKey     []byte
	PubKeyFile string
	PubKey     *rsa.PublicKey
}
