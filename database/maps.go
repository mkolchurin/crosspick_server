package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type DbDeciderMap struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Mode        int
	Icon        string
	Description string
	Author      string
	CreatedAt   int64 `gorm:"autoCreateTime"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = Connect()
	if err != nil {
		log.Fatal(err)
	}
}

func GetMaps(first int, limit int) ([]DbDeciderMap, error) {
	var maps []DbDeciderMap
	tx := db.Table("maps").Where("id >= ?", fmt.Sprintf("%d", first)).Limit(limit).Find(&maps)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maps, nil
}

func InsertMap(name string, mode int, icon_path string, description string, author string) error {
	if mode < 2 || mode > 8 {
		return fmt.Errorf("invalid mode")
	}
	Map := DbDeciderMap{
		Name:        name,
		Mode:        mode,
		Icon:        icon_path,
		Description: description,
		Author:      author,
	}
	tx := db.Table("maps").Create(&Map)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func RemoveMap(name string) error {
	tx := db.Table("maps").Where("name = ?", name).Delete(&DbDeciderMap{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
