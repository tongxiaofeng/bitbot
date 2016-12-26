package main

type OKCoinExchangeAccount struct {
	ExchangeAccountConfig
	okcoin *OKCoinExchange
	symbol string
}

func (ok *OKCoinExchangeAccount) Setup(acctConfig ExchangeAccountConfig) {
	ok.ExchangeAccountConfig = acctConfig
	ok.okcoin = new(OKCoinExchange)
	if ok.ExchangeAccountConfig.ExchangeName == "OKCOINUSD" {
		ok.okcoin.SetURL(OKCOIN_API_URL)
	} else {
		ok.okcoin.SetURL(OKCOIN_API_URL_CHINA)
	}
}
func (ok *OKCoinExchangeAccount) GetName() string {
	return ok.ExchangeName
}

func (ok *OKCoinExchangeAccount) GetLabel() string {
	return ok.Label
}

func (ok *OKCoinExchangeAccount) GetAccount() (ExchangeAccountInfo, error) {
	return ExchangeAccountInfo{}, nil
}

func (ok *OKCoinExchangeAccount) GetEnabledPair() string {
	return ok.EnabledPair
}

func (ok *OKCoinExchangeAccount) Start() {

}

func (ok *OKCoinExchangeAccount) SetDefaults() {

}

func (ok *OKCoinExchangeAccount) GetTicker() Ticker {
	okTicker, err := ok.okcoin.GetTicker("BTCCNY")
	if err != nil {
		return Ticker{}
	}
	ticker := Ticker{}

	ticker.Buy = okTicker.Buy
	ticker.High = okTicker.High
	ticker.Last = okTicker.Last
	ticker.Low = okTicker.Low
	ticker.Sell = okTicker.Sell
	ticker.Volume = okTicker.Vol

	return ticker
}
