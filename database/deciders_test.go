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
	err = db.AutoMigrate(&Deciders{})
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	maps, err := GetMaps(0, 0)
	if err != nil {
		t.Error(err)
	}
	err = CreateDecider("sample", "description", "author", maps)
	if err != nil {
		t.Error(err)
	}
}

func TestGetDeciders(t *testing.T) {
	deciders, err := GetDeciders()
	if err != nil {
		t.Error(err)
	}
	t.Log(deciders)
}
