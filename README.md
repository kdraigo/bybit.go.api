# bybit-go-api
[![GO 1.21.0](https://img.shields.io/badge/Go-1.21.0-brightgreen.svg)](https://github.com/VictorFrWu/bybit-go-api)   [![Contributor Victor](https://img.shields.io/badge/contributor-Victor-blue.svg)](https://github.com/wuhewuhe/bybit-go-api)   [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/VictorFrWu/bybit-go-api/blob/main/LICENSE)
## Table of Contents
- [About](#about)
- [Release](#release)
- [Development](#development)
- [Installation](#installation)
- [Usage](#usage)
- [Contact](#contact)
- [Contributors](#contributors)
## About
The Official Go Lang API connector for Bybit's HTTP and WebSocket APIs.

Dive into a plethora of functionalities:
- Market Data Retrieval
- Trade Execution
- Position Management
- Account and Asset Info Retrieval
- User and Upgrade Management
- Public Websocket Streaming
- Private Websocket Streaming
- Institution Loan 
- Broker

bybit-go-api provides an official, robust, and high-performance go connector to Bybit's trading APIs.

Initially conceptualized by go developer Victor, this module is now maintained by Bybit's in-house go experts.

Your contributions are most welcome!

## Release
- refactor all the project
- Market endpoints:
  - add server response adapter
- Position endpoints:
  - SetPositionRiskLimit is deprecated;
  - SetPositionTpslMode is deprecated.
  - Add move position and history
- trade
  - Rename v5/execution/list to trade history
- account
  - Add transaction log to classical user
  - Get DCP Info
  - New endpoint /v5/account/smp-group
- demo trading
  - Add request coin endpoint
- asset
  - Add Get Convert Coin List
  - Add Request a Quote
  - Add Confirm a Quote
  - Add Get Convert Status
  - Add Get Convert history
- user
  - Query unlimited sub members 
- spot margin uta
   - GetSpotMarginCoin is deprecated.
   - GetSpotMarginBorrowCoin is deprecated.
   - GetSpotMarginLoanAccountInfo is deprecated.
   - GetSpotMarginBorrowOrders is deprecated.
   - GetSpotMarginRepaymentOrders is deprecated.
   - BorrowSpotMarginLoan is deprecated.
   - RepaySpotMarginLoan is deprecated.
   - Add spot margin uta interest history
- ins
   - GetC2cLendingAccountInfo is deprecated.
   - GetC2cLendingOrders is deprecated.
   - GetC2cLendingCoinInfo is deprecated.
   - C2cCancelRedeemFunds is deprecated.
   - C2cRedeemFunds is deprecated.
   - C2cDepositFunds is deprecated.
   - Add associate ins loan id
- broker
  - Add Get Sub Account Deposit Records

## Development
bybit-go-api is under active development with the latest features and updates from Bybit's API implemented promptly. The module utilizes minimal external libraries to provide a lightweight and efficient experience. If you've made enhancements or fixed bugs, please submit a pull request.

## Installation
Ensure you have go 1.21.0 or higher. And use dependencies as below
```go
require (
	github.com/google/uuid v1.4.0
	github.com/gorilla/websocket v1.5.1
	github.com/stretchr/testify v1.8.4
)
```

To import my package you need just to put the link to your go mode file
**github.com/kdraigo/bybit.go.api**

## Usage
Note: Replace placeholders (like YOUR_API_KEY, links, or other details) with the actual information. You can also customize this template to better fit the actual state and details of your Java API.
### Rest API
- Place an order by Map
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "side": "Buy", "positionIdx": 0, "orderType": "Limit", "qty": "0.001", "price": "10000", "timeInForce": "GTC"}
orderResult, err := client.NewUtaBybitServiceWithParams(params).PlaceOrder(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(orderResult))
```

- Place an order by Trade Class
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
orderResult, err := client.NewPlaceOrderService("linear", "XRPUSDT", "Buy", "Market", "10").Do(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(orderResult))
```

- Place batch order
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
params := map[string]interface{}{"category": "option",
	"request": []map[string]interface{}{
		{
			"category":    "option",
			"symbol":      "BTC-10FEB23-24000-C",
			"orderType":   "Limit",
			"side":        "Buy",
			"qty":         "0.1",
			"price":       "5",
			"orderIv":     "0.1",
			"timeInForce": "GTC",
			"orderLinkId": "9b381bb1-401",
			"mmp":         false,
			"reduceOnly":  false,
		},
		{
			"category":    "option",
			"symbol":      "BTC-10FEB23-24000-C",
			"orderType":   "Limit",
			"side":        "Buy",
			"qty":         "0.1",
			"price":       "5",
			"orderIv":     "0.1",
			"timeInForce": "GTC",
			"orderLinkId": "82ee86dd-001",
			"mmp":         false,
			"reduceOnly":  false,
		},
	},
}
orderResult, err := client.NewUtaBybitServiceWithParams(params).PlaceBatchOrder(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(orderResult))
```

- Get Position 
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
params := map[string]interface{}{"category": "linear", "settleCoin": "USDT", "limit": 10}
orderResult, err := client.NewUtaBybitServiceWithParams(params).GetPositionList(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(orderResult))
```

- Get Transaction Log
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
params := map[string]interface{}{"accountType": "UNIFIED", "category": "linear"}
accountResult, err := client.NewUtaBybitServiceWithParams(params).GetTransactionLog(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(accountResult))
```

### Websocket public channel
- Order book Subscribe
```go
ws := bybit.NewBybitPublicWebSocket("wss://stream.bybit.com/v5/public/spot", func(message string) error {
fmt.Println("Received:", message)
return nil
})
_ = ws.Connect([]string{"orderbook.1.BTCUSDT"})
select {}
```

### Websocket private channel
```go
ws := bybit.NewBybitPrivateWebSocket("wss://stream-testnet.bybit.com/v5/private", "YOUR_API_KEY", "YOUR_API_SECRET", func(message string) error {
	fmt.Println("Received:", message)
	return nil
})
_ = ws.Connect([]string{"order"})
select {}
```

## Contact
For support, join our Bybit API community on [Telegram](https://t.me/Bybitapi).

## Contributors
List of other contributors
<table>
  <tr>
    <td align="center">
        <a href="https://github.com/VictorFrWu">
            <img src="https://avatars.githubusercontent.com/u/32245754?v=4" width="100px;" alt=""/>
            <br />
            <sub>   
                <b>Victor</b>
            </sub>
        </a>
        <br />
        <a href="https://github.com/VictorFrWu/bybit.go.api/commits?author=wuhewuhe" title="Code">💻</a>
        <a href="https://github.com/VictorFrWu/bybit.go.api/commits?author=wuhewuhe" title="Documentation">📖</a>
    </td>
  </tr>
</table>