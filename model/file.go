package model

import (
	"NetworkDisk/dao"
	"NetworkDisk/library"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)



// SaveUploadedFile 添加文件信息到数据库中
func SaveUploadedFile(file *multipart.FileHeader, filePath string, userId int, parentId int) (*dao.File, error) {
	fileInfo := &dao.File {}

	src, err := file.Open()
	if err != nil {
		return fileInfo, err
	}
	defer src.Close()

	filenameAll := path.Base(file.Filename)
	// 扩展名
	fileType := strings.Trim(path.Ext(filePath), ".")
	// 文件名
	_, fileName := path.Split(file.Filename)
	// 判断文件是否存在
	fileExist := FileExist(filePath)
	if fileExist {
		currentDateTime := time.Now().Format("20060102_150405")
		fileName = filenameAll[0:len(filenameAll) - len(fileType) - 1] +"_"+ currentDateTime +"."+fileType
		filePath = "./upload/"+fileName
	}
	fmt.Println("filePath = " , filePath)
	// 创建本地目录
	newFile, err := os.Create(filePath)
	if err != nil {
		return fileInfo, err
	}
	defer newFile.Close()

	// 保存到本地
	fileSize, err := io.Copy(newFile, src)
	if err != nil {
		return fileInfo, err
	}

	// 文件信息
	fileInfo.UserId = userId
	fileInfo.ParentId = parentId
	fileInfo.FileName = fileName
	fileInfo.FilePath = filePath
	fileInfo.FileSize = fileSize
	fileInfo.UploadTime = time.Now().Format("2006-01-02 15:04:05")
	// seek 设置文件偏移量
	newFile.Seek(0, 0)
	fileInfo.FileHash = library.FileSha1(newFile)
	// 获取文件类型
	fileInfo.Category= GetFileCategory(file.Filename)
	// 文件扩展名
	fileInfo.FileType = fileType

	// 保存文件信息到数据库中
	_, err = dao.SaveFileInfo(fileInfo)
	if err != nil {
		// 文件数据保存失败
		return fileInfo, err
	}
	return fileInfo, err
}

// GetFileList 文件列表
func GetFileList(parentId int, userId int, category int, status int) (interface{}, error) {
	return dao.GetFileList(parentId, userId, category, status)
}

// AddDirectory 添加目录
func AddDirectory(fileName string, userId int, parentId int) (int64, bool) {

	file, err := dao.GetFileInfoByFileName(fileName, userId, 1)
	log.Println(file)
	if err != nil {
		return 0, false
	}
	// 文件存在随机一个文件名
	if file != nil {
		fileName = fileName+"_"+time.Now().Format("20060102_150405")
	}

	// 添加目录
	fileInfo := &dao.File{
		UserId: userId,
		FileName: fileName,
		FilePath: fileName,
		ParentId: parentId,
		FileHash: "",
		FileSize: 0,
		FileType: "dir",
	}
	log.Println(fileInfo)
	id, err := dao.SaveFileInfo(fileInfo)
	if err != nil {
		return 0, false
	}
	return id, true
}

// 根据 FileHash 查看文件是否存在
//func GetFileByFileHash(fileHash string) (*File, error){
//
//	return nil, nil
//}

// GetFileCategory	获取文件类型
func GetFileCategory(fileName string) int {

	fileType := make(map[string]int)
	// 图片
	fileType["webp"] = 1
	fileType["bmp"] = 1
	fileType["pcx"] = 1
	fileType["tif"] = 1
	fileType["gif"] = 1
	fileType["jpeg"] = 1
	fileType["tga"] = 1
	fileType["exif"] = 1
	fileType["fpx"] = 1
	fileType["svg"] = 1
	fileType["psd"] = 1
	fileType["cdr"] = 1
	fileType["pcd"] = 1
	fileType["dxf"] = 1
	fileType["ufo"] = 1
	fileType["epf"] = 1
	fileType["eps"] = 1
	fileType["ai"] = 1
	fileType["png"] = 1
	fileType["hdri"] = 1
	fileType["raw"] = 1
	fileType["wmf"] = 1
	fileType["flic"] = 1
	fileType["emf"] = 1
	fileType["ico"] = 1
	// 文档
	fileType["iso"] = 2
	fileType["rar"] = 2
	fileType["html"] = 2
	fileType["zip"] = 2
	fileType["exe"] = 2
	fileType["pdf"] = 2
	fileType["txt"] = 2
	fileType["doc"] = 2
	fileType["docx"] = 2
	fileType["xls"] = 2
	fileType["xlsx"] = 2
	fileType["ppt"] = 2
	fileType["pptx"] = 2
	// 视频
	//AVI、mov、rmvb、rm、FLV、mp4、3GP
	fileType["rm"] = 3
	fileType["avi"] = 3
	fileType["rmvb"] = 3
	fileType["flv"] = 3
	fileType["mp4"] = 3
	fileType["3gp"] = 3
	// 音乐
	fileType["mp3"] = 4
	// 种子
	fileType["torrent"] = 5
	// 其他
	fileType["other"] = 6
	fileExt := strings.Trim(path.Ext(fileName), ".")
	fileExt = strings.ToLower(fileExt)

	fileTypeId := int(fileType[fileExt])
	if fileTypeId == 0 {
		return 6
	}
	return fileTypeId
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}




