package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"src/services"
)

/*func register(c *gin.Context) {
	userName := c.PostForm("username")
	passWD:=c.PostForm("password")
	mobile:=c.PostForm("mobile")

	u := models.User{UserName: userName, PasswdSha1: passWD,Mobile:mobile}

	ra, err := u.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}*/

func GetUser(c *gin.Context) {
	id := c.GetInt64("id")

	ra := services.UserService.Get(id)

	msg := fmt.Sprintf("get successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": ra,
	})
}
