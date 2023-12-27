package tapngo

import "crypto/rsa"

type TxnHist struct {
	AppId      string
	ApiKey     []byte
	PubKeyFile string
	PubKey     *rsa.PublicKey
}
