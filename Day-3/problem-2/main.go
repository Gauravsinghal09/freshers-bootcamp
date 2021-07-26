package main

import (
	"Freshers_2021/Day-3/problem-2/Config"
	"Freshers_2021/Day-3/problem-2/Models"
	"Freshers_2021/Day-3/problem-2/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	var er error
	//Config.DB, er = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB, er = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if er != nil {
		fmt.Println("Status", er)
	}
	//defer Config.DB.Close()
	//var models = []interface{}{&Models.Student{}, &Models.SubjectMark{}}

	Config.DB.AutoMigrate(&Models.Student{})
	Config.DB.AutoMigrate(&Models.SubjectMark{})
	routes := Routes.SetupRoutes()
	routes.Run()
}
