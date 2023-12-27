package tapngo

import (
	"testing"
)

func TestPay(t *testing.T) {
	payment := &Payment{
		appId:  "4265838497",
		apiKey: []byte("jmaKZmPsZOhkbPpJzzIT5siKhJC5EmyTbPhj5W1ZJoSeoUV6JTqFKAuAhCfWwRgamAf30M6EFbe70MFjE1RJeQ=="),
		// pubKeyFile: "payment.pem",
	}

	txnHist := &TxnHist{
		appId:  "4265838497",
		apiKey: []byte("jmaKZmPsZOhkbPpJzzIT5siKhJC5EmyTbPhj5W1ZJoSeoUV6JTqFKAuAhCfWwRgamAf30M6EFbe70MFjE1RJeQ=="),
		// pubKeyFile: "txnHist.pem",
	}

	refund := &Refund{
		appId:  "4265838497",
		apiKey: []byte("jmaKZmPsZOhkbPpJzzIT5siKhJC5EmyTbPhj5W1ZJoSeoUV6JTqFKAuAhCfWwRgamAf30M6EFbe70MFjE1RJeQ=="),
		// pubKeyFile: "refund.pem",
	}

	_, err := New(false, WithMerchantId("42062971"), WithPayment(payment), WithTxnHist(txnHist), WithRefund(refund))

	if err != nil {
		t.Fatal(err)
	}
}
