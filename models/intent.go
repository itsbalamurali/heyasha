package models

import "github.com/jinzhu/gorm"

type Intent struct {
	gorm.Model
	Sentence string
	Intent   string
	DomainID string
}