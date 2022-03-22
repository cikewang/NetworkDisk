package controller

import (
	"NetworkDisk/library"
	"NetworkDisk/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	pwd_salt = "!Q@W#E$R"
)

// 用户登录页面
func UserLoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":"用户登录",
	})
	return
}

// 用户注册页面
func UserRegisterHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title":"用户注册",
	})
	return
}

// 用户注册操作
func UserDoRegisterHandler(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	password_verify := c.DefaultPostForm("password_verify", "")
	auth := c.Request.Header.Get("Authorization")

	fmt.Println("auth = ", auth)

	// 用户名长度不能小于6位
	if len(username) < 5 {
		c.JSON(http.StatusForbidden, library.RespMessage(10001,"用户名长度不能小于5位", ""))
		return
	}

	// 密码长度不能小于6位
	if len(password) < 5 {
		c.JSON(http.StatusForbidden, library.RespMessage(10002,"密码长度不能小于5位", ""))
		return
	}

	// 两次密码输入是否一致
	if password != password_verify {
		c.JSON(http.StatusForbidden, library.RespMessage(10003,"两次密码输入不一致", ""))
		return
	}

	// 检查用户名是否已经被注册
	_, err := model.GetUserInfo(username)
	if err == nil {
		c.JSON(http.StatusOK, library.RespMessage(10005,"用户名已经存在", ""))
		return
	}

	// 用户注册
	password = library.MD5([]byte(password + pwd_salt))
	regStatus := model.UserRegister(username, password)
	if regStatus {
		data := &library.Redirect{
			Location: "/user/login",
		}
		c.JSON(http.StatusOK, library.RespMessage(10000,"注册成功，请登录", data))
	} else {
		c.JSON(http.StatusForbidden, library.RespMessage(10006,"注册失败", ""))
	}
}

// 用户登录
func UserDoLoginHandler(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	// 用户名长度不能小于6位
	if len(username) < 5 {
		c.JSON(http.StatusForbidden, library.RespMessage(10001,"用户名长度不能小于5位", ""))
		return
	}

	// 密码长度不能小于6位
	if len(password) < 5 {
		c.JSON(http.StatusForbidden, library.RespMessage(10002,"密码长度不能小于5位", ""))
		return
	}

	// 检查用户名是否已经被注册
	user, err := model.GetUserInfo(username)
	if err != nil {
		c.JSON(http.StatusOK, library.RespMessage(10005,"账户或密码错误", ""))
		return
	}

	password = library.MD5([]byte(password + pwd_salt))
	if (user.Password != password) {
		c.JSON(http.StatusOK, library.RespMessage(10005,"账户或密码错误", ""))
		return
	}

	// 生成 token
	token := CreateToken(user)
	if token == "" {
		c.JSON(http.StatusOK, library.RespMessage(10006,"秘钥生成失败，请重试", ""))
		return
	}

	if user.UserName != "" {
		data := &library.Redirect{
			Location: "/disk/home",
			Token : token,
		}
		c.JSON(http.StatusOK, library.RespMessage(10000,"登录成功", data))
	} else {
		c.JSON(http.StatusForbidden, library.RespMessage(10006,"登录成功", ""))
	}
}

// 用户退出操作
func UserLogoutHandler(c *gin.Context) {
	data := &library.Redirect{
		Location: "/user/login",
	}
	c.JSON(http.StatusOK, library.RespMessage(10000,"退出成功", data))
}

// 用户信息
func UserInfoHandler(c *gin.Context) {
	token, bool := GetToken(c)

	data := make(map[string]interface{})
	code := 10000
	message := "用户信息"
	if bool {
		data = gin.H{
			"username" : token.Username,
		}
	} else {
		message = "用户信息认证失败，请重新登录"
		code = 99999
		data = gin.H{
			"location" : "/user/login",
		}
	}

	c.JSON(http.StatusOK, library.RespMessage(code, message, data))
	return
}
