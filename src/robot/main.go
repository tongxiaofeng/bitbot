package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
)

var verbose bool

var version = "No Version Provided"
var buildstamp = "No Buildstamp Provided"
var githash = "No Githash Provided"
var docker *Docker = &Docker{}

func main() {
	var shutdownChan chan bool

	log.Println("Version:" + version + "   Buildstamp:" + buildstamp + "   Git Hash:" + githash)
	handleInterrupt()

	//verbose = true

	log.Println("Loading config file config.json..")
	err := docker.ReadConfig()
	if err != nil {
		log.Printf("Fatal error opening config.json file. Error: %s", err)
		return
	}
	verbose = docker.Verbose

	log.Println("Config file loaded. Checking settings.. ")

	AdjustGoMaxProcs()
	docker.Start()

	log.Println("Press Ctrl-C to quit.")
	<-shutdownChan
}

func AdjustGoMaxProcs() {
	log.Println("Adjusting bot runtime performance..")
	maxProcsEnv := os.Getenv("GOMAXPROCS")
	maxProcs := runtime.NumCPU()
	log.Println("Number of CPU's detected:", maxProcs)

	if maxProcsEnv != "" {
		log.Println("GOMAXPROCS env =", maxProcsEnv)
		env, err := strconv.Atoi(maxProcsEnv)

		if err != nil {
			log.Println("Unable to convert GOMAXPROCS to int, using", maxProcs)
		} else {
			maxProcs = env
		}
	}
	log.Println("Set GOMAXPROCS to:", maxProcs)
	runtime.GOMAXPROCS(maxProcs)
}

func handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("Captured %v.", sig)
		docker.SaveConfig()
		docker.Stop()
		os.Exit(1)
	}()
}
