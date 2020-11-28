package v1

import (
	"fmt"
	"go-gin-example/e"
	"go-gin-example/models"
	setting "go-gin-example/pkg"
	"go-gin-example/util"
	"net/http"

	"github.com/astaxie/beego/validation"

	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

//获取所有文章
func GetArticles(c *gin.Context) {
	title := c.Query("title")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if title != "" {
		maps["title"] = title
	}

	code := e.INVALID_PARAMS
	msg := e.GetMsg(code)

	data["count"] = models.GetArticlesTotal(maps)
	data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})

}

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS
	var data interface{}
	if id > 0 {
		if models.ExistsArticleById(id) {
			code = e.SUCCESS
			data = models.GetArticleById(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//添加文章
func AddArticle(c *gin.Context) {
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	state := com.StrTo(c.Query("state")).MustInt()

	code := e.INVALID_PARAMS
	msg := e.GetMsg(code)
	valid := validation.Validation{}

	valid.Required(title, "title").Message("文章标题[title]不能为空！")
	valid.MaxSize(title, 100, "title").Message("文章标题[title]不能超过100个字符")
	valid.Required(desc, "desc").Message("文章简介[desc]不能为空！")
	valid.MaxSize(desc, 100, "desc").Message("文章简介[desc]不能超过100个字符")
	valid.Required(createdBy, "createdBy").Message("文章创建者[createdBy]不能为空！")
	valid.Required(content, "content").Message("文章内容[content]不能为空！")
	valid.Required(tagId, "tagId").Message("文章标签[tag_id]不能为空！")

	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["tagId"] = tagId
		data["createdBy"] = createdBy
		data["state"] = state
		code = e.SUCCESS
		msg = e.GetMsg(code)
		models.AddArticle(data)
	} else {
		for _, value := range valid.Errors {
			msg = value.Error()
			break
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make([]int, 0),
	})
}

//修改文章
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	state := com.StrTo(c.Query("state")).MustInt()

	code := e.INVALID_PARAMS
	msg := e.GetMsg(code)
	valid := validation.Validation{}

	fmt.Println(id)
	valid.Min(id, 0, "id").Message("文章ID有误！")
	valid.Required(title, "title").Message("文章标题[title]不能为空！")
	valid.MaxSize(title, 100, "title").Message("文章标题[title]不能超过100个字符")
	valid.Required(desc, "desc").Message("文章简介[desc]不能为空！")
	valid.MaxSize(desc, 100, "desc").Message("文章简介[desc]不能超过100个字符")
	valid.Required(modifiedBy, "modifiedBy").Message("文章创建者[modified_by]不能为空！")
	valid.Required(content, "content").Message("文章内容[content]不能为空！")
	valid.Required(tagId, "tagId").Message("文章标签[tag_id]不能为空！")

	if !valid.HasErrors() {
		if models.ExistsArticleById(id) {
			data := make(map[string]interface{})
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["tagId"] = tagId
			data["modifiedBy"] = modifiedBy
			data["state"] = state
			code = e.SUCCESS
			fmt.Println(data)
			models.EditArticle(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
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
		"data": make([]int, 0),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS
	if id > 0 {
		if models.ExistsArticleById(id) {
			code = e.SUCCESS
			models.DeleteArticle(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make([]int, 0),
	})
}
