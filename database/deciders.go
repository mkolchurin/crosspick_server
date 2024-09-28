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

func DeleteDecider(id uint) error {
	tx := db.Delete(&Deciders{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetDeciderByTitle(title string) (*Deciders, error) {
	var decider Deciders
	tx := db.Where("Title = ?", title).First(&decider)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &decider, nil
}

func GetDecider(id uint) (*Deciders, error) {
	var decider Deciders
	tx := db.Preload("Maps").Where("ID == ?", id).First(&decider)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &decider, nil
}

func GetDeciders() ([]Deciders, error) {
	var deciders []Deciders
	tx := db.Preload("Maps").Find(&deciders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return deciders, nil
}
