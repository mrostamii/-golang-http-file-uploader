package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	url := "http://127.0.0.1:8080/upload/image"
	method := "POST"

	//to get the file form cmd
	//var uploadingFilePath string = "C:/" + os.Args[1]

	//to get the file from input
	fmt.Printf("Enter the file name that located on: C:/")
	var inputPath string
	fmt.Scanln(&inputPath)
	var uploadingFilePath string = "C:/" + inputPath

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(uploadingFilePath)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base(uploadingFilePath))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {

		fmt.Println(errFile1)
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Printf("\nCongrats, the File has uploaded.\n\nHit the 'ENTER' key to exit.\n	 MIMTech")
	var input string
	fmt.Scanln(&input)
}
