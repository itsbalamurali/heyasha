//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 17/06/2016 1:47 AM
package main

import (
	"fmt"
	"github.com/itsbalamurali/heyasha/core/database"
	"github.com/itsbalamurali/heyasha/models"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
)

func main() {

	app := cli.NewApp()
	app.Name = "Asha Cli Tool for Intent and Knowledge Management"
	app.Usage = "asha-cli "
	app.Version = "v0.0.1-alpha"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Balamurali Pandranki",
			Email: "balamurali@live.com",
		},
	}
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:"newintent",
			Usage:  "newintent --name [intent_name] --sentence [sample sentence] --domain [domain id]",
			Action: createNewIntent,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "Intent Name",
				},
				cli.StringFlag{
					Name:  "sentence, s",
					Usage: "Sample sentence for traning",
				},
				cli.IntFlag{
					Name:  "domain, d",
					Usage: "Domain ID",
				},
			},
		},
		{
			Name:   "migratedb",
			Usage:  "migratedb",
			Action: migrateDatabase,
		},
	}

	app.Run(os.Args)
}

func createNewIntent(c *cli.Context) {
	db := database.MysqlCon()
	dbfile := &models.Intent{
		Sentence: c.String("sentence"),
		Intent:   c.String("name"),
		DomainID: uint64(c.Int("domain")),
	}
	db.Save(&dbfile)
	fmt.Println("Success: Created Intent Successfully with id: "+ strconv.Itoa(int(dbfile.ID)))
}

func migrateDatabase(c *cli.Context) {
	db := database.MysqlCon()
	db.AutoMigrate(&models.User{}, &models.ConversationLog{}, &models.Intent{}, &models.Aiml{}, &models.Personality{}, &models.SraiLookup{}, &models.Wordcensor{}, &models.Session{}, &models.File{})
	fmt.Println("Success: Migrated All models!")
}
