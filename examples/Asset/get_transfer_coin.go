package main

import (
	"context"
	"fmt"

	bybit "github.com/kdraigo/bybit.go.api"
)

func PlaceOrder() {
	client := bybit.NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"fromAccountType": "UNIFIED", "toAccountType": "CONTRACT"}
	response, err := client.NewUtaBybitServiceWithParams(params).GetTransferableCoin(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(response))
}
