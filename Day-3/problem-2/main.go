package main

import (
	"Freshers_2021/Day-3/problem-2/Config"
	"Freshers_2021/Day-3/problem-2/Models"
	"Freshers_2021/Day-3/problem-2/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

func main() {
	var er error
	Config.DB, er = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if er != nil {
		fmt.Println("Status", er)
	}
	defer Config.DB.Close()
	var models = []interface{}{&Models.Student{}, &Models.SubjectMark{}}
	Config.DB.AutoMigrate(models...)
	routes := Routes.SetupRoutes()
	routes.Run()
}
