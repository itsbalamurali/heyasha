//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:30 PM
package main

/*import (
	"os"
	"github.com/BurntSushi/toml"
	"log"
)*/

// Info from config file
type Config struct {
	Baseurl   string
	Public    string
	Admin     string
	Metadata  string
	Index     string
}
/*
// Reads info from config file
func ReadConfig() Config {
	var configfile = ""//flag.Configfile
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
*/