package models

import "github.com/jinzhu/gorm"

type Aiml struct {
	gorm.Model
	Aiml	string
	Pattern string
	Thatpattern string
	Template string
	Topic string
	Filename string
}

type Personality  struct{
	gorm.Model
	Name string
	Value string
}

type SraiLookup struct {
	gorm.Model
	Pattern string
	TemplateID int64
}

type Wordcensor struct {
	CensorID int64
	WordToCensor string
	ReplaceWith	string
	BotExclude	string
}