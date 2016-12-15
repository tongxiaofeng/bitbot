package bitbot

import (
	"time"
)

type KLineIntervalType int

const (
	KLINE_1_MIN KLineIntervalType = 1 + iota
	KLINE_5_MIN
	KLINE_15_MIN
	KLINE_30_MIN
	KLINE_1_HOUR
	KLINE_1_DAY
)

type BotStatus int

const (
	BOT_RUNNING BotStatus = 1 + iota
	BOT_STOPPED
	BOT_STARTING
	BOT_STOPPING
)

type BotLogType int

const (
	ERROR_LOG BotLogType = 1 + iota
	PROFIT_LOG
	INFO_LOG
	BUY_LOG
	SELL_LOG
	CANCEL_LOG
)

type BotLog struct {
	Date     time.Time
	Exchange ExchangeType
	Type     BotLogType
	Price    float64
	Amount   float64
	Message  string
}

type Bot struct {
	ID                  int
	Name                string
	StrategyID          int
	DockerID            int
	KLineInterval               KLineIntervalType
	ExchangeAccountAPIs []ExchangeAPI
	Status              BotStatus
	StatusDescription   string

	StartTime time.Time
	StopTime  time.Time

	StrategyVarabiles []StrategyVarabile
	strategy_code     string
}

var bots []Bot

//ReadConfig reads bots configurations and construct bots.
func(bot *Bot) ReadConfig(){

}  

func (bot *Bot) Run(){

}

func (bot *Bot) Stop(){

}

func (bot *Bot) Remove(){
	
}

func (Bot *Bot) OnConfigChanged(){

}