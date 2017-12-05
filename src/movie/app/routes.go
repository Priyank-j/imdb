
package app

import (
"github.com/gin-gonic/gin"
//"cccServer/middlewares"
"os"
"github.com/jinzhu/gorm"
//"cccServer/users"
//"github.com/gin-gonic/contrib/sessions"

	"movie/movie"
)
var DB *gorm.DB

const (
	// Port at which the server starts listening
	Port = "8088"
)

var router *gin.Engine

func InitializeRoutes() {
	//gin.SetMode(gin.ReleaseMode)

	router = gin.Default()




	router.Use(func(c *gin.Context) {
		//c.Next()
		c.Writer.Header().Set("Cache-Control", "no-cache")
	})

	router.POST("/SearchMovieByTitle",movie.SearchMovieByTitle)

	router.POST("/SearchMovie",movie.SearchMovie)


	router.POST("/UpdateRating",movie.UpdateRating)
	router.POST("/UpdateGenres",movie.UpdateGenres)

	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)

}






