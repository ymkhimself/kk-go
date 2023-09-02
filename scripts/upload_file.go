package main

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "/Users/ymk/Documents/魔法骰子/美食"
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fileName := filepath.Base(path)
		if fileName == ".DS_Store" {
			return nil
		}
		typeKey := strings.TrimSuffix(fileName, filepath.Ext(fileName)) // 拿到tykeKey
		fmt.Println(typeKey)
		client := http.Client{}
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		theme := "10"
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fmt.Println(fileName)
		formFile, err := writer.CreateFormFile("file", file.Name())
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(formFile, file)
		if err != nil {
			panic(err)
		}

		writer.WriteField("theme", theme)
		writer.WriteField("type", typeKey)
		writer.WriteField("operation", "OVERWRITE_ALL")
		writer.WriteField("slotsNeedDetect", "Sheet2")

		err = writer.Close()
		if err != nil {
			panic(err)
		}

		request, err := http.NewRequest("POST", "http://localhost:8156/magicDice/uploadTypeAsync", body)
		if err != nil {
			panic(err)
		}
		request.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		all, err := io.ReadAll(resp.Body)
		fmt.Println(string(all))
		file.Close()
		fmt.Println("\n")
		return nil
	})
}
