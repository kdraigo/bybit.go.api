package main

import (
	"context"
	"fmt"

	bybit "github.com/kdraigo/bybit.go.api"
)

func GetMarkPriceKline() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarkPriceKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
