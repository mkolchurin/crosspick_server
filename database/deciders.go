package database

import "gorm.io/gorm"

type Deciders struct {
	gorm.Model
	Id          uint   `gorm:"primary_key;auto_increment;not_null;type:serial"`
	Title       string `gorm:"uniq"`
	Description string
	Author      string
	CreatedAt   int64 `gorm:"autoCreateTime"`
	Maps        []Maps
}

func CreateDecider(title string, description string, author string, maps []Maps) error {

	decider := Deciders{
		Title:       title,
		Description: description,
		Author:      author,
		Maps:        maps,
	}
	tx := db.Model(&Deciders{}).Preload("Maps").Create(&decider)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetDeciders() ([]Deciders, error) {
	var Deciders []Deciders
	err := db.Find(&Deciders)
	if err.Error != nil {
		return nil, err.Error
	}
	return Deciders, nil
}
