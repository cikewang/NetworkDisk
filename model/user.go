package model

import (
	"NetworkDisk/data/database/mysql"
	"fmt"
	"log"
)

// 新用户注册
func UserRegister(username string, password string) bool  {
	stmt, err := mysql.DBConn().Prepare("insert into user (`user_name`, `password`) values(?, ?)")
	if err != nil {
		log.Println("User Insert Err : "+err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, password)
	if err != nil {
		log.Println("User insert Exec Err : "+err.Error())
		return false
	}

	if rowAffected, err := ret.RowsAffected(); nil == err && rowAffected > 0 {
		return true
	}
	return false
}

// 用户登录
func UserLogin(username string, password string) bool {
	stmt , err := mysql.DBConn().Prepare("select id, user_name, password from user where user_name = ? AND password = ? AND status = 1")
	if err != nil {
		fmt.Println("UserLogin select ERR : "+err.Error())
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, password)
	if err != nil {
		fmt.Println("UserLogin select Query ERR : "+ err.Error())
		return false
	} else if rows == nil {
		fmt.Println("Userlogin username not found : "+ username)
		return false
	}
	return true
	//pRows := mysql.ParseRows(rows)
}

// 更新 Token
func UpdateToken(username string, token string) bool {
	// replace
	stmt, err := mysql.DBConn().Prepare("insert user_token (`user_name`, `token`) values(?, ?)")
	if err != nil {
		fmt.Println("UpdateToken insert ERR : "+ err.Error())
		return false
	}

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println("UpdateToken insert exec ERR : "+err.Error())
		return false
	}
	return true
}

type User struct {
	Id int
	UserName string
	Password string
	Email string
	Phone string
	CreateTime string
	UpdateTime string
	Status	int
}

// 获取用户信息
func GetUserInfo(username string) (User, error) {
	user := User{}
	stmt, err := mysql.DBConn().Prepare(
		"select id, user_name, password, create_time from user where user_name = ?")
	if err != nil {
		log.Println("GetUserInfo select err : "+err.Error())
		return user, err
	}
	defer stmt.Close()

	// 执行查询的SQL语句
	err = stmt.QueryRow(username).Scan(&user.Id, &user.UserName, &user.Password, &user.CreateTime)
	return user, err
}