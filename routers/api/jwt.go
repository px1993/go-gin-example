package api

import (
	"go-gin-example/e"
	"go-gin-example/models"
	"go-gin-example/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//获取token
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	data := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空！")
	valid.Required(password, "password").Message("用户密码不能为空！")

	code := e.INVALID_PARAMS
	msg := e.GetMsg(code)
	if !valid.HasErrors() {
		checkRes := models.CheckAuth(username, password)

		if checkRes {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				code = e.SUCCESS
				data["token"] = token
			}
		} else {
			code = e.ERROR_AUTH
		}
		msg = e.GetMsg(code)
	} else {
		for _, value := range valid.Errors {
			msg = value.Error()
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}