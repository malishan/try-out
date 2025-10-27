package models

import "errors"

type ApiModels struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data"`
}

// current price of BTC, market trade volume, intra day high price, market cap
type CreateNotification struct {
	CurrentPrice      float64 `json:"current_price"`
	MarketTradeVolume float64 `json:"market_trade_volume"`
	IntraDayHighPrice float64 `json:"intraday_high_price"`
	MarketCap         float64 `json:"market_cap"`
}

func (o *CreateNotification) Validation() error {
	if o.CurrentPrice == 0 {
		return errors.New("current price cannot be zero")
	}
	return nil
}
