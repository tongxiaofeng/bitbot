package main

import (
	"errors"
	"log"
)

type DockerStatus int

const (
	DOCKER_CONNECTED DockerStatus = iota
	DOCKER_DISCONNECTED
	DOCKER_CONNECTING
	DOCKER_DISCONNECTING
)

type DockerConfig struct {
	ID            int
	ServerAddress string
	IP            string
	Port          int
	OS            string
	OSVersion     string
	Status        DockerStatus
	IsLocal       bool // Is it running on the same machine as server?
	Verbose       bool

	Cryptocurrencies       string
	ExchangeConfigs        map[string]ExchangeConfig
	ExchangeAccountConfigs map[string]ExchangeAccountConfig
	StrategieConfigs       map[string]StrategyConfig
	BotConfigs             map[string]BotConfig
}

type Docker struct {
	DockerConfig
	Bots []*Bot
}

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

func (docker *Docker) Stop() {
	log.Println("Docker shutting down...")

	for _, bot := range docker.Bots {
		if bot.Enabled {
			bot.Stop()
		}
	}

	log.Println("All bots stopped.")
}

// func (docker *Docker) Login()
// func (docker *Docker) Logout()

func (docker *Docker) Start() error {
	err := docker.CheckConfigValues()

	if err != nil {
		log.Println("Fatal error checking config values. Error:", err)
		return err
	}

	for _, botConfig := range docker.BotConfigs {
		var bot Bot
		bot.BotConfig = botConfig
		bot.docker = docker
		docker.Bots = append(docker.Bots, &bot)

		if !bot.Enabled {
			continue
		} else {
			if err := bot.Start(); err != nil {
				log.Fatal(err)
			}
		}
	}
	return err
}

func (docker *Docker) version() string {
	return "Version:" + version + "   Buildstamp:" + buildstamp + "   Git Hash:" + githash
}

func (docker *Docker) SaveConfig() {
	//Saving configs
	var store IConfigStore
	if docker.IsLocal {
		store = new(FileConfigStore)
	} else {
		log.Fatal("No implementation for other config stores.")
		return
	}

	if err := store.Write(docker.DockerConfig); err != nil {
		log.Fatal("Error saving configs to store.")
	}
}

func (docker *Docker) ReadConfig() (err error) {
	//Reading configs
	var store IConfigStore
	store = FileConfigStore{}

	if docker.DockerConfig, err = store.Read(); err != nil {
		return
	}

	if err == nil && !docker.IsLocal {
		return errors.New("No implementation for other config stores now.")
	}

	return
}
