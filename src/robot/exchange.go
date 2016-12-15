package bitbot

import (
	//"time"
)

type ExchangeType int

const (
	EXCH_BTCC = 1 + iota
	EXCH_OKCOIN_CHINA
	EXCH_OKCOIN_INTERNATIONAL
)

//IBotExchange : Enforces standard functions for all exchanges supported in gocryptotrader
type IBotExchange interface {
	Setup(exconfig ExchangeAPI)
	Start()
	SetDefaults()
	GetName() string
	//GetLabel() string
	/*
						GetUSDCNY
						GetRate
						SetRate
						GetCurrency
						GetTicker
						GetDepth
						GetTrades
						GetRecords(Period)
						Buy(Price, Amount)
						Sell(Price, Amount)
						GetOrders
						GetOrder(orderId)
						CancelOrder(orderId)
						GetMinStock
						GetMinPrice
						GetFee
						GetRawJSON

			//Futures:
			Position 结构
			{
			MarginLevel	:杆杠大小, 796期货有可能为5, 10, 20三个参数, OKCoin为10或者20, BitVC期货和OK期货的全仓模式返回为固定的10, 因为原生API不支持
			Amount	:持仓量, 796期货表示持币的数量, BitVC指持仓的总金额(100的倍数), OKCoin表示合约的份数(整数且大于1)
			CanCover	:可平量, 只有股票有此选项, 表示可以平仓的数量(股票为T+1)今日仓不能平
			FrozenAmount	:冻结量, 支持传统期货与股票, 数字货币只支持796交易所
			Price	:持仓均价
			Profit	:持仓浮动盈亏(数据货币单位：BTC/LTC, 传统期货单位:RMB, 股票不支持此字段, 注: OKCoin期货全仓情况下指实现盈余, 并非持仓盈亏, 逐仓下指持仓盈亏)
			Type	:PD_LONG为多头仓位(CTP中用closebuy_today平仓), PD_SHORT为空头仓位(CTP用closesell_today)平仓, (CTP期货中)PD_LONG_YD为咋日多头仓位(用closebuy平), PD_SHORT_YD为咋日空头仓位(用closesell平)
			ContractType	:商品期货为合约代码, 股票为'交易所代码_股票代码', 具体参数SetContractType的传入类型
			}

			GetPosition	 获取当前持仓信息
			返回一个Position数组, (BitVC和OKCoin)可以传入一个参数, 指定要获取的合约类型
			SetMarginLevel(MarginLevel)	设置杆杠大小
			SetDirection(Direction)	设置Buy或者Sell下单类型
			Direction可以取buy, closebuy, sell, closesell四个参数, 传统期货多出closebuy_today,与closesell_today, 指平今仓, 默认为closebuy/closesell为平咋仓
			对于CTP传统期货, 可以设置第二个参数"1"或者"2"或者"3", 分别指"投机", "套利", "套保", 不设置默认为投机
			exchange.SetMarginLevel(5);
			exchange.SetDirection("buy");
			exchange.Buy(1000, 2);
			exchange.SetMarginLevel(5);
			exchange.SetDirection("closebuy");
			exchange.Sell(1000, 2);

			SetContractType(ContractType)	设置合约类型
			数字货币796支持: "week", "weekcny", 默认为子账户A, 要指定子账户是A还是B, 在合约后加"@A"或"@B", 如: "day@A" 为日合约A子账户
			BitVC有week和quarter和next_week三个可选参数, OKCoin期货有this_week, next_week, quarter三个参数


		期货交易中Buy, Sell, CancelOrder和现货交易的区别
		Buy或Sell之前需要调用SetMarginLevel和SetDirection明确操作类型
		数字货币796的CancelOrder之前需要调用SetDirection明确订单类型
		如: exchange.SetDirection("sell"); exchange.Sell(1000, 2);
		如: exchange.SetDirection("buy"); exchange.CancelOrder(123);

		TA - 优化重写部分常用指标函数库, talib - http://ta-lib.org/


	*/
	IsEnabled() bool
	GetTickerPrice(currency string) TickerPrice
	GetEnabledCurrencies() []string
	GetExchangeAccountInfo() (ExchangeAccountInfo, error)
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
