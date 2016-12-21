package main

import (
	"log"
	"net/smtp"
	"time"

	"fmt"
	"net"

	"os"
	"strings"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

//"github.com/d4l3k/talib"

type BotStatus int

const (
	BOT_RUNNING BotStatus = iota
	BOT_STOPPED
	BOT_STARTING
	BOT_STOPPING
)

type BotLogType int

const (
	ERROR_LOG BotLogType = iota
	PROFIT_LOG
	INFO_LOG
	BUY_LOG
	SELL_LOG
	CANCEL_LOG
)

type BotLog struct {
	Stamp        time.Time
	ExchangeName string
	Type         BotLogType
	Price        float64
	Amount       float64
	Message      string
}

type Bot struct {
	Name                  string
	StrategyID            string
	DockerID              string
	KLineInterval         KLineIntervalType
	ExchangeAccountAPIIDs string
	Status                BotStatus
	StatusDescription     string
	Enabled               bool
	Backtesting           bool
	PaperTrading          bool

	//StartTime time.Time
	//StopTime  time.Time
	StrategyVarabileValues map[string]StrategyVarabileValue

	jsVM      *otto.Otto
	docker    *Docker
	exchanges []IExchange

	botLogs       []BotLog
	commands      string
	lastError     string
	errorFilter   string
	botStatusLog  string
	enableAutoLog bool
}

func (bot *Bot) Start(docker *Docker, code string, exchanges []IExchange) (err error) {
	log.Printf("Bot %s Started!\n", bot.Name)
	bot.docker = docker
	bot.exchanges = exchanges

	bot.jsVM = otto.New()

	bot.setupJSVM()

	go func() {
		if _, err = bot.jsVM.Run(code); err != nil {
			log.Print(err.Error())
			return
		}
		mainfunc, _ := bot.jsVM.Get("main")
		_, err = mainfunc.Call(otto.NullValue())
		if err != nil {
			log.Print(err.Error())
			return
		}
	}()

	return
}
func (bot *Bot) setupJSVM() {
	bot.jsVM.Set("Version", bot.docker.version)
	bot.jsVM.Set("exchanges", bot.exchanges)
	bot.jsVM.Set("exchange", bot.exchanges[0])

	bot.jsVM.Set("Sleep", bot.Sleep)
	bot.jsVM.Set("LogProfit", bot.LogProfit)
	bot.jsVM.Set("IsBacktesting", bot.IsBacktesting)
	bot.jsVM.Set("IsPaperTrading", bot.IsPaperTrading)
	bot.jsVM.Set("Dial", bot.Dial)
	bot.jsVM.Set("HttpQuery", bot.HttpQuery)
	bot.jsVM.Set("Mail", bot.Mail)
	bot.jsVM.Set("GetCommand", bot.GetCommand)
	bot.jsVM.Set("GetPid", bot.GetPid)
	bot.jsVM.Set("Log", bot.Log)
	bot.jsVM.Set("GetLastError", bot.GetLastError)
	bot.jsVM.Set("SetErrorFilter", bot.SetErrorFilter)
	bot.jsVM.Set("EnableLog", bot.EnableLog)
	bot.jsVM.Set("LogStatus", bot.LogStatus)
	bot.jsVM.Set("LogReset", bot.LogReset)
	bot.jsVM.Set("LogProfitReset", bot.LogProfitReset)
	bot.jsVM.Set("LogProfit", bot.LogProfit)
	//bot.jsVM.Set("Sin", talib.Sin)

}

func (bot *Bot) Stop() {
	if bot.jsVM != nil {
		exitfunc, _ := bot.jsVM.Get("onExit")
		_, err := exitfunc.Call(otto.NullValue())
		if err != nil {
			log.Print(err.Error())
		}
	} else {
		log.Println("jsvm is nil")
	}
}

/*TODO
Chart({...})
_G(K, V)
//EnableWebsocket()
*/

// func (bot Bot) Remove() {

// }

// func (Bot Bot) OnConfigChanged() {

// }

func (bot *Bot) Sleep(ms int64) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

//Log is exported to js VM. only log strings
func (bot *Bot) Log(log string) {
	if verbose {
		fmt.Println("Get log from JS VM:" + log)
	}
	var botLog BotLog
	botLog.Message = log
	botLog.Type = INFO_LOG
	bot.botLogs = append(bot.botLogs, botLog)
	return
}

func (bot *Bot) LogProfit(profit float64) {
	var botLog BotLog
	botLog.Amount = profit
	botLog.Type = PROFIT_LOG
	bot.botLogs = append(bot.botLogs, botLog)
	return
}

func (bot *Bot) IsBacktesting() bool {
	return bot.Backtesting
}

func (bot *Bot) IsPaperTrading() bool {
	return bot.PaperTrading
}

//Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only),
//"udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".
//For TCP and UDP networks, addresses have the form host:port.
//If host is a literal IPv6 address it must be enclosed in square brackets as in "[::1]:80"
//or "[ipv6-host%zone]:80". The functions JoinHostPort and SplitHostPort manipulate addresses in this form.
//If the host is empty, as in ":80", the local system is assumed.

func (bot *Bot) Dial(network, address string, timeout time.Duration) net.Conn {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		fmt.Println(err.Error())
	}
	//TODO
	return conn
}

func (bot *Bot) HttpQuery(url string, postData string, cookies string, headers string, isReturnHeader bool) string {
	//TODO
	// var client http.Client
	// var method string
	// if postData == "" || postData == nil {
	// 	method = "POST"
	// } else {
	// 	method = "GET"
	// }

	// if req, err := http.NewRequest(method, url, nil); err != nil {
	// 	log.Println(err.Error)
	// 	return ""
	// }

	return ""
}

func (bot *Bot) Mail(smtpServer string, smtpUsername string, smtpPassword string, mailTo string, title string, body string) bool {
	// Set up authentication information.
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := strings.Split(mailTo, ";")

	msg := []byte("To:" + to[0] + "\r\n" +
		"Subject: " + title + "\r\n" +
		body + " \r\n")
	err := smtp.SendMail(smtpServer, auth, "noreply@bitbot.com.cn", to, msg)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (bot *Bot) GetCommand() string {
	cmd := bot.commands
	bot.commands = ""
	return cmd
}

func (bot *Bot) GetPid() int {
	return os.Getpid()
}

func (bot *Bot) GetLastError() string {
	return bot.lastError
}

func (bot *Bot) SetErrorFilter(filter string) {
	bot.errorFilter = filter
	return
}

func (bot *Bot) EnableLog(enable bool) {
	bot.enableAutoLog = enable
	return
}

func (bot *Bot) LogStatus(status string) {
	bot.botStatusLog = status
	return
}

func (bot *Bot) LogReset(recordsToLeave interface{}) {

	if records, ok := recordsToLeave.(int); ok {
		if records < len(bot.botLogs) {
			bot.botLogs = bot.botLogs[len(bot.botLogs)-records:]
		}
	} else {
		bot.botLogs = bot.botLogs[0:0]
	}
	return
}

func (bot *Bot) LogProfitReset(recordsToLeave int) {
	return
}

// func runUnsafe(unsafe string) {
// 	start := time.Now()
// 	defer func() {
// 		duration := time.Since(start)
// 		if caught := recover(); caught != nil {
// 			if caught == halt {
// 				fmt.Fprintf(os.Stderr, "Some code took to long! Stopping after: %v\n", duration)
// 				return
// 			}
// 			panic(caught) // Something else happened, repanic!
// 		}
// 		fmt.Fprintf(os.Stderr, "Ran code successfully: %v\n", duration)
// 	}()

// 	vm := otto.New()
// 	vm.Interrupt = make(chan func(), 1) // The buffer prevents blocking

// 	go func() {
// 		time.Sleep(2 * time.Second) // Stop after two seconds
// 		vm.Interrupt <- func() {
// 			panic(halt)
// 		}
// 	}()

// 	vm.Run(unsafe) // Here be dragons (risky code)
// }
