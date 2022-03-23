package router

import (
	"NetworkDisk/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path/filepath"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()
	// CORS
	r.Use(cors.Default())
	// 引入 gin 框架模板文件
	//r.LoadHTMLGlob("view/**/*")
	r.HTMLRender = loadTemplates("./view")
	// 引入 static
	r.Static("/static", "static")
	r.Static("/upload", "upload")


	// 无权限认证
	r.GET("/", controller.DiskHomeHandler)
	// 用户登录页面
	r.GET("/user/login", controller.UserLoginHandler)
	// 用户登录
	r.POST("/user/doLogin", controller.UserDoLoginHandler)
	// 用户注册页面
	r.GET("/user/register", controller.UserRegisterHandler)
	// 用户注册操作
	r.POST("/user/doRegister", controller.UserDoRegisterHandler)

	/* 网盘页面 */
	r.GET("/disk/home", controller.DiskHomeHandler)
	r.GET("/disk/recycle", controller.DiskRecycleHandler)
	// 文件上传页面
	r.GET("/disk/upload", controller.DiskUploadHandler)

	// 以下接口验证权限
	// 权限拦截器
	r.Use(controller.Authorize())
	// 用户退出
	r.GET("/user/logout", controller.UserLogoutHandler)
	// 获取用户信息
	r.GET("/user/info", controller.UserInfoHandler)

	// 文件管理
	r.POST("/file/upload", controller.UploadFileHandler)
	r.POST("/file/uploads", controller.UploadFilesHandler)


	// 文件列表
	r.GET("/disk/list", controller.DiskFileListHandler)
	// 添加目录
	r.POST("/disk/add", controller.AddDirectoryHandler)

	return r
}

// 加载模板文件
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layout_simple, err := filepath.Glob(templatesDir + "/layout/simple.html")
	if err != nil {
		panic(err.Error())
	}
	layout_home, err := filepath.Glob(templatesDir + "/layout/home.html")
	if err != nil {
		panic(err.Error())
	}
	users, err := filepath.Glob(templatesDir + "/user/*.html")
	if err != nil {
		panic(err.Error())
	}
	disk, err := filepath.Glob(templatesDir + "/disk/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range users {
		layoutCopy := make([]string, len(layout_simple))
		copy(layoutCopy, layout_simple)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	for _, include := range disk {
		layoutCopy := make([]string, len(layout_home))
		copy(layoutCopy, layout_home)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	return r
}

