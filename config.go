package kvnc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configFileTemplate = `Server %s
APIKey %s
CustomID %s
ViewerCmd %s
ServerCmd %s
Password %s
`

// Config defines configuration file entries.
type Config struct {
	Server    string `json:"server,omitempty"`
	APIKey    string `json:"api_key,omitempty"`    // client-only
	CustomID  string `json:"custom_id,omitempty"`  // client-only
	ViewerCmd string `json:"viewer_cmd,omitempty"` // client-only
	ServerCmd string `json:"server_cmd,omitempty"` // agent-only
	Password  string `json:"password,omitempty"`   // agent-only
}

// LoadConfig loads a configuration file from specified path.
func LoadConfig(path string) (Config, error) {
	var c Config
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return c, fmt.Errorf("failed to load %s: %w", path, err)
	}
	if err := json.Unmarshal(content, &c); err != nil {
		return c, fmt.Errorf("failed to unmarshal %s: %w", path, err)
	}
	return c, nil
}

// Save saves a configuration file to specified path.
func (c Config) Save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", path, err)
	}
	defer SafeClose(f, path)
	content, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		return err
	}
	return nil
}
