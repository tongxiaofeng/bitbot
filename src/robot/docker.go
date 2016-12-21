package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type DockerStatus int

const (
	DOCKER_CONNECTED DockerStatus = iota
	DOCKER_DISCONNECTED
	DOCKER_CONNECTING
	DOCKER_DISCONNECTING
)

type Docker struct {
	ID            int
	ServerAddress string
	IP            string
	Port          int
	OS            string
	OSVersion     string
	DockerVersion string
	Status        DockerStatus
	IsLocal       bool // Is it running on the same machine as server?
	Verbose       bool

	Cryptocurrencies string
	DisplayCurrency  string
	ExchangeConfigs  map[string]ExchangeConfig
	ExchangeAccounts map[string]ExchangeAccount
	Strategies       map[string]Strategy
	Bots             map[string]*Bot
}

const (
	CONFIG_FILE = "config.json"
)

var (
	ErrExchangeNameEmpty                       = "Exchange #%d in config: Exchange name is empty."
	ErrExchangeAvailablePairsEmpty             = "Exchange %s: Available pairs is empty."
	ErrExchangeEnabledPairsEmpty               = "Exchange %s: Enabled pairs is empty."
	ErrExchangeBaseCurrenciesEmpty             = "Exchange %s: Base currencies is empty."
	WarningExchangeAuthAPIDefaultOrEmptyValues = "WARNING -- Exchange %s: Authenticated API support disabled due to default/empty APIKey/Secret/ClientID values."
	ErrExchangeNotFound                        = "Exchange %s: Not found."
	ErrNoEnabledExchanges                      = "No Exchanges enabled."
	ErrCryptocurrenciesEmpty                   = "Cryptocurrencies variable is empty."
)

func (docker *Docker) CheckConfigValues() error {
	// if docker.Cryptocurrencies == "" {
	// 	return errors.New(ErrCryptocurrenciesEmpty)
	// }

	// exchanges := 0
	// for i, exch := range docker.ExchangeAccounts {
	// 	if exch.Enabled {
	// 		if exch.Name == "" {
	// 			return fmt.Errorf(ErrExchangeNameEmpty, i)
	// 		}
	// 		if exch.AvailablePairs == "" {
	// 			return fmt.Errorf(ErrExchangeAvailablePairsEmpty, exch.Name)
	// 		}
	// 		if exch.AuthenticatedAPISupport { // non-fatal error
	// 			if exch.APIKey == "" || exch.APISecret == "" || exch.APIKey == "Key" || exch.APISecret == "Secret" {
	// 				config.ExchangeAccounts[i].AuthenticatedAPISupport = false
	// 				log.Printf(WarningExchangeAuthAPIDefaultOrEmptyValues, exch.Name)
	// 				continue
	// 			} else if exch.Name == "ITBIT" || exch.Name == "Bitstamp" || exch.Name == "Coinbase" {
	// 				if exch.ClientID == "" || exch.ClientID == "ClientID" {
	// 					config.ExchangeAccounts[i].AuthenticatedAPISupport = false
	// 					log.Printf(WarningExchangeAuthAPIDefaultOrEmptyValues, exch.Name)
	// 					continue
	// 				}
	// 			}
	// 		}
	// 		exchanges++
	// 	}
	// }
	// if exchanges == 0 {
	// 	return errors.New(ErrNoEnabledExchanges)
	// }
	return nil
}

func (docker *Docker) ReadConfig() error {
	file, err := ioutil.ReadFile(CONFIG_FILE)

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, docker)
	return err
}

func (docker *Docker) SaveConfig() error {
	log.Println("Saving config")
	payload, err := json.MarshalIndent(docker, "", "  ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(CONFIG_FILE, payload, 0644)

	if err != nil {
		return err
	}
	retrieved, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		return err
	}

	if !bytes.Equal(retrieved, payload) {
		return fmt.Errorf("File %q content doesn't match, read %s expected %s\n", CONFIG_FILE, retrieved, payload)
	}

	return nil
}

func (docker *Docker) handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("Captured %v.", sig)
		docker.shutdown()
	}()
}

func (docker *Docker) shutdown() {
	log.Println("Bot shutting down..")
	err := docker.SaveConfig()

	if err != nil {
		log.Println("Unable to save config.")
	} else {
		log.Println("Config file saved successfully.")
	}

	for _, bot := range docker.Bots {
		if bot.Enabled {
			bot.Stop()
		}
	}

	log.Println("Exiting.")
	os.Exit(1)
}

// func (docker *Docker) Login()
// func (docker *Docker) Logout()

func (docker *Docker) Start() {

	for _, bot := range docker.Bots {
		if !bot.Enabled {
			continue
		}

		var code string
		var exchanges []IExchange
		var exchange IExchange

		//getting strategy code
		strategy := docker.Strategies[bot.StrategyID]
		if strategy.Code != "" {
			code = strategy.Code
		} else {
			if retrieved, err := ioutil.ReadFile(strategy.LocalPath); err != nil {
				log.Print(err.Error())
				continue
			} else {
				code = string(retrieved)
			}
		}

		//setup exchanges
		acctIDs := strings.Split(bot.ExchangeAccountAPIIDs, ",")
		for _, acctID := range acctIDs {
			acct := docker.ExchangeAccounts[acctID]
			switch acct.Exchange {
			case "BTCC":
				exchange = new(BTCC)
			case "HUOBI":
				exchange = new(HUOBI)
			case "OKCOINCHINA":
				exchange = new(OKCoin)
			case "OKCOININTERNATIONAL":
				exchange = new(OKCoin)
			default:
				log.Printf("Unknown Exchange %s.", acct.Exchange)
			}
			if exchange != nil {
				exchange.SetDefaults()
				exchange.Setup(docker.ExchangeConfigs[acct.Exchange])
				//exchange.Start()
				exchanges = append(exchanges, exchange)
			}
		}

		if err := bot.Start(docker, code, exchanges); err != nil {
			log.Fatal("Error starting bot." + err.Error())
		}
	}
	return

}

func (docker *Docker) version() string {
	return "Version:" + version + "\n" + "Buildstamp:" + buildstamp + "\n" + "Git Hash:" + githash
}
