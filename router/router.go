package router

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RouterFunc func(req map[string]interface{}) interface{}
type MultipartFunc func(req *multipart.Form) interface{}

type Router struct {
	eng *gin.Engine
}

func NewRouter(r *gin.Engine) *Router {
	m := &Router{eng: r}
	return m
}

func (m *Router) SetGetRouter(router string, f RouterFunc) {
	m.eng.GET(router, func(c *gin.Context) {
		m.handlerRouter(c, f)
	})
}
func (m *Router) SetPostRouter(router string, f RouterFunc) {
	m.eng.POST(router, func(c *gin.Context) {
		m.handlerRouter(c, f)
	})
}
func (m *Router) SetPutRouter(router string, f RouterFunc) {
	m.eng.PUT(router, func(c *gin.Context) {
		m.handlerRouter(c, f)
	})
}
func (m *Router) SetDelRouter(router string, f RouterFunc) {
	m.eng.DELETE(router, func(c *gin.Context) {
		m.handlerRouter(c, f)
	})
}

func (m *Router) SetRouterMulti(router string, f MultipartFunc) {
	m.eng.POST(router, func(c *gin.Context) {
		m.handlerRouterMultiFunc(c, f)
	})
}

func (m *Router) handlerRouter(c *gin.Context, routerFunc RouterFunc) {
	data := make(map[string]interface{})
	c.ShouldBindBodyWith(&data, binding.JSON)

	datab, _ := json.Marshal(data)
	fmt.Println("reqeust:", string(datab))

	res := routerFunc(data)

	resb, _ := json.Marshal(res)
	fmt.Println("response:", string(resb))

	c.JSON(http.StatusOK, res)
}

func (m *Router) handlerRouterMultiFunc(c *gin.Context, routerFunc MultipartFunc) {
	formdata, err := c.MultipartForm()
	if err != nil {
		fmt.Println("get MultipartForm error", err.Error())
		c.JSON(http.StatusBadRequest, formdata)
	}
	fmt.Println("formdata", formdata)
	res := routerFunc(formdata)
	//fmt.Println("response", res)
	c.JSON(http.StatusOK, res)
}
