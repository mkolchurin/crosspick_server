package database

import (
	"fmt"

	"gorm.io/gorm"
)

type Maps struct {
	gorm.Model
	Name        string `gorm:"uniq"`
	Mode        uint16
	Icon        string
	Description string
	Author      string
	Order       int16
	CreatedAt   int64 `gorm:"autoCreateTime"`
	UploaderId  uint  `gorm:"many2many:users;"`
}

func GetMaps(first int, limit int) ([]Maps, error) {
	var maps []Maps
	tx := db.Table("maps").Where("id >= ?", fmt.Sprintf("%d", first)).Limit(limit).Find(&maps)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maps, nil
}

func InsertMap(name string, mode uint16, icon_path string, description string, author string, uploaderId uint) error {
	if mode < 2 || mode > 8 {
		return fmt.Errorf("invalid mode")
	}
	return InsertMapByStruct(&Maps{
		Name:        name,
		Mode:        mode,
		Icon:        icon_path,
		Description: description,
		Order:       0,
		Author:      author,
		UploaderId:  uploaderId,
	})
}

func InsertMapByStruct(maps *Maps) error {
	tx := db.Table("maps").Create(&maps)
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
