package tapngo

import "crypto/rsa"

type TxnHist struct {
	appId      string
	apiKey     []byte
	pubKeyFile string
	pubKey     *rsa.PublicKey
}
