package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	ftpserverlib "github.com/fclairamb/ftpserverlib"
	"github.com/spf13/afero"
)

type UserConfig struct {
	Path     string `json:"path"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type FTPDriver struct {
	settings *ftpserverlib.Settings
	users    map[string]UserConfig
}

func (d *FTPDriver) GetSettings() (*ftpserverlib.Settings, error) {
	d.settings.ListenAddr = "0.0.0.0:2121"
	d.settings.PassiveTransferPortRange = &ftpserverlib.PortRange{Start: 2122, End: 2130}
	return d.settings, nil
}

func (d *FTPDriver) ClientConnected(cc ftpserverlib.ClientContext) (string, error) {
	return "Connection established!", nil
}

func (d *FTPDriver) ClientDisconnected(cc ftpserverlib.ClientContext) {
	log.Printf("User disconnected: %s", cc.RemoteAddr())
}

func (d *FTPDriver) AuthUser(cc ftpserverlib.ClientContext, user, pass string) (ftpserverlib.ClientDriver, error) {
	if uCfg, ok := d.users[user]; ok && uCfg.Password == pass {
		return afero.NewBasePathFs(afero.NewOsFs(), uCfg.Path), nil
	}
	return nil, fmt.Errorf("Authentication failed")
}

func (d *FTPDriver) GetTLSConfig() (*tls.Config, error) {
	return &tls.Config{
		InsecureSkipVerify: true,
	}, nil
}

func main() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Could not read config file: %v", err)
	}

	var configs []UserConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		log.Fatalf("Could not parse config: %v", err)
	}

	userMap := make(map[string]UserConfig)
	for _, cfg := range configs {
		userMap[cfg.User] = cfg
	}

	driver := &FTPDriver{
		settings: &ftpserverlib.Settings{ListenAddr: "0.0.0.0:2121"},
		users:    userMap,
	}

	server := ftpserverlib.NewFtpServer(driver)
	log.Printf("Starting FTP server: %s", driver.settings.ListenAddr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start FTP server: %v", err)
	}
}
