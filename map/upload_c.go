package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
	"mime/multipart"
	"os"
	"io"
)

func postFile(filename string, targeUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作, 文件名 不要带路径
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "test.log")
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// 打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	defer fh.Close()

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targeUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string((resp_body)))
	return nil

}

//sample usage
func main() {
	target_url := "http://localhost:9090/upload"
	filename := "/Users/lzp/Desktop/GoTools/goStudy/map/array.go"
	postFile(filename, target_url)

}

