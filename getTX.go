package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/rpcnode"
)

// GetTX teste
func GetTX() {
	ctx := context.Background()

	config := &rpcnode.Config{
		Host:     os.Getenv("RPC_HOST"),
		Username: os.Getenv("RPC_USERNAME"),
		Password: os.Getenv("RPC_PASSWORD"),
	}
	fmt.Printf("Connect to %s as %s password : %s\n", config.Host, config.Username, config.Password)

	node, err := rpcnode.NewNode(config)
	if err != nil {
		fmt.Printf("Failed to create node : %s\n", err.Error())
	}

	txid, err := bitcoin.NewHash32FromStr(os.Getenv("TX_ID"))
	fmt.Printf("Get Tx : %s\n", txid.String())

	if tx, err := node.GetTX(ctx, txid); err != nil {
		fmt.Printf("Failed to get tx : %s\n", err.Error())
	} else {
		fmt.Printf("Tx : %s\n%+v\n", tx.TxHash().String(), tx)
	}
}
