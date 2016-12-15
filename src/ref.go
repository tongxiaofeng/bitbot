package bitbot

/*
Record

{
Time	:一个时间戳, 精确到毫秒，与Javascript的 new Date().getTime() 得到的结果格式一样
Open	:开盘价
High	:最高价
Low	:最低价
Close	:收盘价
Volume	:交易量
}

MarketOrder
{
Price	:价格
Amount	:数量
}

Ticker
{
High	:最高价
Low	:最低价
Sell	:卖一价
Buy	:买一价
Last	:最后成交价
Volume	:最近成交量
}


Order

{
Id	:交易单唯一标识
Price	:下单价格
Amount	:下单数量
DealAmount	:成交数量
Status	:订单状态, 参考常量里的订单状态
Type	:订单类型, 参考常量里的订单类型
}


Depth

{
Asks	:卖单数组, MarketOrder数组, 按价格从低向高排序
Bids	:买单数组, MarketOrder数组, 按价格从高向低排序
}


Trade
{
Time	:时间(Unix timestamp 毫秒)
Price	:价格
Amount	:数量
Type	:订单类型, 参考常量里的订单类型
}

Fee
{
Sell	:卖出手续费, 为一个浮点数, 如0.2表示0.2%的手续费
Buy	:买入手续费, 格式同上
}

Account
{
Balance	:余额(人民币或者美元, 在Poloniex交易所里BTC_ETC这样的品种, Balance就指的是BTC的数量, Stocks指的是ETC数量, BTC38的ETC_BTC相当于Poloniex的BTC_ETC, 指的是以BTC计价)
FrozenBalance	:冻结的余额
Stocks	:BTC/LTC数量, 现货为当前可操作币的余额(去掉冻结的币), 期货的话为合约当前可用保证金(传统期货为此属性)
FrozenStocks	:冻结的BTC/LTC数量(传统期货无此属性)
}
*/
