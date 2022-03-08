package router

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(staticFiles fs.FS) *gin.Engine {
	router := gin.Default()

	//静态文件目录
	router.StaticFS("/", http.FS(staticFiles))

	return router
}
