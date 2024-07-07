package database

import (
	"log"
	"testing"

	"github.com/mkolchurin/crosspick_server/appconfig"
)

func init() {
	cfg, err := appconfig.GetConfig("../config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	Connect(cfg)
	db.Table("maps").AutoMigrate(&Maps{})
}

func TestGetMaps(t *testing.T) {
	maps, err := GetMaps(0, 1)
	if err != nil {
		t.Error(err)
	}
	if len(maps) != 1 {
		t.Error("Expected 1 map, got", len(maps))
	}
	t.Log(maps)
	maps, err = GetMaps(1, 3)
	if err != nil {
		t.Error(err)
	}
	if len(maps) != 3 {
		t.Error("Expected 3 maps, got", len(maps))
	}
	t.Log(maps)
}

func TestInsertMap(t *testing.T) {
	err := InsertMap("test", 2, "test", "test", "test")
	if err != nil {
		t.Error(err)
	}
	// RemoveMap("test")
}

func TestRemoveMap(t *testing.T) {
	err := RemoveMap("test")
	if err != nil {
		t.Error(err)
	}
}
