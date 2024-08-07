package database

import (
	"testing"
)

func TestCreate(t *testing.T) {
	maps, err := GetMaps(0, -1)
	if err != nil {
		t.Error(err)
	}
	user, err := GetUser("admin")
	if err != nil {
		t.Error(err)
	}
	err = CreateDecider("sample", "description", user.ID, maps)
	if err != nil {
		t.Error(err)
	}
	t.Log("Decider created with maps:", maps)
}

func TestGetDeciders(t *testing.T) {
	deciders, err := GetDeciders()
	if err != nil {
		t.Error(err)
	}
	t.Log(deciders)
}
