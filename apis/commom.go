package apis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func ParseJson(str string) (data map[string]interface{}) {
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		fmt.Println("Unmarshal error", err.Error())
	}
	return
}

func ParseFileBase64(file string) (data string) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("ReadFile error", err.Error())
	}
	data = base64.StdEncoding.EncodeToString(fileBytes)
	return
}

func CreateFormdataFile(formdata *multipart.Form, key string) (fileName string, err error) {
	// 获取上传文件
	files := formdata.File[key]
	if len(files) == 0 {
		return fileName, fmt.Errorf("no file")
	}
	fileName = "./" + files[0].Filename
	file, err := files[0].Open()
	if err != nil {
		return fileName, err
	}
	defer file.Close()
	out, err := os.Create(fileName)
	if err != nil {
		return fileName, err
	}
	defer out.Close()

	io.Copy(out, file)
	return
}
