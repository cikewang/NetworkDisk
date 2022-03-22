package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

func UploadHandler(c *gin.Context) {
	// 接收文件流及存储到本地目录
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)



	//defer file.Close()
	//
	//fileMeta := meta.FileMeta{
	//	FileName: head.Filename,
	//	FilePath: "./upload/"+head.Filename,
	//	UploadTime: time.Now().Format("2006-01-02 15:04:05"),
	//
	//}
	//// 创建本地目录
	//newFile, err := os.Create(fileMeta.FilePath)
	//if err != nil {
	//	fmt.Println("Failed to create file, err: %s\n", err.Error())
	//	return
	//}
	//defer newFile.Close()
	//
	//// 保存到本地
	//fileMeta.FileSize, err = io.Copy(newFile, file)
	//if err != nil {
	//	fmt.Println("Failed to save data into file, err :%s \n", err.Error())
	//	return
	//}
	//// seek 设置文件偏移量
	//newFile.Seek(0, 0)
	//fileMeta.FileSha1 = util.FileSha1(newFile)
	//// meta.UpdateFileMeta(fileMeta)
	//
	//ossPath := string("oss/"+fileMeta.FileSha1)
	//err = oss.Bucket().PutObjectFromFile(ossPath, fileMeta.FilePath)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	w.Write([]byte("OSS 文件上传失败"))
	//	return
	//}
	//fileMeta.FilePath = ossPath
	//
	//// 信息保存到数据库中
	//_ = meta.UpdateFileMetaDB(fileMeta)
	//status := database.UserFileUploadFinished(1, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
	//if status {
	//	http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	//} else {
	//	w.Write([]byte("更新失败"))
	//}
}