package main

import (
	"time"
)

type ExchangeAPI struct {
	Name                    string
	Type                    ExchangeType
	APIKey                  string
	APISecret               string
	ClientID                string
	Enabled                 bool
	Verbose                 bool
	RESTPollingDelay        time.Duration //ms
	EnabledPairs            string
	AvailablePairs          string
	BaseCurrencies          string
	AuthenticatedAPISupport bool
}
