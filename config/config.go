//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:30 PM
package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	ps "github.com/itsbalamurali/heyasha/core/engine/stt"
	"path/filepath"
)

// Info from config file
type Config struct {
	Environment  string
	Databases    map[string]database  `toml:"database"`
	Email        email                `toml:"email"`
	Sms          sms                  `toml:"sms"`
	GCloud       map[string]gcloud    `toml:"google_cloud"`
	Pocketsphinx map[string]ps.Config `toml:"pocketsphinx"`
}

type database struct {
	Name     string `toml:"dbname"`
	Host     string `toml:"dbhost"`
	Port     int    `toml:"dbport"`
	User     string `toml:"dbuser"`
	Password string `toml:"dbpassword"`
}

type email struct {
}

type sms struct {
}

type gcloud struct {
	Bucket string `toml:"bucket_name"`
}

// Reads info from config file
func LoadConfig() Config {
	configfile, err := filepath.Abs("config.toml") //flag.Configfile
	if err != nil {
		log.Fatalln("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatalln(err)
	}
	log.Infoln("Loaded configuration file: config.toml")
	return config
}
