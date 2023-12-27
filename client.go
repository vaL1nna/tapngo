package tapngo

import (
	"crypto/rsa"
	"net/http"
	"os"
	"tapngo/utils"
)

const (
	UatGateWay        = "gateway.sandbox.tapngo.com.hk"
	ProductionGateWay = "gateway2.tapngo.com.hk"
)

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
		if payment.pubKeyFile != "" {
			payment.pubKey, err = c.LoadPublicKeyFromFile(payment.pubKeyFile)
		}
		c.payment = payment
		return err
	}
}

func WithTxnHist(txnHist *TxnHist) OptionFunc {
	return func(c *Client) error {
		var err error
		if txnHist.pubKeyFile != "" {
			txnHist.pubKey, err = c.LoadPublicKeyFromFile(txnHist.pubKeyFile)
		}
		c.txnHist = txnHist
		return err
	}
}

func WithRefund(refund *Refund) OptionFunc {
	return func(c *Client) error {
		var err error
		if refund.pubKeyFile != "" {
			refund.pubKey, err = c.LoadPublicKeyFromFile(refund.pubKeyFile)
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
