package tapngo

import "crypto/rsa"

type Refund struct {
	AppId      string
	ApiKey     []byte
	PubKeyFile string
	PubKey     *rsa.PublicKey
}
