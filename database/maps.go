package database

import (
	"fmt"

	"gorm.io/gorm"
)

type Maps struct {
	gorm.Model
	Id          uint   `gorm:"primary_key;auto_increment;not_null;type:serial"`
	Name        string `gorm:"uniq"`
	Mode        uint16
	Icon        string
	Description string
	Author      string
	Order       int16
	CreatedAt   int64 `gorm:"autoCreateTime"`
}

func GetMaps(first int, limit int) ([]Maps, error) {
	var maps []Maps
	tx := db.Table("maps").Where("id >= ?", fmt.Sprintf("%d", first)).Limit(limit).Find(&maps)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maps, nil
}

func InsertMap(name string, mode uint16, icon_path string, description string, author string) error {
	if mode < 2 || mode > 8 {
		return fmt.Errorf("invalid mode")
	}
	tx := db.Table("maps").Create(&Maps{
		Name:        name,
		Mode:        mode,
		Icon:        icon_path,
		Description: description,
		Order:       0,
		Author:      author,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func RemoveMap(name string) error {
	tx := db.Table("maps").Where("name = ?", name).Delete(&Maps{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
