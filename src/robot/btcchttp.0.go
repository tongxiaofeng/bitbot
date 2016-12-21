package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//"AvailablePairs": "BTCCNY,LTCCNY,LTCBTC"
const (
	BTCC_API_URL                  = "https://api.btcc.com/"
	BTCC_API_AUTHENTICATED_METHOD = "api_trade_v1.php"
	BTCC_API_VER                  = "2.0.1.4"
	BTCC_ORDER_BUY                = "buyOrder2"
	BTCC_ORDER_SELL               = "sellOrder2"
	BTCC_ORDER_CANCEL             = "cancelOrder"
	BTCC_ICEBERG_BUY              = "buyIcebergOrder"
	BTCC_ICEBERG_SELL             = "sellIcebergOrder"
	BTCC_ICEBERG_ORDER            = "getIcebergOrder"
	BTCC_ICEBERG_ORDERS           = "getIcebergOrders"
	BTCC_ICEBERG_CANCEL           = "cancelIcebergOrder"
	BTCC_ACCOUNT_INFO             = "getAccountInfo"
	BTCC_DEPOSITS                 = "getDeposits"
	BTCC_MARKETDEPTH              = "getMarketDepth2"
	BTCC_ORDER                    = "getOrder"
	BTCC_ORDERS                   = "getOrders"
	BTCC_TRANSACTIONS             = "getTransactions"
	BTCC_WITHDRAWAL               = "getWithdrawal"
	BTCC_WITHDRAWALS              = "getWithdrawals"
	BTCC_WITHDRAWAL_REQUEST       = "requestWithdrawal"
	BTCC_STOPORDER_BUY            = "buyStopOrder"
	BTCC_STOPORDER_SELL           = "sellStopOrder"
	BTCC_STOPORDER_CANCEL         = "cancelStopOrder"
	BTCC_STOPORDER                = "getStopOrder"
	BTCC_STOPORDERS               = "getStopOrders"
	BTCC_ArchivedOrder            = "getArchivedOrder"
	BTCC_ArchivedOrderS           = "getArchivedOrders"
)

/*
 Code Message -32000 内部错误
-32003 人民币账户余额不足
-32004 比特币账户余额不足
-32005 挂单没有找到
-32006 无效的用户
-32007 无效的货币
-32008 无效的金额
-32009 无效的钱包地址
-32010 没有找到提现记录
-32011 没有找到充值记录
-32017 无效的类型
-32018 无效的价格
-32019 无效的参数
-32025 订单已取消
-32026 订单已完成
-32062 市场深度不足以成交该市场交易
-32065 无效的货币参数
-32086 订单处理中
*/

type BTCC struct {
	Name                    string
	Label                   string
	Enabled                 bool
	RESTPollingDelay        time.Duration
	AuthenticatedAPISupport bool
	APISecret, APIKey       string
	Fee                     float64
	BaseCurrencies          []string
	AvailablePairs          []string
	EnabledPairs            []string
}

type BTCCTicker struct {
	High       float64 `json:",string"`
	Low        float64 `json:",string"`
	Buy        float64 `json:",string"`
	Sell       float64 `json:",string"`
	Last       float64 `json:",string"`
	Vol        float64 `json:",string"`
	Date       int64
	Vwap       float64 `json:",string"`
	Prev_close float64 `json:",string"`
	Open       float64 `json:",string"`
}

type BTCCProfile struct {
	Username             string
	TradePasswordEnabled bool    `json:"trade_password_enabled,bool"`
	OTPEnabled           bool    `json:"otp_enabled,bool"`
	TradeFee             float64 `json:"trade_fee"`
	TradeFeeCNYLTC       float64 `json:"trade_fee_cnyltc"`
	TradeFeeBTCLTC       float64 `json:"trade_fee_btcltc"`
	DailyBTCLimit        float64 `json:"daily_btc_limit"`
	DailyLTCLimit        float64 `json:"daily_ltc_limit"`
	BTCDespoitAddress    string  `json:"btc_despoit_address"`
	BTCWithdrawalAddress string  `json:"btc_withdrawal_address"`
	LTCDepositAddress    string  `json:"ltc_deposit_address"`
	LTCWithdrawalAddress string  `json:"ltc_withdrawal_request"`
	APIKeyPermission     int64   `json:"api_key_permission"`
}

type BTCCCurrencyGeneric struct {
	Currency      string
	Symbol        string
	Amount        string
	AmountInt     string `json:"amount_integer"`
	AmountDecimal int64  `json:"amount_decimal"`
}

type BTCCBTCCCurrency struct {
	Btc BTCCCurrencyGeneric
	Ltc BTCCCurrencyGeneric
	Cny BTCCCurrencyGeneric
}

type BTCCOrder struct {
	ID         int64
	Type       string
	Price      float64
	Currency   string
	Amount     float64
	AmountOrig float64 `json:"amount_original"`
	Date       int64
	Status     string
	Detail     BTCCOrderDetail
}

type BTCCOrderDetail struct {
	Dateline int64
	Price    float64
	Amount   float64
}

type BTCCWithdrawal struct {
	ID          int64
	Address     string
	Currency    string
	Amount      float64
	Date        int64
	Transaction string
	Status      string
}

type BTCCDeposit struct {
	ID       int64
	Address  string
	Currency string
	Amount   float64
	Date     int64
	Status   string
}

type BTCCBidAsk struct {
	Price  float64
	Amount float64
}

type BTCCDepth struct {
	Bid []BTCCBidAsk
	Ask []BTCCBidAsk
}

type BTCCTransaction struct {
	ID        int64
	Type      string
	BTCAmount float64 `json:"btc_amount"`
	LTCAmount float64 `json:"ltc_amount"`
	CNYAmount float64 `json:"cny_amount"`
	Date      int64
}

type BTCCIcebergOrder struct {
	ID              int64
	Type            string
	Price           float64
	Market          string
	Amount          float64
	AmountOrig      float64 `json:"amount_original"`
	DisclosedAmount float64 `json:"disclosed_amount"`
	Variance        float64
	Date            int64
	Status          string
}

type BTCCStopOrder struct {
	ID          int64
	Type        string
	StopPrice   float64 `json:"stop_price"`
	TrailingAmt float64 `json:"trailing_amount"`
	TrailingPct float64 `json:"trailing_percentage"`
	Price       float64
	Market      string
	Amount      float64
	Date        int64
	Status      string
	OrderID     int64 `json:"order_id"`
}

func (b *BTCC) SetDefaults() {
	b.Name = "BTCC"
	b.Enabled = false
	b.Fee = 0
	b.RESTPollingDelay = 10000
}

//Setup is run on startup to setup exchange with config values
func (b *BTCC) Setup(exconfig ExchangeConfig) {
	if !exconfig.Enabled {
		b.SetEnabled(false)
	} else {
		b.Enabled = true
		b.RESTPollingDelay = exconfig.RESTPollingDelay
		b.AvailablePairs = SplitStrings(exconfig.AvailablePairs, ",")
	}
}

func (k *BTCC) GetEnabledCurrencies() []string {
	return k.EnabledPairs
}

//Start is run if exchange is enabled, after Setup
func (b *BTCC) Start() {
	go func() {
		if verbose {
			log.Printf("%s polling delay: %d ms.\n", b.GetName(), b.RESTPollingDelay)
			log.Printf("%s %d currencies enabled: %s.\n", b.GetName(), len(b.EnabledPairs), b.EnabledPairs)
		}
	}()
}

func (b *BTCC) GetName() string {
	return b.Name
}

func (b *BTCC) GetLabel() string {
	return b.Label
}

func (b *BTCC) SetEnabled(enabled bool) {
	b.Enabled = enabled
}

func (b *BTCC) IsEnabled() bool {
	return b.Enabled
}

func (b *BTCC) GetFee() float64 {
	return b.Fee
}

// 	for b.Enabled {

// 		for _, x := range b.EnabledPairs {
// 			currency := StringToLower(x)
// 			go func() {
// 				ticker := b.GetTicker(currency)
// 				if currency != "ltcbtc" {
// 					tickerLastUSD, _ := ConvertCurrency(ticker.Last, "CNY", "USD")
// 					tickerHighUSD, _ := ConvertCurrency(ticker.High, "CNY", "USD")
// 					tickerLowUSD, _ := ConvertCurrency(ticker.Low, "CNY", "USD")
// 					log.Printf("BTCC %s: Last %f (%f) High %f (%f) Low %f (%f) Volume %f\n", currency, tickerLastUSD, ticker.Last, tickerHighUSD, ticker.High, tickerLowUSD, ticker.Low, ticker.Vol)
// 					//AddExchangeInfo(b.GetName(), StringToUpper(currency[0:3]), StringToUpper(currency[3:]), ticker.Last, ticker.Vol)
// 					//AddExchangeInfo(b.GetName(), StringToUpper(currency[0:3]), "USD", tickerLastUSD, ticker.Vol)
// 				} else {
// 					log.Printf("BTCC %s: Last %f High %f Low %f Volume %f\n", currency, ticker.Last, ticker.High, ticker.Low, ticker.Vol)
// 					//AddExchangeInfo(b.GetName(), StringToUpper(currency[0:3]), StringToUpper(currency[3:]), ticker.Last, ticker.Vol)
// 				}
// 			}()
// 		}
// 		time.Sleep(time.Millisecond * b.RESTPollingDelay)
// 	}
// }

func (b *BTCC) GetTicker(symbol string) BTCCTicker {
	type Response struct {
		Ticker BTCCTicker
	}

	resp := Response{}
	req := fmt.Sprintf("%sdata/ticker?market=%s", BTCC_API_URL, symbol)
	err := SendHTTPGetRequest(req, true, &resp)
	if err != nil {
		log.Println(err)
		return BTCCTicker{}
	}
	return resp.Ticker
}

func (b *BTCC) GetTickerPrice(currency string) TickerPrice {
	var tickerPrice TickerPrice
	ticker := b.GetTicker(currency)
	tickerPrice.Ask = ticker.Sell
	tickerPrice.Bid = ticker.Buy
	tickerPrice.CryptoCurrency = currency
	tickerPrice.Low = ticker.Low
	tickerPrice.Last = ticker.Last
	tickerPrice.Volume = ticker.Vol
	tickerPrice.High = ticker.High

	return tickerPrice
}

func (b *BTCC) GetTradesLast24h(symbol string) bool {
	req := fmt.Sprintf("%sdata/trades?market=%s", BTCC_API_URL, symbol)
	err := SendHTTPGetRequest(req, true, nil)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (b *BTCC) GetTradeHistory(symbol string, limit, sinceTid int64, time time.Time) bool {
	req := fmt.Sprintf("%sdata/historydata?market=%s", BTCC_API_URL, symbol)
	v := url.Values{}

	if limit > 0 {
		v.Set("limit", strconv.FormatInt(limit, 10))
	}
	if sinceTid > 0 {
		v.Set("since", strconv.FormatInt(sinceTid, 10))
	}
	if !time.IsZero() {
		v.Set("sincetype", strconv.FormatInt(time.Unix(), 10))
	}

	req = EncodeURLValues(req, v)
	err := SendHTTPGetRequest(req, true, nil)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (b *BTCC) GetOrderBook(symbol string, limit int) bool {
	req := fmt.Sprintf("%sdata/orderbook?market=%s&limit=%d", BTCC_API_URL, symbol, limit)
	err := SendHTTPGetRequest(req, true, nil)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//参数可以是“all”，“balance”，“frozen”, “loan”或者“profile”，默认为“all”.
//包含或者是如下对象：profile, balance, frozen, loan
func (b *BTCC) GetAccount() (ExchangeAccountInfo, error) {
	info := ExchangeAccountInfo{}

	type ResultStruct struct {
		Profile BTCCProfile
		Balance BTCCBTCCCurrency
		Frozen  BTCCBTCCCurrency
		Loan    BTCCBTCCCurrency
	}

	type Response struct {
		Id     string
		Result ResultStruct
	}
	resp := Response{}
	params := make([]interface{}, 0)

	params = append(params, "all")
	respstr, err := b.SendAuthenticatedHTTPRequest(BTCC_ACCOUNT_INFO, params)

	if err != nil {
		log.Println(err)
		return ExchangeAccountInfo{}, err
	}

	err = JSONDecode([]byte(respstr), &resp)
	if err != nil {
		log.Println(err.Error())
		return info, errors.New("unable to JSON Unmarshal response")
	}

	info.ExchangeName = resp.Result.Profile.Username
	var currency ExchangeAccountCurrencyInfo

	currency.CurrencyName = "BTC"
	currency.Hold, _ = strconv.ParseFloat(resp.Result.Frozen.Btc.Amount, 64)
	currency.TotalValue, _ = strconv.ParseFloat(resp.Result.Balance.Btc.Amount, 64)
	info.Currencies = append(info.Currencies, currency)

	currency.CurrencyName = "LTC"
	currency.Hold, _ = strconv.ParseFloat(resp.Result.Frozen.Ltc.Amount, 64)
	currency.TotalValue, _ = strconv.ParseFloat(resp.Result.Balance.Btc.Amount, 64)
	info.Currencies = append(info.Currencies, currency)

	currency.CurrencyName = "CNY"
	currency.Hold, _ = strconv.ParseFloat(resp.Result.Frozen.Cny.Amount, 64)
	currency.TotalValue, _ = strconv.ParseFloat(resp.Result.Balance.Btc.Amount, 64)
	info.Currencies = append(info.Currencies, currency)

	return info, nil
}

func (b *BTCC) PlaceOrder(buyOrder bool, price, amount float64, market string) {
	params := make([]interface{}, 0)
	params = append(params, strconv.FormatFloat(price, 'f', -1, 64))
	params = append(params, strconv.FormatFloat(amount, 'f', -1, 64))

	if len(market) > 0 {
		params = append(params, market)
	}

	req := BTCC_ORDER_BUY
	if !buyOrder {
		req = BTCC_ORDER_SELL
	}

	_, err := b.SendAuthenticatedHTTPRequest(req, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) CancelOrder(orderID int64, market string) {
	params := make([]interface{}, 0)
	params = append(params, orderID)

	if len(market) > 0 {
		params = append(params, market)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_ORDER_CANCEL, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetDeposits(currency string, pending bool) {
	params := make([]interface{}, 0)
	params = append(params, currency)

	if pending {
		params = append(params, pending)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_DEPOSITS, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetMarketDepth(market string, limit int64) {
	params := make([]interface{}, 0)

	if limit > 0 {
		params = append(params, limit)
	}

	if len(market) > 0 {
		params = append(params, market)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_MARKETDEPTH, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetOrder(orderID int64, market string, detailed bool) {
	params := make([]interface{}, 0)
	params = append(params, orderID)

	if len(market) > 0 {
		params = append(params, market)
	}

	if detailed {
		params = append(params, detailed)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_ORDER, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetOrders(openonly bool, market string, limit, offset, since int64, detailed bool) {
	params := make([]interface{}, 0)

	if openonly {
		params = append(params, openonly)
	}

	if len(market) > 0 {
		params = append(params, market)
	}

	if limit > 0 {
		params = append(params, limit)
	}

	if offset > 0 {
		params = append(params, offset)
	}

	if since > 0 {
		params = append(params, since)
	}

	if detailed {
		params = append(params, detailed)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_ORDERS, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetTransactions(transType string, limit, offset, since int64, sinceType string) {
	params := make([]interface{}, 0)

	if len(transType) > 0 {
		params = append(params, transType)
	}

	if limit > 0 {
		params = append(params, limit)
	}

	if offset > 0 {
		params = append(params, offset)
	}

	if since > 0 {
		params = append(params, since)
	}

	if len(sinceType) > 0 {
		params = append(params, sinceType)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_TRANSACTIONS, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetWithdrawal(withdrawalID int64, currency string) {
	params := make([]interface{}, 0)
	params = append(params, withdrawalID)

	if len(currency) > 0 {
		params = append(params, currency)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_WITHDRAWAL, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetWithdrawals(currency string, pending bool) {
	params := make([]interface{}, 0)
	params = append(params, currency)

	if pending {
		params = append(params, pending)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_WITHDRAWALS, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) RequestWithdrawal(currency string, amount float64) {
	params := make([]interface{}, 0)
	params = append(params, currency)
	params = append(params, amount)

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_WITHDRAWAL_REQUEST, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) IcebergOrder(buyOrder bool, price, amount, discAmount, variance float64, market string) {
	params := make([]interface{}, 0)
	params = append(params, strconv.FormatFloat(price, 'f', -1, 64))
	params = append(params, strconv.FormatFloat(amount, 'f', -1, 64))
	params = append(params, strconv.FormatFloat(discAmount, 'f', -1, 64))
	params = append(params, strconv.FormatFloat(variance, 'f', -1, 64))

	if len(market) > 0 {
		params = append(params, market)
	}

	req := BTCC_ICEBERG_BUY
	if !buyOrder {
		req = BTCC_ICEBERG_SELL
	}

	_, err := b.SendAuthenticatedHTTPRequest(req, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetIcebergOrder(orderID int64, market string) {
	params := make([]interface{}, 0)
	params = append(params, orderID)

	if len(market) > 0 {
		params = append(params, market)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_ICEBERG_ORDER, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetIcebergOrders(limit, offset int64, market string) {
	params := make([]interface{}, 0)

	if limit > 0 {
		params = append(params, limit)
	}

	if offset > 0 {
		params = append(params, offset)
	}

	if len(market) > 0 {
		params = append(params, market)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_ICEBERG_ORDERS, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) CancelIcebergOrder(orderID int64, market string) {
	params := make([]interface{}, 0)
	params = append(params, orderID)

	if len(market) > 0 {
		params = append(params, market)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_ICEBERG_CANCEL, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) PlaceStopOrder(buyOder bool, stopPrice, price, amount, trailingAmt, trailingPct float64, market string) {
	params := make([]interface{}, 0)

	if stopPrice > 0 {
		params = append(params, stopPrice)
	}

	params = append(params, strconv.FormatFloat(price, 'f', -1, 64))
	params = append(params, strconv.FormatFloat(amount, 'f', -1, 64))

	if trailingAmt > 0 {
		params = append(params, strconv.FormatFloat(trailingAmt, 'f', -1, 64))
	}

	if trailingPct > 0 {
		params = append(params, strconv.FormatFloat(trailingPct, 'f', -1, 64))
	}

	if len(market) > 0 {
		params = append(params, market)
	}

	req := BTCC_STOPORDER_BUY
	if !buyOder {
		req = BTCC_STOPORDER_SELL
	}

	_, err := b.SendAuthenticatedHTTPRequest(req, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetStopOrder(orderID int64, market string) {
	params := make([]interface{}, 0)
	params = append(params, orderID)

	if len(market) > 0 {
		params = append(params, market)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_STOPORDER, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) GetStopOrders(status, orderType string, stopPrice float64, limit, offset int64, market string) {
	params := make([]interface{}, 0)

	if len(status) > 0 {
		params = append(params, status)
	}

	if len(orderType) > 0 {
		params = append(params, orderType)
	}

	if stopPrice > 0 {
		params = append(params, stopPrice)
	}

	if limit > 0 {
		params = append(params, limit)
	}

	if offset > 0 {
		params = append(params, limit)
	}

	if len(market) > 0 {
		params = append(params, market)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_STOPORDERS, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) CancelStopOrder(orderID int64, market string) {
	params := make([]interface{}, 0)
	params = append(params, orderID)

	if len(market) > 0 {
		params = append(params, market)
	}

	_, err := b.SendAuthenticatedHTTPRequest(BTCC_STOPORDER_CANCEL, params)

	if err != nil {
		log.Println(err)
	}
}

func (b *BTCC) SendAuthenticatedHTTPRequest(method string, params []interface{}) (resp string, err error) {
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)[0:16]
	encoded := fmt.Sprintf("tonce=%s&accesskey=%s&requestmethod=post&id=%d&method=%s&params=", nonce, b.APIKey, 1, method)

	if len(params) == 0 {
		params = make([]interface{}, 0)
	} else {
		items := make([]string, 0)
		for _, x := range params {
			xType := fmt.Sprintf("%T", x)
			switch xType {
			case "int64", "int":
				{
					items = append(items, fmt.Sprintf("%d", x))
				}
			case "string":
				{
					items = append(items, fmt.Sprintf("%s", x))
				}
			case "float64":
				{
					items = append(items, fmt.Sprintf("%f", x))
				}
			case "bool":
				{
					if x == true {
						items = append(items, "1")
					} else {
						items = append(items, "")
					}
				}
			default:
				{
					items = append(items, fmt.Sprintf("%v", x))
				}
			}
		}
		encoded += JoinStrings(items, ",")
	}
	if verbose {
		log.Println(encoded)
	}

	hmac := GetHMAC(HASH_SHA1, []byte(encoded), []byte(b.APISecret))
	postData := make(map[string]interface{})
	postData["method"] = method
	postData["params"] = params
	postData["id"] = 1
	apiURL := BTCC_API_URL + BTCC_API_AUTHENTICATED_METHOD
	data, err := JSONEncode(postData)

	if err != nil {
		return "", errors.New("Unable to JSON Marshal POST data")
	}

	if verbose {
		log.Printf("Sending POST request to %s calling method %s with params %s\n", apiURL, method, data)
	}

	headers := make(map[string]string)
	headers["Content-type"] = "application/json-rpc"
	headers["Authorization"] = "Basic " + Base64Encode([]byte(b.APIKey+":"+HexEncodeToString(hmac)))
	headers["Json-Rpc-Tonce"] = nonce

	resp, err = SendHTTPRequest("POST", apiURL, headers, strings.NewReader(string(data)))

	if verbose {
		log.Printf("Recv'd :%s\n", resp)
	}

	return resp, nil
}
