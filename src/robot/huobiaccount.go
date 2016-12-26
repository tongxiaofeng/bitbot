package main

type HUOBIExchangeAccount struct {
	ExchangeAccountConfig
	huobi *HUOBIExchange
}

func (h *HUOBIExchangeAccount) Setup(acctConfig ExchangeAccountConfig) {
	h.ExchangeAccountConfig = acctConfig
	//ToDo this should be changed
	h.huobi = new(HUOBIExchange)
}
func (h *HUOBIExchangeAccount) GetName() string {
	return h.ExchangeName
}

func (h *HUOBIExchangeAccount) GetLabel() string {
	return h.Label
}

func (h *HUOBIExchangeAccount) GetAccount() (ExchangeAccountInfo, error) {
	return ExchangeAccountInfo{}, nil
}

func (h *HUOBIExchangeAccount) GetEnabledPair() string {
	return h.EnabledPair
}

func (h *HUOBIExchangeAccount) Start() {

}

func (h *HUOBIExchangeAccount) SetDefaults() {

}

func (h *HUOBIExchangeAccount) GetTicker() Ticker {
	huobiTicker := h.huobi.GetTicker("BTCCNY")
	ticker := Ticker{}

	ticker.Buy = huobiTicker.Buy
	ticker.High = huobiTicker.High
	ticker.Last = huobiTicker.Last
	ticker.Low = huobiTicker.Low
	ticker.Sell = huobiTicker.Sell
	ticker.Volume = huobiTicker.Vol

	return ticker
}
