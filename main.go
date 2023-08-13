package main

import (
	"fmt"
	"net/http"
	"time"
	controllers "web-service-gin/controller"
	"web-service-gin/middlewares"
	"web-service-gin/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// album struct
type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

// albums slice
var albums = []album{
	{ID: "1", Title: "Greatest Hits 2Pac", Artist: "Tupac Shakur", Price: 12.99},
	{ID: "2", Title: "Hans Zimmer Masterpieces", Artist: "Hans Zimmer", Price: 13.99},
	{ID: "3", Title: "Chill Out Classics", Artist: "Various Artists", Price: 15.99},
}

func main() {

	models.ConnectDataBase()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization"},
		AllowCredentials: true,
	}))

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.Use(LoggerMiddleWare)
	router.GET("albums/:id", getAlbumById)
	router.POST("albums", postAlbum)
	router.DELETE("albums/:id", removeAlbumById)

	// all routes in the protected group will use the jwt middleware
	protected := router.Group("/api")

	protected.Use(middlewares.JwtAuthMiddleware())
	// protected.Use(token.CheckTokenExpiration())
	protected.GET("/albums", getAllAlbums)
	protected.GET("/user", controllers.CurrentUser)

	router.Run("localhost:8080")

}

func LoggerMiddleWare(c *gin.Context) {
	fmt.Println("Request has been received at", time.Now())
	c.Next()

	fmt.Println("Request handled")
}

// get all albums
func getAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// get album by id
func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return

		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// remove album by id
// we append the index upto but not including the albums[:index] with albums[index+1:] which is from albums index +1 to the end of the slice
func removeAlbumById(c *gin.Context) {
	id := c.Param("id")
	for index, album := range albums {
		if album.ID != id {
			albums = append(albums[:index], albums[index+1:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

// create new album
func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)

}
