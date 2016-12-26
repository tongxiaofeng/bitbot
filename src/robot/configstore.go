package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	CONFIG_FILE = "config.json"
)

type IConfigStore interface {
	Read() (DockerConfig, error)
	Write(DockerConfig) error
}

type FileConfigStore struct {
}

func (store FileConfigStore) Read() (DockerConfig, error) {
	var dockerConfig DockerConfig
	file, err := ioutil.ReadFile(CONFIG_FILE)

	if err != nil {
		return DockerConfig{}, err
	}

	err = json.Unmarshal(file, &dockerConfig)
	return dockerConfig, err
}

func (store FileConfigStore) Write(dockerConfig DockerConfig) (err error) {
	log.Println("Saving config to file...")
	payload, err := json.MarshalIndent(dockerConfig, "", "  ")

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
