package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kaginawa/kvnc"
)

// clientConfig defines configuration file entries.
type clientConfig struct {
	Server    string `json:"server,omitempty"`
	APIKey    string `json:"api_key,omitempty"`
	CustomID  string `json:"custom_id,omitempty"`
	ViewerCmd string `json:"viewer_cmd,omitempty"`
}

// loadConfig loads a configuration file from specified path.
func loadConfig(path string) (clientConfig, error) {
	var c clientConfig
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return c, fmt.Errorf("failed to load %s: %w", path, err)
	}
	if err := json.Unmarshal(content, &c); err != nil {
		return c, fmt.Errorf("failed to unmarshal %s: %w", path, err)
	}
	return c, nil
}

// save saves a configuration file to specified path.
func (c clientConfig) save(path string) error {
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
