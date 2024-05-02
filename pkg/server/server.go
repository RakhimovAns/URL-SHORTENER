package server

import (
	"github.com/RakhimovAns/URL-SHORTENER/initializers"
	"github.com/RakhimovAns/URL-SHORTENER/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

var R *gin.Engine

func StartServer() {
	R = gin.Default()
	R.POST("/short", GetLink)
	R.Static("/static", "./static")
	R.GET("/", func(c *gin.Context) {
		c.File(filepath.Join("static", "index.html"))
	})
	go LoadEndPoints(R)
	err := R.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("failed to start server")
	}
}
func GetLink(c *gin.Context) {
	var link model.Link
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	link.Link = initializers.Parser(link.Link)
	//resp, err := http.Get(link.Link)
	//if err != nil || resp.StatusCode != http.StatusOK {
	//	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "service unavailable"})
	//	return
	//}
	Hashed := initializers.HashString(link.Link)
	link.Short = Hashed
	if exist, _ := initializers.IsLinkExists(link.Link); exist == false {
		R.GET("/"+link.Short, func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, link.Link)
		})
		initializers.AddLink(link)
	}
	c.JSON(http.StatusOK, gin.H{"link": link.Short})
	return
}

func LoadEndPoints(r *gin.Engine) {
	links, err := initializers.GetAll()
	if err != nil {
		log.Println("failed to get all links from db: ", err)
	}
	for _, link := range links {
		r.GET("/"+link.Short, func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, link.Link)
		})
	}
}
