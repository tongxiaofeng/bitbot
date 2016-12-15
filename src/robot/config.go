package main

import (
//"encoding/json"
//"errors"
//"fmt"
//"io/ioutil"
//"log"
//"time"
)

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

type Config struct {
	Cryptocurrencies string
	ExchangeAccounts []ExchangeAPI
}

// func GetEnabledExchanges() int {
// 	counter := 0
// 	for i := range bot.config.Exchanges {
// 		if bot.config.Exchanges[i].Enabled {
// 			counter++
// 		}
// 	}
// 	return counter
// }

// func GetExchangeConfig(name string) (Exchanges, error) {
// 	for i, _ := range bot.config.Exchanges {
// 		if bot.config.Exchanges[i].Name == name {
// 			return bot.config.Exchanges[i], nil
// 		}
// 	}
// 	return Exchanges{}, fmt.Errorf(ErrExchangeNotFound, name)
// }

// func UpdateExchangeConfig(e Exchanges) error {
// 	for i, _ := range bot.config.Exchanges {
// 		if bot.config.Exchanges[i].Name == e.Name {
// 			bot.config.Exchanges[i] = e
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf(ErrExchangeNotFound, e.Name)
// }

// func CheckExchangeConfigValues() error {
// 	if bot.config.Cryptocurrencies == "" {
// 		return errors.New(ErrCryptocurrenciesEmpty)
// 	}

// 	exchanges := 0
// 	for i, exch := range bot.config.Exchanges {
// 		if exch.Enabled {
// 			if exch.Name == "" {
// 				return fmt.Errorf(ErrExchangeNameEmpty, i)
// 			}
// 			if exch.AvailablePairs == "" {
// 				return fmt.Errorf(ErrExchangeAvailablePairsEmpty, exch.Name)
// 			}
// 			if exch.EnabledPairs == "" {
// 				return fmt.Errorf(ErrExchangeEnabledPairsEmpty, exch.Name)
// 			}
// 			if exch.BaseCurrencies == "" {
// 				return fmt.Errorf(ErrExchangeBaseCurrenciesEmpty, exch.Name)
// 			}
// 			if exch.AuthenticatedAPISupport { // non-fatal error
// 				if exch.APIKey == "" || exch.APISecret == "" || exch.APIKey == "Key" || exch.APISecret == "Secret" {
// 					bot.config.Exchanges[i].AuthenticatedAPISupport = false
// 					log.Printf(WarningExchangeAuthAPIDefaultOrEmptyValues, exch.Name)
// 					continue
// 				} else if exch.Name == "ITBIT" || exch.Name == "Bitstamp" || exch.Name == "Coinbase" {
// 					if exch.ClientID == "" || exch.ClientID == "ClientID" {
// 						bot.config.Exchanges[i].AuthenticatedAPISupport = false
// 						log.Printf(WarningExchangeAuthAPIDefaultOrEmptyValues, exch.Name)
// 						continue
// 					}
// 				}
// 			}
// 			exchanges++
// 		}
// 	}
// 	if exchanges == 0 {
// 		return errors.New(ErrNoEnabledExchanges)
// 	}
// 	return nil
// }

// func ReadConfig() (Config, error) {
// 	file, err := ioutil.ReadFile(CONFIG_FILE)

// 	if err != nil {
// 		return Config{}, err
// 	}

// 	cfg := Config{}
// 	err = json.Unmarshal(file, &cfg)
// 	return cfg, err
// }

// func SaveConfig() error {
// 	log.Println("Saving config")
// 	payload, err := json.MarshalIndent(bot.config, "", " ")

// 	if err != nil {
// 		return err
// 	}

// 	err = ioutil.WriteFile(CONFIG_FILE, payload, 0644)

// 	if err != nil {
// 		return err
// 	}
// 	retrieved, err := ioutil.ReadFile(CONFIG_FILE)
// 	if err != nil {
// 		return err
// 	}

// 	if !bytes.Equal(retrieved, payload) {
// 		return fmt.Errorf("file %q content doesn't match, read %s expected %s\n", CONFIG_FILE, retrieved, payload)
// 	}

// 	return nil
// }
