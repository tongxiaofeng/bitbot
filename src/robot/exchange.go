package main

import (
	"time"
)

type ExchangeConfig struct {
	RESTPollingDelay time.Duration
	AvailablePairs   string
	BaseCurrencies   string
}

//IExchange : Enforces standard functions for all exchanges supported
type IExchange interface {
	GetAvailablePairs() []string
	GetName() string
	Setup(exconfig ExchangeConfig)
	Start()
}

type Fee struct {
	Sell float64 //percentage 0.2 for 0.2%
	Buy  float64 //percentage
}

type MarketOrder struct {
	Price  float64
	Amount float64
}

type KLineIntervalType int

const (
	KLINE_1_MIN KLineIntervalType = iota
	KLINE_5_MIN
	KLINE_15_MIN
	KLINE_30_MIN
	KLINE_1_HOUR
	KLINE_1_DAY
)

type Depth struct {
	Asks []MarketOrder
	Bids []MarketOrder
}

type OrderType int

const (
	ORDER_TYPE_BUY = iota
	ORDER_TYPE_SELL
)

type Trade struct {
	Time   time.Time
	Price  float64
	Amount float64
	Type   OrderType
}

type Record struct {
	Time   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type Order struct {
	ID         string
	Price      float64
	Amount     float64
	DealAmount float64
	Status     OrderStatus
	Type       OrderType
}

type OrderStatus int

const (
	ORDER_STATE_PENDING = iota
	ORDER_STATE_CLOSED
	ORDER_STATE_CANCELED
)

type Ticker struct {
	High   float64
	Low    float64
	Sell   float64
	Buy    float64
	Last   float64
	Volume float64
}

type ExchangeAccountInfo struct {
	Balance       float64 //Fiat
	FrozenBalance float64
	Stocks        float64 //Crypto
	FrozenStocks  float64
}

type PositionType int

const (
	PD_LONG PositionType = iota
	PD_SHORT
)

type Position struct {
	MarginLevel int
	Amount      float64
	//CanCover  //Only for stock
	FrozenAmount float64
	Price        float64
	Profit       float64
	Type         PositionType
	ContractType string
}
