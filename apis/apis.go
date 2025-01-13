package apis

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/liyonge-cm/mock/router"
)

type Apis struct {
	*router.Router
}

type ApiRouter struct {
	Router   string                 `json:"router"`
	Method   string                 `json:"method"`
	Request  map[string]interface{} `json:"request"`
	Response map[string]interface{} `json:"response"`
}

func NewApis(m *router.Router) *Apis {
	return &Apis{m}
}

func (a *Apis) InitRouter(filePath string) error {
	apis, err := a.readeJsonApi(filePath)
	if err != nil {
		return err
	}

	for _, v := range apis {
		fmt.Println("mock api: ", v.Router)

		method := strings.ToLower(v.Method)
		switch method {
		case "get":
			a.SetGetRouter(v.Router, a.makeRouteFunc(v))
		case "post":
			a.SetPostRouter(v.Router, a.makeRouteFunc(v))
		case "put":
			a.SetPutRouter(v.Router, a.makeRouteFunc(v))
		case "delete":
			a.SetDelRouter(v.Router, a.makeRouteFunc(v))
		default:
			fmt.Println("mock api: ", v.Router, "method error", method)
		}
	}
	return nil
}

func (a *Apis) makeRouteFunc(api *ApiRouter) func(req map[string]interface{}) interface{} {
	return func(req map[string]interface{}) interface{} {
		reqb, _ := json.Marshal(api.Request)
		fmt.Println("sample request", string(reqb))
		return api.Response
	}
}

func (a *Apis) readeJsonApi(jsonPath string) (apis []*ApiRouter, err error) {
	apis = []*ApiRouter{}
	files, err := os.ReadDir(jsonPath)
	if err != nil {
		return apis, err
	}

	for _, file := range files {
		if !file.IsDir() { // 只读取文件，不读取子文件夹
			fileApis, _ := a.parseApiJson(path.Join(jsonPath, file.Name()))
			apis = append(apis, fileApis...)
		}
	}
	return apis, err
}

func (a *Apis) parseApiJson(filePath string) (apis []*ApiRouter, err error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &apis)
	if err != nil {
		fmt.Println("json unmarshal file ", filePath, "error", err.Error())
		return nil, err
	}
	return apis, err
}
