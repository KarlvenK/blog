package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//add category
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//to do search all the articles below the category

//search cate list
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//edit cate
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//delete cate
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
