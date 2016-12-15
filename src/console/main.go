package main

import (
	//"errors"
	"log"
	//"os"
	//"os/signal"
	//"runtime"
	//"strconv"
	//"syscall"
)

// type Exchange struct {
// 	btcc        BTCC
// 	okcoinChina OKCoin
// 	okcoinIntl  OKCoin
// 	huobi       HUOBI
// }

// type Bot struct {
// 	config    Config
// 	exchange  Exchange
// 	exchanges []IBotExchange
// 	shutdown  chan bool
// }

//var bot Bot

// func setupBotExchanges() {
// 	for _, exch := range bot.config.Exchanges {
// 		for i := 0; i < len(bot.exchanges); i++ {
// 			if bot.exchanges[i] != nil {
// 				if bot.exchanges[i].GetName() == exch.Name {
// 					bot.exchanges[i].Setup(exch)
// 					if bot.exchanges[i].IsEnabled() {
// 						log.Printf("%s: Exchange support: %s (Authenticated API support: %s - Verbose mode: %s).\n", exch.Name, IsEnabled(exch.Enabled), IsEnabled(exch.AuthenticatedAPISupport), IsEnabled(exch.Verbose))
// 						bot.exchanges[i].Start()
// 					} else {
// 						log.Printf("%s: Exchange support: %s\n", exch.Name, IsEnabled(exch.Enabled))
// 					}
// 				}
// 			}
// 		}
// 	}
// }

func main() {
	//HandleInterrupt()
	log.Println("Loading config file config.json..")

	//err := errors.New("")
	//bot.config, err = ReadConfig()
	// if err != nil {
	// 	log.Printf("Fatal error opening config.json file. Error: %s", err)
	// 	return
	// }
	// log.Println("Config file loaded. Checking settings.. ")

	// err = CheckExchangeConfigValues()
	// if err != nil {
	// 	log.Println("Fatal error checking config values. Error:", err)
	// 	return
	// }

	//log.Printf("Bot '%s' started.\n", bot.config.Name)
	//AdjustGoMaxProcs()

	// log.Printf("Available Exchanges: %d. Enabled Exchanges: %d.\n", len(bot.config.Exchanges), GetEnabledExchanges())
	// log.Println("Bot Exchange support:")

	// bot.exchanges = []IBotExchange{
	// 	new(BTCC),
	// 	new(OKCoin),
	// 	new(OKCoin),
	// 	new(HUOBI),
	// }

	// for i := 0; i < len(bot.exchanges); i++ {
	// 	if bot.exchanges[i] != nil {
	// 		bot.exchanges[i].SetDefaults()
	// 		log.Printf("Exchange %s successfully set default settings.\n", bot.exchanges[i].GetName())
	// 	}
	// }

	// setupBotExchanges()

	//err = RetrieveConfigCurrencyPairs(bot.config)

	// if err != nil {
	// 	log.Println("Fatal error retrieving config currency AvailablePairs. Error: ", err)
	// }

	//<-bot.shutdown
	//Shutdown()
}

// func AdjustGoMaxProcs() {
// 	log.Println("Adjusting bot runtime performance..")
// 	maxProcsEnv := os.Getenv("GOMAXPROCS")
// 	maxProcs := runtime.NumCPU()
// 	log.Println("Number of CPU's detected:", maxProcs)

// 	if maxProcsEnv != "" {
// 		log.Println("GOMAXPROCS env =", maxProcsEnv)
// 		env, err := strconv.Atoi(maxProcsEnv)

// 		if err != nil {
// 			log.Println("Unable to convert GOMAXPROCS to int, using", maxProcs)
// 		} else {
// 			maxProcs = env
// 		}
// 	}
// 	log.Println("Set GOMAXPROCS to:", maxProcs)
// 	runtime.GOMAXPROCS(maxProcs)
// }

// func HandleInterrupt() {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
// 	go func() {
// 		sig := <-c
// 		log.Printf("Captured %v.", sig)
// 		Shutdown()
// 	}()
// }

// func Shutdown() {
// 	log.Println("Bot shutting down..")
// 	err := SaveConfig()

// 	if err != nil {
// 		log.Println("Unable to save config.")
// 	} else {
// 		log.Println("Config file saved successfully.")
// 	}

// 	log.Println("Exiting.")
// 	os.Exit(1)
// }
