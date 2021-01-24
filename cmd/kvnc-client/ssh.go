package main

import (
	"strings"

	"github.com/kaginawa/kaginawa-sdk-go"
	"golang.org/x/crypto/ssh"
)

func newKaginawaClient() (*kaginawa.Client, error) {
	endpoint := config.Server
	if !strings.HasPrefix(endpoint, "http://") && !strings.HasPrefix(endpoint, "https://") {
		if config.Server == "localhost" {
			endpoint = "http://" + endpoint
		} else {
			endpoint = "https://" + endpoint
		}
	}
	return kaginawa.NewClient(endpoint, config.APIKey)
}

func createSSHConfig(user, key, password string) (*ssh.ClientConfig, error) {
	config := ssh.ClientConfig{
		User:            user,
		Auth:            make([]ssh.AuthMethod, 0),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if len(key) > 0 {
		parsed, err := ssh.ParsePrivateKey([]byte(key))
		if err != nil {
			return nil, err
		}
		config.Auth = append(config.Auth, ssh.PublicKeys(parsed))
	}
	if len(password) > 0 {
		config.Auth = append(config.Auth, ssh.Password(password))
	}
	return &config, nil
}
