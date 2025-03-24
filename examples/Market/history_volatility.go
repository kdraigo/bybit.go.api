package main

import (
	"context"
	"fmt"

	bybit "github.com/kdraigo/bybit.go.api"
)

func GetHistoryVolatility() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET), bybit.WithDebug(true))
	params := map[string]interface{}{"category": "option", "baseCoin": "BTC"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetHistoryVolatility(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
