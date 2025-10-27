package models

import "time"

type Notification struct {
	Id                int64   `json:"id"`
	CurrentPrice      float64 `json:"current_price"`
	MarketTradeVolume float64 `json:"market_trade_volume"`
	IntraDayHighPrice float64 `json:"intraday_high_price"`
	MarketCap         float64 `json:"market_cap"`
}

type MailStatus struct {
	Client   string    `json:"client"`
	Status   string    `json:"status"`
	SentTime time.Time `json:"sent_time"`
}
