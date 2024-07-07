package database

import (
	"testing"
)

func TestCreateRole(t *testing.T) {
	err := CreateRole("admin")
	if err != nil {
		t.Error(err)
	}
	t.Log("Role created")
}

func TestCreateUser(t *testing.T) {
	role, err := GetRole("admin")
	if err != nil {
		t.Error(err)
	}
	roles := []Roles{role}
	err = CreateUser("admin", "password", "email", roles)
	if err != nil {
		t.Error(err)
	}
	t.Log("User created with roles:", roles)
}

func TestGetUsers(t *testing.T) {
	users, err := GetUsers()
	if err != nil {
		t.Error(err)
	}
	t.Log(users)
}

func TestGetRoles(t *testing.T) {
	roles, err := GetRoles()
	if err != nil {
		t.Error(err)
	}
	t.Log(roles)
}
