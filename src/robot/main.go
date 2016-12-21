package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
)

var verbose bool

var version = "No Version Provided"
var buildstamp = "No Buildstamp Provided"
var githash = "No Githash Provided"

func main() {
	var shutdownChan chan bool
	var docker Docker

	log.Println("Version:" + version + "\n" + "Buildstamp:" + buildstamp + "\n" + "Git Hash:" + githash)
	docker.handleInterrupt()

	verbose = true

	log.Println("Loading config file config.json..")
	err := docker.ReadConfig()
	if err != nil {
		log.Printf("Fatal error opening config.json file. Error: %s", err)
		return
	}
	verbose = docker.Verbose

	log.Println("Config file loaded. Checking settings.. ")

	err = docker.CheckConfigValues()
	if err != nil {
		log.Println("Fatal error checking config values. Error:", err)
		return
	}

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
