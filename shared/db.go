//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:27 PM

package shared

import "gopkg.in/mgo.v2"

//mongodb
type Db struct {
	Url        string
	MgoSession *mgo.Session
}

//get new db session
func (d *Db) GetSession() *mgo.Session {
	if d.MgoSession == nil {
		var err error
		d.MgoSession, err = mgo.Dial(d.Url)
		if err != nil {
			panic(err) // no, not really
		}
	}
	return d.MgoSession.Clone()

}

