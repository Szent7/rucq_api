package requester

import (
	"fmt"
	"net/http"
	"rucq_api/crud"
	"rucq_api/data_scheme"

	"github.com/gin-gonic/gin"
)

// Sign in
func Authenticate(c *gin.Context) {
	/*var User data_scheme.User
	if err := c.BindJSON(&User); err != nil {
		return
	}*/
	c.IndentedJSON(http.StatusOK, "hello")

}

// * Sign up
func AddUser(c *gin.Context) {
	var User data_scheme.User
	if err := c.BindJSON(&User); err != nil {
		fmt.Println(err)
		return
	}
	if err := crud.AddUserDB(&User); err != nil {
		fmt.Println(err)
		return
	}
}

// Get last 50 messages
func GetMessages(c *gin.Context) {
	var User data_scheme.User
	if err := c.BindJSON(&User); err != nil {
		return
	}

}

// Send message from user
func SendMessage(c *gin.Context) {
	var Mes data_scheme.MesContainer
	if err := c.BindJSON(&Mes); err != nil {
		fmt.Println(err)
		return
	}
	if err := crud.AddMessageDB(&Mes); err != nil {
		fmt.Println(err)
		return
	}
}

// Add user to room
func ConnectRoom(c *gin.Context) {
	var User data_scheme.User
	if err := c.BindJSON(&User); err != nil {
		return
	}

}

// Create room by user
func CreateRoom(c *gin.Context) {
	var User data_scheme.User
	if err := c.BindJSON(&User); err != nil {
		return
	}

}
