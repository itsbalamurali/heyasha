//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:27 PM
package database

import (
	"os"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/itsbalamurali/heyasha/models"
)

var (
	//Mysql Connection
	Db *gorm.DB
)

const (
	// MySQLDBUrl is the default Mysql url that will be used to connect to the
	// database.
	MySQLDBUrl = "admin_twa:Aydim4VdBK@tcp(128.199.81.183:3306)/admin_twa?charset=utf8&parseTime=True"
)

// Connect connects to mysql
func MysqlCon() *gorm.DB {
	uri := os.Getenv("MYSQLDB_URL")
	if len(uri) == 0 {
		uri = MySQLDBUrl
	}
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Printf("Can't connect to MySQL, go error %v\n", err)
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{},&models.ConversationLog{},&models.Intent{},&models.Aiml{},&models.Personality{},&models.SraiLookup{},&models.Wordcensor{},&models.Session{})
	log.Println("Connected to MySQL Server")
	Db = db
	return Db
}
