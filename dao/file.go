package dao

import (
	"NetworkDisk/data/database/mysql"
	sql2 "database/sql"
	"fmt"
)

type File struct {
	Id 			int	`json:"id"`
	UserId 		int `json:"userId"`
	ParentId	int `json:"parentId"`
	FileName string `json:"fileName"`
	FilePath string `json:"filePath"`
	FileSize int64 `json:"fileSize"`
	FileHash string `json:"fileHash"`
	FileType string `json:"fileType"`
	Category int 	`json:"category"`
	CreateTime string `json:"createTime"`
	UploadTime string `json:"updateTime"`
}

// SaveFileInfo 保存文件信息到数据库
func SaveFileInfo(file *File) (int64, error) {
	sql := "insert into file (`user_id`, `file_hash`, `file_name`, `file_size`, `file_path`, `file_type`, `category`, `parent_id`) "+
			"values (?,?,?,?,?,?,?,?)"
	stmt, err := mysql.DBConn().Prepare(sql)
	if err != nil {
		fmt.Println("Failed to prepare statement, err :"+err.Error())
		return  0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(file.UserId, file.FileHash, file.FileName, file.FileSize, file.FilePath, file.FileType, file.Category, file.ParentId)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return result.LastInsertId()
}

// GetFileInfoByFileName 根据文件名查询文件信息
func GetFileInfoByFileName(fileName string, userId int, status int) (*File, error){

	sql := "select id, parent_id, file_hash, file_name, file_size, file_path FROM file where file_name = ? and user_id = ? and status = ?"
	stmt, err := mysql.DBConn().Prepare(sql)
	if err != nil {
		fmt.Println("Failed to prepare statement, err :"+err.Error())
		return nil, err
	}
	defer stmt.Close()

	file := &File{}
	err = stmt.QueryRow(fileName, userId, status).Scan(&file.Id, &file.ParentId, &file.FileHash, &file.FileName, &file.FileSize, &file.FilePath)
	if err == sql2.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	fmt.Println(file)
	return file, nil
}

// GetFileList 文件列表
func GetFileList(parentId int, userId int, category int, status int) (file []File, err error) {

	sql_str := ""
	// 查询目录下文件
	if category > 0 {
		sql_str = "select id, parent_id, file_hash, file_name, file_size, file_path, file_type, create_time, update_time, category FROM file "+
			"where user_id = ? and category = ? and status = ?"
	} else {
		sql_str = "select id, parent_id, file_hash, file_name, file_size, file_path, file_type, create_time, update_time, category FROM file "+
			"where user_id = ? and parent_id = ? and status = ?"
	}

	stmt, err := mysql.DBConn().Prepare(sql_str)
	if err != nil {
		fmt.Println("Failed to prepare statement, err :"+err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows := &sql2.Rows{}
	if category > 0 {
		rows, err = stmt.Query(userId, category, status)
	} else {
		rows, err = stmt.Query(userId, parentId, status)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var fileList []File
	for rows.Next() {
		file := File{}
		err = rows.Scan(&file.Id, &file.ParentId, &file.FileHash, &file.FileName, &file.FileSize, &file.FilePath, &file.FileType, &file.CreateTime, &file.UploadTime, &file.Category)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fileList = append(fileList, file)
	}
	return fileList, err
}
