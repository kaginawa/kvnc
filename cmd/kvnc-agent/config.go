package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kaginawa/kvnc"
)

type kaginawaConfig struct {
	APIKey        string `json:"api_key"`
	Server        string `json:"server"`
	CustomID      string `json:"custom_id"`
	SSHLocalPort  int    `json:"ssh_local_port"`
	UpdateEnabled bool   `json:"update_enabled"`
}

// loadConfig loads a kaginawa configuration file from specified path.
func loadConfig(path string) (kaginawaConfig, error) {
	c := kaginawaConfig{
		SSHLocalPort:  5900,
		UpdateEnabled: false,
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return c, fmt.Errorf("failed to load %s: %w", path, err)
	}
	if err := json.Unmarshal(content, &c); err != nil {
		return c, fmt.Errorf("failed to unmarshal %s: %w", path, err)
	}
	return c, nil
}

// save saves a kaginawa configuration file to specified path.
func (c kaginawaConfig) save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", path, err)
	}
	defer kvnc.SafeClose(f, path)
	content, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		return err
	}
	return nil
}
