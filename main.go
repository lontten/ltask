package main

import "github.com/gin-gonic/gin"

func main() {
	// PowerShell 脚本（每1秒输出一次）
	//script := `
	//foreach ($i in 1..10) {
	//    Write-Host "进度: $i/10 - $(Get-Date -Format 'HH:mm:ss')"
	//    Start-Sleep -Seconds 1
	//}
	//`
	//ok := util.ShellExec(script, 11)
	//fmt.Println("命令执行", ok)

	if true {
		router := gin.Default()
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		router.Run() // listen and serve on 0.0.0.0:8080
	}
}
