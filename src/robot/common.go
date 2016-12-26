package main

// type TickerPrice struct {
// 	CryptoCurrency string  `json:"CryptoCurrency"`
// 	FiatCurrency   string  `json:"FiatCurrency"`
// 	Last           float64 `json:"Last"`
// 	High           float64 `json:"High"`
// 	Low            float64 `json:"Low"`
// 	Bid            float64 `json:"Bid"`
// 	Ask            float64 `json:"Ask"`
// 	Volume         float64 `json:"Volume"`
// }

// type OldTicker struct {
// 	Price        map[string]map[string]TickerPrice
// 	ExchangeName string
// }

// func (t *Ticker) PriceToString(cryptoCurrency, fiatCurrency, priceType string) string {
// 	switch priceType {
// 	case "last":
// 		return strconv.FormatFloat(t.Price[cryptoCurrency][fiatCurrency].Last, 'f', -1, 64)
// 	case "high":
// 		return strconv.FormatFloat(t.Price[cryptoCurrency][fiatCurrency].High, 'f', -1, 64)
// 	case "low":
// 		return strconv.FormatFloat(t.Price[cryptoCurrency][fiatCurrency].Low, 'f', -1, 64)
// 	case "bid":
// 		return strconv.FormatFloat(t.Price[cryptoCurrency][fiatCurrency].Bid, 'f', -1, 64)
// 	case "ask":
// 		return strconv.FormatFloat(t.Price[cryptoCurrency][fiatCurrency].Ask, 'f', -1, 64)
// 	case "volume":
// 		return strconv.FormatFloat(t.Price[cryptoCurrency][fiatCurrency].Volume, 'f', -1, 64)
// 	default:
// 		return ""
// 	}
// }

// func AddTickerPrice(m map[string]map[string]TickerPrice, cryptocurrency, fiatcurrency string, price TickerPrice) {
// 	mm, ok := m[cryptocurrency]
// 	if !ok {
// 		mm = make(map[string]TickerPrice)
// 		m[cryptocurrency] = mm
// 	}
// 	mm[fiatcurrency] = price
// }

// func NewTicker(exchangeName string, prices []TickerPrice) *Ticker {
// 	ticker := &Ticker{}
// 	ticker.ExchangeName = exchangeName
// 	ticker.Price = make(map[string]map[string]TickerPrice, 0)

// 	for x, _ := range prices {
// 		AddTickerPrice(ticker.Price, prices[x].CryptoCurrency, prices[x].FiatCurrency, prices[x])
// 	}

// 	return ticker
// }

// const (
// 	LIMIT_ORDER = iota
// 	MARKET_ORDER
// )

// var Orders []*Order

// type Order struct {
// 	OrderID  int
// 	Exchange string
// 	Type     int
// 	Amount   float64
// 	Price    float64
// }

// func NewOrder(Exchange string, amount, price float64) int {
// 	order := &Order{}
// 	if len(Orders) == 0 {
// 		order.OrderID = 0
// 	} else {
// 		order.OrderID = len(Orders)
// 	}

// 	order.Exchange = Exchange
// 	order.Amount = amount
// 	order.Price = price
// 	Orders = append(Orders, order)
// 	return order.OrderID
// }

// func DeleteOrder(orderID int) bool {
// 	for i := range Orders {
// 		if Orders[i].OrderID == orderID {
// 			Orders = append(Orders[:i], Orders[i+1:]...)
// 			return true
// 		}
// 	}
// 	return false
// }

// func GetOrdersByExchange(exchange string) ([]*Order, bool) {
// 	orders := []*Order{}
// 	for i := range Orders {
// 		if Orders[i].Exchange == exchange {
// 			orders = append(orders, Orders[i])
// 		}
// 	}
// 	if len(orders) > 0 {
// 		return orders, true
// 	}
// 	return nil, false
// }

// func GetOrderByOrderID(orderID int) (*Order, bool) {
// 	for i := range Orders {
// 		if Orders[i].OrderID == orderID {
// 			return Orders[i], true
// 		}
// 	}
// 	return nil, false
// }
