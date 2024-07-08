package database

import "gorm.io/gorm"

type Deciders struct {
	gorm.Model
	Title       string `gorm:"uniq"`
	Description string
	CreatedAt   int64  `gorm:"autoCreateTime"`
	Maps        []Maps `gorm:"many2many:deciders_maps;"`
	CreatorId   uint   `gorm:"many2many:users;"`
}

func CreateDecider(title string, description string, creatorId uint, maps []Maps) error {

	decider := Deciders{
		Title:       title,
		Description: description,
		CreatorId:   creatorId,
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
