package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"src/models"
	"src/services"
	"strconv"
)

// @log in
// @Summary login
// @Description check user passwd
// @Tags user
// @Accept json
// @Security Bearer
// @Produce  json
// @Resource Name
// @Success 200 {object} models.User
// @Router /users/signin [post]
func SignIn(c *gin.Context) {
	var userJson models.User
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

	msg := ""
	userMap := map[string]interface{}{"user_name": userJson.UserName}
	//fmt.Println("userMap: "+*userMap)
	user := services.UserService.GetUserByOpt(userMap)

	fmt.Println("user.PasswdSha1: ", user.PasswdSha1, ",userJson.PasswdSha1: ", userJson.PasswdSha1)
	//fmt.Println(user)

	if user.PasswdSha1 == userJson.PasswdSha1 {
		msg = "sign in ok"
	} else {
		msg = "password error"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// @register
// @Summary adduser
// @Description add user
// @Tags user
// @Accept json
// @Security Bearer
// @Produce  json
// @Param  user body models.User true "Add user"
// @Success 200 {integer} integer 0
// @Router /users/register [post]
func Register(c *gin.Context) {
	var userJson models.User
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := &models.User{}
	user.UserName = userJson.UserName
	user.PasswdSha1 = userJson.PasswdSha1
	user.Mobile = userJson.Mobile

	msg := ""

	ra := services.UserService.Post(user)

	msg = fmt.Sprintf("post successful")
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": ra,
	})
}

// @get one user by id
// @Summary getuser
// @Description get user
// @Tags user
// @Accept json
// @Security Bearer
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {integer} integer 0
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	//id := c.GetInt64("id")
	idStr := c.Param("id")
	msg := ""
	if idStr == "" {
		msg = fmt.Sprintf("get param id is error")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		msg = fmt.Sprintf("param string convert to int64 failed")
	}
	ra := services.UserService.Get(id)

	msg = fmt.Sprintf("get successful")
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": ra,
	})
}

// @put one user by id
// @Summary PutUser
// @Description put user
// @Tags user
// @Accept json
// @Security Bearer
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {integer} integer 0
// @Router /users/{id} [put]
func PutUser(c *gin.Context) {
	idStr := c.Param("id")
	msg := ""
	if idStr == "" {
		msg = fmt.Sprintf("get param id is error")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		msg = fmt.Sprintf("param string convert to int64 failed")
	}

	user := &models.User{}
	user.Id = id
	user.UserName = c.PostForm("username")
	user.PasswdSha1 = c.PostForm("password")
	user.Mobile = c.PostForm("mobile")

	ra := services.UserService.Put(user)

	msg = fmt.Sprintf("put successful")
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": ra,
	})
}

// @delete one user by id
// @Summary DelUser
// @Description delete user
// @Tags user
// @Accept json
// @Security Bearer
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {integer} integer 0
// @Router /users/{id} [delete]
func DelUser(c *gin.Context) {
	idStr := c.Param("id")
	msg := ""
	if idStr == "" {
		msg = fmt.Sprintf("get param id is error")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		msg = fmt.Sprintf("param string convert to int64 failed")
	}
	ra := services.UserService.Delete(id)

	msg = fmt.Sprintf("delete successful")
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": ra,
	})
}

// @get users
// @Summary GetUsers
// @Description get users
// @Tags user
// @Accept json
// @Security Bearer
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	msg := ""

	pageStr, isPage := c.GetQuery("page")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		msg = fmt.Sprintf("param string convert to int64 failed")

	}
	perpageStr, isPerPage := c.GetQuery("perpage")
	perpage, err := strconv.ParseInt(perpageStr, 10, 64)
	if err != nil {
		msg = fmt.Sprintf("param string convert to int64 failed")
	}
	ra := &[]models.User{}
	if isPage && isPerPage {
		offset := (page - 1) * perpage
		ra = services.UserService.GetUsersPaged(offset, perpage)
		msg = fmt.Sprintf("get users successful")
		c.JSON(http.StatusOK, gin.H{
			"msg":     msg,
			"data":    ra,
			"page":    page,
			"perPage": perpage,
		})
	} else {
		ra = services.UserService.GetUsers()
		c.JSON(http.StatusOK, gin.H{
			"msg":  msg,
			"data": ra,
		})
	}

}
