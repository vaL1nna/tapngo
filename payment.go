package tapngo

import "crypto/rsa"

type Payment struct {
	appId      string
	apiKey     []byte
	pubKeyFile string
	pubKey     *rsa.PublicKey
}
