package main

import (
	_ "golang-swagger/docs" // 一定要加上這個 docs 資料夾的路徑

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title <swagger 的 title>
// @version 2.0
// @description <就是描述>
// @host localhost:8080
func main() {
	router := gin.Default()

	router.Group("/swagger-example").
		GET("/hello", hello).
		POST("/echo", echo)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run()
}
