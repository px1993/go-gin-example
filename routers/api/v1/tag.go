package v1

import (
	"go-gin-example/e"
	"go-gin-example/models"
	setting "go-gin-example/pkg"
	"go-gin-example/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	//将传参赋值给map
	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagsTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//添加标签
func AddTasg(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.Query("state")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空！")
	valid.MaxSize(name, 100, "name").Message("name最长为100个字符")
	valid.Required(state, "state").Message("state不能为空！")
	valid.Range(state, 0, 1, "state").Message("state只能是0和1")
	valid.Required(createdBy, "created_by").Message("createBy不能为空！")
	valid.MaxSize(createdBy, 100, "created_by").Message("createBy最长为100个字符")

	code := e.INVALID_PARAMS
	msg := e.GetMsg(code)
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
		msg = e.GetMsg(code)
	} else {
		for _, value := range valid.Errors {
			msg = value.Error()
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make([]int, 0),
	})
}

//修改标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	ModifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空！")
	valid.Required(name, "name").Message("name不能为空！")
	valid.Required(ModifiedBy, "modified_by").Message("modified_by不能为空！")

	valid.MaxSize(name, 100, "name").Message("name不能超过100个字符")
	valid.MaxSize(ModifiedBy, 100, "modified_by").Message("modified_by不能超过100个字符")

	code := e.INVALID_PARAMS
	msg := e.GetMsg(code)
	if !valid.HasErrors() {
		if models.ExistTagById(id) {
			if !models.ExistTagByName(name) {
				code = e.SUCCESS
				models.EditTag(id, name, ModifiedBy)
			} else {
				code = e.ERROR_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
		msg = e.GetMsg(code)
	} else {
		for _, value := range valid.Errors {
			msg = value.Error()
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make([]int, 0),
	})
}

//删除标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS
	if id > 0 {
		if models.ExistTagById(id) {
			code = e.SUCCESS
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	msg := e.GetMsg(code)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make([]int, 0),
	})
}
