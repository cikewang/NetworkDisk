package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DiskHomeHandler(c *gin.Context) {

	PageName := "全部文件"

	category, _ := strconv.Atoi(c.Request.FormValue("category"))
	fmt.Printf("category = %d \r\n", category)
	switch category {
	case 1:
		PageName = "图片"
	case 2:
		PageName = "文档"
	case 3:
		PageName = "视频"
	case 4:
		PageName = "音乐"
	case 5:
		PageName = "种子"
	case 6:
		PageName = "其他"
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"PageName" : PageName,
	})
}

func DiskUploadHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"PageName" : "上传",
	})
}

func DiskRecycleHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "recycle.html", gin.H{
		"PageName" : "回收站",
	})
}

