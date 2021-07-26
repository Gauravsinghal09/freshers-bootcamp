package Models

import (
	"Freshers_2021/Day-3/problem-2/Config"
	"fmt"
	"gorm.io/gorm/clause"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]Student) (err error) {

	if err = Config.DB.Preload(clause.Associations).Find(user).Error; err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *Student) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *Student, id string) (err error) {
	if err = Config.DB.Preload(clause.Associations).Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *Student, id string) (err error) {
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
