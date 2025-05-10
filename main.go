package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/ltask/api"
	"io/fs"
	"net/http"
)

//go:embed web/dist/*
var staticFS embed.FS

// 自定义文件系统处理前端路由
type spaFS struct {
	fs.FS
}

func (f *spaFS) Open(name string) (fs.File, error) {
	// 尝试打开请求的文件
	file, err := f.FS.Open(name)
	if err == nil {
		return file, nil
	}
	// 文件不存在时返回 index.html
	return f.FS.Open("index.html")
}

func main() {
	// 初始化Gin
	gin.ForceConsoleColor()
	router := gin.Default()

	// 创建子文件系统并包装为SPA文件系统
	distFS, err := fs.Sub(staticFS, "web/dist")
	if err != nil {
		panic("无法加载静态资源: " + err.Error())
	}

	// 注册API路由
	registerAPIRoutes(router)

	// 注册静态文件处理
	router.Use(staticHandler(distFS))

	// 启动服务器
	router.Run(":8080")
}

// 静态文件处理中间件
func staticHandler(distFS fs.FS) gin.HandlerFunc {
	// 包装为SPA文件系统
	spaFileSystem := &spaFS{distFS}

	return gin.WrapH(http.FileServer(http.FS(spaFileSystem)))
}

// 注册API路由
func registerAPIRoutes(router *gin.Engine) {
	// 基础路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// API路由组
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/user/login", api.Login)
		apiGroup.GET("/hello", api.Hello)
		apiGroup.GET("/version", api.Version)
	}
}
