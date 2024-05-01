package main

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/RakhimovAns/URL-SHORTENER/cmd/initializers"
	"github.com/RakhimovAns/URL-SHORTENER/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	initializers.ConnectToDB()
	initializers.CreateTable()
}
func main() {
	r := gin.Default()
	r.GET("/short", GetLink)
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("failed to start server")
	}
}

func GetLink(c *gin.Context) {
	var link model.Link
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	Hashed := HashString(link.Link)
	link.Short = Hashed
	if exist, _ := initializers.IsLinkExists(link.Link); exist == false {
		initializers.AddLink(link)
	}
	c.JSON(http.StatusOK, gin.H{"link": link.Short})
	return
	//c.JSON(http.StatusOK, gin.H{"link": "https://localhost:8080/" + Hashed})
}

func HashString(link string) string {
	data := []byte(link)
	hash := md5.Sum(data)
	hashBase64 := base64.StdEncoding.EncodeToString(hash[:])
	return hashBase64[:7]
}
