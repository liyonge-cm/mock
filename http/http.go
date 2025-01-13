package http

import (
	"fmt"
	"net/http"

	"github.com/liyonge-cm/mock/apis"
	"github.com/liyonge-cm/mock/router"

	"github.com/gin-gonic/gin"
)

const PORT = 8080

// 初始化路由
func init() {
	// 设置为发布模式（初始化路由之前设置）
	gin.SetMode(gin.ReleaseMode)
	// gin 默认中间件
	r := gin.Default()
	r.Use(cors)

	// 访问一个错误网站时，返回404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})

	m := router.NewRouter(r)
	a := apis.NewApis(m)
	err := a.InitRouter("./json")
	if err != nil {
		panic(err)
	}
	if err = r.Run(fmt.Sprintf(":%v", PORT)); err != nil {
		panic(err)
	}
}

func cors(c *gin.Context) {
	origin := c.GetHeader("origin")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Headers", "*")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
