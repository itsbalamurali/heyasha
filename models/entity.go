package models

import (
	"github.com/jinzhu/gorm"
)

type EntityValue struct {
	ID int
	Value string
	Expressions string
	Metadata string
}

type Entity struct  {
	gorm.Model
	EId string
	Doc string
	Values []EntityValue
	Lang string
	Lookup string
}


