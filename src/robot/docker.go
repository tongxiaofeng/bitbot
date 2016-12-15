package bitbot

import (
	"time"
)

type DockerStatus int

const (
	DOCKER_CONNECTED DockerStatus = 1 + iota
	DOCKER_DISCONNECTED
	DOCKER_CONNECTING
	DOCKER_DISCONNECTING
)

type Docker struct {
	ID            int
	IP            string
	Port          int
	OS            string
	OSVersion     string
	DockerVersion string
	Status        DockerStatus
	LastHeartBeat time.Time
	//IsLocal       bool // Is it running on the same machine as server?
}

//Functions

/*
	RunBot ([]ExchangeAPI,Strategy, StrategyConfigs) botid,error
	StopBot (botid) error

	Login()
	Logout()

	Start()

	var config Config
	SendConfig()

	//ModifyStrategyCode
	//ModifyStrategyVaraibles
*/
