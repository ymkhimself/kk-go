package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fileNames := []string{"a", "b", "c"}
	prePath := "/sa"
	client := http.Client{}
	for _, name := range fileNames {
		filePath := filepath.Join(prePath, name)
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
			return
		}
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		formFile, err := writer.CreateFormFile("file", name)
		if err != nil {
			panic(err)
			return
		}
		_, err = io.Copy(formFile, file)
		_ = writer.WriteField("", "") // 写k-v 其他数据放进去
		writer.Close()

		request, err := http.NewRequest("POST", "", body)
		if err != nil {
			panic(err)
			return
		}
		request.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := client.Do(request)
		if err != nil {
			panic(err)
			return
		}
		all, err := io.ReadAll(resp.Body)
		fmt.Println(string(all))
		file.Close()
	}
}
