package movie

import (
	"movie/DbModel"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/eefret/gomdb"

	"strings"
)

func SearchMovieByTitle (c *gin.Context){
	movieData:=DbModel.MovieData{}
	DbModel.Body = make(map[string]interface{})
		DbModel.DB.Model(DbModel.MovieData{}).Where("title=?",c.PostForm("title")).Find(&movieData)
		if movieData.ID>0 {

			DbModel.Body["title"]=movieData.Title
			DbModel.Body["Rated"]=movieData.Rated
			DbModel.Body["Year"]=movieData.Year
			DbModel.Body["ImdbID"]=movieData.ImdbID
			DbModel.Body["Genre"]=strings.Split(movieData.Genre, ",")
			c.JSON(http.StatusOK, gin.H{
				"status":  1,
				"body":    DbModel.Body,
			})
		}else{

			query := &gomdb.QueryData{Title: c.PostForm("title")}
			res, err := gomdb.MovieByTitle(query)
			if err != nil {
				fmt.Println(err)

				c.JSON(http.StatusBadRequest, gin.H{
					"status":  0,
					"error":   err,
					"body":    DbModel.Body,
					"message": "Movie not found",
				})
				return
			}

			MovieData:=DbModel.MovieData{
				Title:res.Title,
				Rated:res.Rated,
				Year:res.Year,
				ImdbID:res.ImdbID,
				Genre:res.Genre,
			}


			DbModel.DB.Save(&MovieData)




			DbModel.Body["title"]=res.Title
			DbModel.Body["Rated"]=res.Rated
			DbModel.Body["Year"]=res.Year
			DbModel.Body["ImdbID"]=res.ImdbID
			DbModel.Body["Genre"]= strings.Split(res.Genre, ",")

			c.JSON(http.StatusOK, gin.H{
				"status":  1,
				"body":    DbModel.Body,
			})


		}


}


func SearchMovie (c *gin.Context){
	movieData:=DbModel.MovieData{}
	DbModel.Body = make(map[string]interface{})
	genre:=c.PostForm("genre")
	imdbId:=c.PostForm("imdbId")
	year:=c.PostForm("year")
	//rating:=c.PostForm("rating")
	condition:=""
	if genre!=""{

		condition=condition+"genre like('%"+genre+"%')"
	}

	if year!=""{
		if condition !=""{
			condition=condition+" and "
		}

		condition=condition+" year="+year
	}

	if imdbId!=""{
		if condition !=""{
			condition=condition+" and "
		}

		condition=condition+" imdb_id  like('%"+imdbId+"%')"
	}
fmt.Println(condition)
	DbModel.DB.Where(condition).Find(&movieData)
	fmt.Println(movieData)
	if movieData.ID>0 {

		DbModel.Body["title"]=movieData.Title
		DbModel.Body["Rated"]=movieData.Rated
		DbModel.Body["Year"]=movieData.Year
		DbModel.Body["ImdbID"]=movieData.ImdbID
		DbModel.Body["Genre"]=strings.Split(movieData.Genre, ",")
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"body":    DbModel.Body,
		})
	}else{


			c.JSON(http.StatusBadRequest, gin.H{
				"status":  0,
				"error":   "movie not found",
				"body":    DbModel.Body,
				"message": "Movie not found",
			})
		}



}


func UpdateRating (c *gin.Context){
	movieData:=DbModel.MovieData{}
	imdbId:=c.PostForm("imdbId")

	DbModel.DB.Where("imdb_id=?",imdbId).Find(&movieData)

	if movieData.ID>0 {
		DbModel.DB.Model(DbModel.MovieData{}).Where("imdb_id=?",imdbId).Update("rated",c.PostForm("rated"))


		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "sucessfully Updated",
		})
	}else{


		c.JSON(http.StatusBadRequest, gin.H{
			"status":  0,
			"error":   "movie not found",
			"body":    DbModel.Body,
		})
	}



}


func UpdateGenres (c *gin.Context){
	movieData:=DbModel.MovieData{}
	imdbId:=c.PostForm("imdbId")
fmt.Println(imdbId)
	//rating:=c.PostForm("rating")

	DbModel.DB.Where("imdb_id=?",imdbId).Find(&movieData)
fmt.Println(movieData)
	if movieData.ID>0 {
		DbModel.DB.Model(DbModel.MovieData{}).Where("imdb_id=?",imdbId).Update("genre",c.PostForm("genre"))


		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "sucessfully Updated",
		})
	}else{


		c.JSON(http.StatusBadRequest, gin.H{
			"status":  0,
			"error":   "movie not found",
		})
	}



}

