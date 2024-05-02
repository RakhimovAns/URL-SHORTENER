package main

import (
	"github.com/RakhimovAns/URL-SHORTENER/initializers"
	"github.com/RakhimovAns/URL-SHORTENER/pkg/server"
)

func init() {
	initializers.ConnectToDB()
	initializers.CreateTable()
	server.StartServer()
}

func main() {

}
