package tapngo

import "crypto/rsa"

type Refund struct {
	appId      string
	apiKey     []byte
	pubKeyFile string
	pubKey     *rsa.PublicKey
}
