package main

import (
	"time"
)

//IExchange : Enforces standard functions for all exchanges supported
type IExchange interface {

	//######################spot exchange interface######################
	GetName() string
	GetLabel() string

	//GetUSDCNY()float64
	//GetRate() float64
	//SetRate(rate float64)
	// GetCurrency() string
	//GetTicker() Ticker
	// GetDepth() []Depth
	// GetTrades() []Trade
	// GetRecords(KLineIntervalType) []Record
	GetAccount() (ExchangeAccountInfo, error)
	// LimitBuy(price , amount float64)
	// LimitSell(price , amount float64)
	// MarketBuy( price float64)
	// MarketSell( amount float64)
	// GetOrders() []Order
	// GetOrder(orderID string) Order
	// CancelOrder(orderID string) bool

	//######################future exchange interface######################
	//返回一个Position数组, (BitVC和OKCoin)可以传入一个参数, 指定要获取的合约类型
	//GetPosition() //	 获取当前持仓信息

	//SetMarginLevel(MarginLevel) //设置杆杠大小

	//Direction可以取buy, closebuy, sell, closesell四个参数, 传统期货多出closebuy_today,与closesell_today, 指平今仓, 默认为closebuy/closesell为平咋仓
	//SetDirection(Direction)	//设置Buy或者Sell下单类型

	//数字货币796支持: "week", "weekcny", 默认为子账户A, 要指定子账户是A还是B, 在合约后加"@A"或"@B", 如: "day@A" 为日合约A子账户
	//BitVC有week和quarter和next_week三个可选参数, OKCoin期货有this_week, next_week, quarter三个参数
	//SetContractType(ContractType)	设置合约类型

	//期货交易中Buy, Sell, CancelOrder和现货交易的区别
	//Buy或Sell之前需要调用SetMarginLevel和SetDirection明确操作类型
	//数字货币796的CancelOrder之前需要调用SetDirection明确订单类型
	//如: exchange.SetDirection("sell"); exchange.Sell(1000, 2);
	//如: exchange.SetDirection("buy"); exchange.CancelOrder(123);

	//GetFuturesTicker()
	//GetFuturesDepth()
	//OKCoinFuturesTrades
	//GetFuturesTrades()
	//GetFuturesIndex
	//GetFuturesExchangeRate()
	//GetFuturesEstimatedPrice()
	//GetFutureRecords()
	//GetFuturesHoldAmount()
	//GetFuturesExplosive()
	//

	//######################Exchange Scope interface######################
	// GetMinStock() float64
	// GetMinPrice() float64
	// GetFee() Fee
	//GetRawJSON

	//Support: GetTicker, GetDepth, GetTrades, GetRecords, GetAccount, GetOrders, GetOrder, CancelOrder, Buy, Sell, GetPosition
	//Go(Method, Args...)

	//IO("api",apiName string, args interface{}) //this is original from botvs
	//API(apiName string, args interface{}) string

	//above functions are exported to VM

	GetEnabledCurrencies() []string
	Setup(exconfig ExchangeConfig)
	Start()
	SetDefaults()
}

type ExchangeAccount struct {
	Exchange                string
	Label                   string
	APIKey                  string
	APISecret               string
	ClientID                string
	EnabledPairs            string
	AuthenticatedAPISupport bool
}

type ExchangeConfig struct {
	Enabled          bool
	RESTPollingDelay time.Duration
	AvailablePairs   string
	BaseCurrencies   string
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

type Account struct {
	Balance       float64 //Fiat
	FrozenBalance float64
	Stocks        float64 //Crypto
	FrozenStocks  float64
}

//ExchangeAccountInfo : Generic type to hold each exchange's holdings in all enabled currencies
type ExchangeAccountInfo struct {
	ExchangeName string
	Currencies   []ExchangeAccountCurrencyInfo
}

//ExchangeAccountCurrencyInfo : Sub type to store currency name and value
type ExchangeAccountCurrencyInfo struct {
	CurrencyName string
	TotalValue   float64
	Hold         float64
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
