package controller

import (
	"NetworkDisk/library"
	"NetworkDisk/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// UploadFileHandler 上传单个文件
func UploadFileHandler(c *gin.Context) {

	token, bool := GetToken(c)

	if !bool {
		c.JSON(http.StatusOK, library.RespMessage(10001,"权限认证失败", ""))
		return
	}

	// 接收文件并保存
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, library.RespMessage(10001,"文件上传失败，请重试", ""))
		return
	}

	parentId, _ := strconv.Atoi(c.DefaultPostForm("parentId", "0"))

	// 文件路径
	filePath := "./upload/" + file.Filename
	// 上传文件至指定的完整文件路径
	_, err = model.SaveUploadedFile(file, filePath, token.UserId, parentId)
	if err != nil {
		c.JSON(http.StatusOK, library.RespMessage(10002,"文件保存失败，请重试", ""))
		return
	}

	data := &library.Redirect{
		Location: "/disk/home?parentId="+strconv.Itoa(parentId),
	}
	c.JSON(http.StatusOK, library.RespMessage(10000,"上传成功", data))
}

// UploadFilesHandler 上传目录文件
func UploadFilesHandler(c *gin.Context) {

	token, bool := GetToken(c)
	if !bool {
		c.JSON(http.StatusOK, library.RespMessage(10001,"权限认证失败", ""))
		return
	}

	parentId, _ := strconv.Atoi(c.DefaultPostForm("parentId", "0"))

	form, _ := c.MultipartForm()
	files := form.File["files"]
	fileLength := len(files)
	if  fileLength <= 0 {
		c.JSON(http.StatusOK, library.RespMessage(10002,"上传文件不能为空", ""))
		return
	}
	// 获取文件目录
	dirFile := files[0]
	dir, _ := path.Split(dirFile.Filename)
	dir = strings.Trim(dir, "/")
	// 添加目录
	currentParentId, status := model.AddDirectory(dir, token.UserId, parentId)
	if !status {
		c.JSON(http.StatusOK, library.RespMessage(10003,"上传失败，请重试", ""))
		return
	}

	number := 0
	for _, file := range files {
		_, fileName := path.Split(file.Filename)
		// 文件路径
		filePath := "./upload/" + fileName
		// 上传文件至指定的完整文件路径
		_, err := model.SaveUploadedFile(file, filePath, token.UserId, int(currentParentId))
		if err == nil {
			number++
		}
	}

	if number == 0 {
		c.JSON(http.StatusOK, library.RespMessage(10002,"文件保存失败，请重试", ""))
		return
	}

	data := &library.Redirect{
		Location: "/disk/home?parentId=" + strconv.Itoa(parentId),
	}
	c.JSON(http.StatusOK, library.RespMessage(10000,"上传成功", data))

}