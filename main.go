package main

import (
	"fmt"
	"rucq_api/requester"

	"github.com/gin-gonic/gin"
)

func main() {
	//crud.InitDB()
	fmt.Println(requester.GeneratePassword(10))
	//go requester.StartWebSocketServer()
	r := gin.Default()
	r.GET("/auth", requester.Authenticate)
	r.POST("/addUser", requester.AddUser)

	r.GET("/getMessages", requester.GetMessages)
	r.POST("/sendMessage", requester.SendMessage)

	r.POST("/connectRoom", requester.ConnectRoom)
	r.POST("/createRoom", requester.CreateRoom)

	r.Run(":10015")
}
