package api

import (
	"go-gin-example/e"
	"go-gin-example/models"
	"go-gin-example/pkg/logging"
	"go-gin-example/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// @Version 1.0
// @Summary 获取JwtToken
// @Tags 用户模块
// @Produce json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 "{"code":200,"msg":"ok", "data":{}}"
// @Router /auth [get]
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
				logging.Info(data["token"])
			}
		} else {
			code = e.ERROR_AUTH
		}
		msg = e.GetMsg(code)
	} else {
		for _, value := range valid.Errors {
			logging.Info(value.Key, value.Message)
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
