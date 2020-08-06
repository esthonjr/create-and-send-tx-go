package main

import (
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/txbuilder"
	"github.com/tokenized/pkg/wire"
)

// CreateTX teste
func CreateTX() {
	// Private key
	wif := "cS2nfuBWemJekA82fUioJ8neGtcsHs6WPazYLhwhb3UdBzAeyYXV"
	key, err := bitcoin.KeyFromStr(wif)
	if err != nil {
		panic(err)
	}

	// Change address
	changeAddress, _ := bitcoin.DecodeAddress("mmsio7EbL2ZwJConqaUovnSYFEc8vxnbuf")

	// Create an instance of the TxBuilder using 512 as the dust limit and 1.1 sat/byte fee rate.
	builder := txbuilder.NewTxBuilder(512, 1.1)
	builder.SetChangeAddress(bitcoin.NewRawAddressFromAddress(changeAddress), "")

	// Add an input: txid, output index, value, locking script, spend address
	hash, _ := bitcoin.NewHash32FromStr("699eba6bd9f9fd6a82dcb4ae69d756003a6b1356e4d2a2eefa5d8ca916b71588")
	index := uint32(0)
	value := uint64(5000000000)
	spendAddress, _ := bitcoin.DecodeAddress("mikAzBEzQtbVbgAFEv4n1cEkDGXJi9mfhd")
	lockingScript, _ := bitcoin.NewRawAddressFromAddress(spendAddress).LockingScript()
	_ = builder.AddInput(*wire.NewOutPoint(hash, index), lockingScript, value)

	// add an output to the recipient (1000 satoshis)
	paymentAddress, _ := bitcoin.DecodeAddress("mmdaHdaiSqzu9AL1DjnBR7NgUcBKZMzwGZ")
	_ = builder.AddPaymentOutput(bitcoin.NewRawAddressFromAddress(paymentAddress), 1000, false)

	// sign the first and only input
	_ = builder.Sign([]bitcoin.Key{key})
	fmt.Println(builder)
	// get the raw TX bytes
	rawTX, _ := builder.Serialize()
	fmt.Println(rawTX)
}
