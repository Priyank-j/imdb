package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"movie/DbModel"
	"movie/app"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

		// Set Gin to production mode



		db, err := gorm.Open("mysql", DbModel.DATABASEUSERNAME+":"+DbModel.DATABASEPASSWORD+"@tcp("+DbModel.DBSERVERIP+":"+DbModel.PORT+")/"+DbModel.DATABASENAME+"?charset=utf8&parseTime=True&loc=Local")
		if err!=nil{
			log.Println("errrrr",err)
		}
		defer db.Close()
		DbModel.DB=db
		DbModel.AutoMigrateDb()
		// Initialize the routes
		app.InitializeRoutes()



}