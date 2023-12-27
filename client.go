package tapngo

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"github.com/val1nna/tapngo/utils"
	"net/http"
	"os"
)

const (
	UatGateWay        = "gateway.sandbox.tapngo.com.hk"
	ProductionGateWay = "gateway2.tapngo.com.hk"
)

var (
	ErrNoPubKeyFile = errors.New("tapngo: no public key file")
	ErrNoApiKey     = errors.New("tapngo: no api key")
)

type Request interface {
	Params() map[string]string
}

type OptionFunc func(c *Client) error

type Client struct {
	environment string
	host        string
	merchantId  string
	payment     *Payment
	txnHist     *TxnHist
	refund      *Refund
	httpClient  *http.Client
}

func New(isProduction bool, opts ...OptionFunc) (*Client, error) {
	client := &Client{}

	if isProduction {
		client.environment = "production"
		client.host = ProductionGateWay
	} else {
		client.environment = "uat"
		client.host = UatGateWay
	}

	client.httpClient = &http.Client{}

	for _, opt := range opts {
		if opt != nil {
			err := opt(client)
			if err != nil {
				return nil, err
			}
		}
	}

	return client, nil
}

func WithMerchantId(merchantId string) OptionFunc {
	return func(c *Client) error {
		c.merchantId = merchantId
		return nil
	}
}

func WithPayment(payment *Payment) OptionFunc {
	return func(c *Client) error {
		var err error
		if payment.PubKeyFile != "" {
			payment.PubKey, err = c.LoadPublicKeyFromFile(payment.PubKeyFile)
		}
		c.payment = payment
		return err
	}
}

func WithTxnHist(txnHist *TxnHist) OptionFunc {
	return func(c *Client) error {
		var err error
		if txnHist.PubKeyFile != "" {
			txnHist.PubKey, err = c.LoadPublicKeyFromFile(txnHist.PubKeyFile)
		}
		c.txnHist = txnHist
		return err
	}
}

func WithRefund(refund *Refund) OptionFunc {
	return func(c *Client) error {
		var err error
		if refund.PubKeyFile != "" {
			refund.PubKey, err = c.LoadPublicKeyFromFile(refund.PubKeyFile)
		}
		c.refund = refund
		return err
	}
}

func (c *Client) LoadPublicKeyFromFile(file string) (*rsa.PublicKey, error) {
	key, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}

	return utils.ParseRSAPublicKeyFromPEM(key)
}

func (c *Client) Payload(plainPayload []byte, pubKey *rsa.PublicKey) (string, error) {
	var encrypted []byte
	var err error

	if pubKey == nil {
		return "", ErrNoPubKeyFile
	}

	// Encrypt the payload
	if encrypted, err = rsa.EncryptOAEP(sha1.New(), rand.Reader, pubKey, plainPayload, nil); err != nil {
		return "", err
	}

	// Return the Base64-encoded encrypted data
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (c *Client) Sign(request Request, apiKey []byte) (string, error) {
	if apiKey == nil {
		return "", ErrNoApiKey
	}
	str := utils.BuildQueryString(request.Params())

	h := hmac.New(sha512.New, apiKey)
	h.Write([]byte(str))
	hash := h.Sum(nil)

	return base64.StdEncoding.EncodeToString(hash), nil
}
