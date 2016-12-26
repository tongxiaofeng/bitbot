package main

type BTCCExchangeAccount struct {
	ExchangeAccountConfig
	btcc *BTCCExchange
}

func (b *BTCCExchangeAccount) Setup(acctConfig ExchangeAccountConfig) {
	b.ExchangeAccountConfig = acctConfig
	b.btcc = new(BTCCExchange)
}
func (b *BTCCExchangeAccount) GetName() string {
	return b.ExchangeName
}

func (b *BTCCExchangeAccount) GetLabel() string {
	return b.Label
}

func (b *BTCCExchangeAccount) GetAccount() (ExchangeAccountInfo, error) {
	return ExchangeAccountInfo{}, nil
}

func (b *BTCCExchangeAccount) GetEnabledPair() string {
	return b.EnabledPair
}

func (b *BTCCExchangeAccount) Start() {

}

func (b *BTCCExchangeAccount) SetDefaults() {

}

func (b *BTCCExchangeAccount) GetTicker() Ticker {
	btccTicker := b.btcc.GetTicker("BTCCNY")
	ticker := Ticker{}

	ticker.Buy = btccTicker.Buy
	ticker.High = btccTicker.High
	ticker.Last = btccTicker.Last
	ticker.Low = btccTicker.Low
	ticker.Sell = btccTicker.Sell
	ticker.Volume = btccTicker.Vol

	return ticker
}
