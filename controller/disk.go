package controller

import (
	"NetworkDisk/library"
	"NetworkDisk/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DiskHomeHandler 网盘首页
func DiskHomeHandler(c *gin.Context) {

	pageName := "全部文件"

	category, _ := strconv.Atoi(c.Request.FormValue("category"))
	parentId, _ := strconv.Atoi(c.Request.FormValue("parentId"))

	switch category {
	case 1:
		pageName = "图片"
	case 2:
		pageName = "文档"
	case 3:
		pageName = "视频"
	case 4:
		pageName = "音乐"
	case 5:
		pageName = "种子"
	case 6:
		pageName = "其他"
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"pageName" : pageName,
		"category" : category,
		"parentId" : parentId,
	})
}

// DiskUploadHandler 上传页面
func DiskUploadHandler(c *gin.Context) {
	parentId, _ := strconv.Atoi(c.Request.FormValue("parentId"))
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"pageName" : "上传",
		"parentId" : parentId,
	})
}

// DiskRecycleHandler 回收站页面
func DiskRecycleHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "recycle.html", gin.H{
		"pageName" : "回收站",
	})
}

// DiskFileListHandler 网盘文件列表
func DiskFileListHandler(c *gin.Context) {

	token, _ := GetToken(c)
	category, _ := strconv.Atoi(c.Request.FormValue("category"))
	parentId, _ := strconv.Atoi(c.Request.FormValue("parentId"))

	// path := c.DefaultPostForm("path", "")

	fileList, err := model.GetFileList(parentId , token.UserId, category, 1)
	if err != nil {
		c.JSON(http.StatusOK, library.RespMessage(10000,"没有文件", ""))
	} else {
		c.JSON(http.StatusOK, library.RespMessage(10000,"获取成功", fileList))
	}
}

// AddDirectoryHandler 添加目录
func AddDirectoryHandler(c *gin.Context) {
	fileName := c.DefaultPostForm("fileName", "")
	parentId, _ := strconv.Atoi(c.DefaultPostForm("parentId", "0"))
	if len(fileName) < 1 {
		c.JSON(http.StatusOK, library.RespMessage(10001,"文件名不能为空", ""))
		return
	}

	token, _ := GetToken(c)
	_, status := model.AddDirectory(fileName, token.UserId, parentId)

	if status {
		c.JSON(http.StatusOK, library.RespMessage(10000,"添加成功", ""))
	} else {
		c.JSON(http.StatusOK, library.RespMessage(10002,"添加失败", ""))
	}

}

