package DbModel

import (
"github.com/jinzhu/gorm"
"github.com/gin-gonic/gin"

)



const DATABASEUSERNAME = "root"
const DATABASEPASSWORD = "root"
const DATABASENAME = "movie_data"
const DBSERVERIP="127.0.0.1"
const PORT="3306"


var Result gin.H


const Encryption_key= "y/B?E(H+MbQeThVmYq3t6w9z$C&F)J@N"


var DB *gorm.DB




func AutoMigrateDb(){


	DB.AutoMigrate(&MovieData{})

}
var Body map[string]interface{}



type MovieData struct{
	gorm.Model
	Title  string
	Year   string
	ImdbID string
	Type   string
	Rated  string
	Genre  string
}


