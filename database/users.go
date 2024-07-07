package database

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id       uint   `gorm:"primary_key;auto_increment;not_null;type:serial"`
	Username string `gorm:"uniq"`
	Password string
	Email    string
	Roles    []Roles `gorm:"many2many:user_roles;"`
}

type Roles struct {
	gorm.Model
	Id   uint   `gorm:"primary_key;auto_increment;not_null;type:serial"`
	Name string `gorm:"uniq"`
}

func CreateUser(username string, password string, email string, roles []Roles) error {
	user := Users{
		Username: username,
		Password: password,
		Email:    email,
		Roles:    roles,
	}
	tx := db.Model(&Users{}).Preload("Roles").Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetUsers() ([]Users, error) {
	var users []Users
	err := db.Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}
	return users, nil
}

func GetRoles() ([]Roles, error) {
	var roles []Roles
	err := db.Find(&roles)
	if err.Error != nil {
		return nil, err.Error
	}
	return roles, nil
}

func CreateRole(name string) error {
	role := Roles{
		Name: name,
	}
	tx := db.Model(&Roles{}).Create(&role)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func RemoveRole(name string) error {
	tx := db.Table("roles").Where("name = ?", name).Delete(&Roles{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func RemoveUser(username string) error {
	tx := db.Table("users").Where("username = ?", username).Delete(&Users{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetUser(username string) (Users, error) {
	var user Users
	tx := db.Preload("Roles").Where("username = ?", username).First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func GetRole(name string) (Roles, error) {
	var role Roles
	tx := db.Where("name = ?", name).First(&role)
	if tx.Error != nil {
		return role, tx.Error
	}
	return role, nil
}

func AddRoleToUser(username string, roleName string) error {

	user, err := GetUser(username)
	if err != nil {
		return err
	}
	var role Roles
	role, err = GetRole(roleName)
	if err != nil {
		return err
	}
	return db.Model(&user).Association("Roles").Append(&role)
}

func RemoveRoleFromUser(username string, roleName string) error {
	user, err := GetUser(username)
	if err != nil {
		return err
	}
	var role Roles
	role, err = GetRole(roleName)
	if err != nil {
		return err
	}
	return db.Model(&user).Association("Roles").Delete(&role)
}

func AddRoleToUserById(userId uint, roleId uint) error {
	var user Users
	var role Roles
	user = Users{Id: userId}
	role = Roles{Id: roleId}
	return db.Model(&user).Association("Roles").Append(&role)
}

func RemoveRoleFromUserById(userId uint, roleId uint) error {
	var user Users
	var role Roles
	user = Users{Id: userId}
	role = Roles{Id: roleId}
	return db.Model(&user).Association("Roles").Delete(&role)
}
