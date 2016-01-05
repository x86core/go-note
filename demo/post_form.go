package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"log"
)

type UploadFile struct {
	FieldName string
	FilePath  string
}

func writeFields(writer *multipart.Writer, fields map[string]string) {
	for key, val := range fields {
		writer.WriteField(key, val)
	}
}

func writeFiles(writer *multipart.Writer, files []UploadFile) {
	for _, aFile := range files {
		file, err := os.Open(aFile.FilePath)

		if err != nil {
			//continue
		}

		part, err := writer.CreateFormFile(aFile.FieldName, aFile.FilePath)
		io.Copy(part, file) //file to writer
		if err != nil {
			// copy error
			log.Println("copy err:", err)
		}
	}
}

func CollectionData(fields map[string]string, files []UploadFile) (body *bytes.Buffer, contentType string, er error) {
	body = &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writeFields(writer, fields)
	writeFiles(writer, files)
	contentType = writer.FormDataContentType()
	writer.Close()

	return body, contentType, nil
}

func PostAction(url string, body *bytes.Buffer, contentType string) {
	log.Println(contentType)
	request, err := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", contentType)
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Println("resp err:", err)
	}else{
		response.Body.Close()
	}
}

// test
func main() {
	var f = map[string]string{
		"key":"word",
		"time":"2016",
	}
	var fSet = []UploadFile{
		UploadFile{
			FieldName:"upfile",
			FilePath:"/home/forme/Pictures/240451-13050106450911.jpg",
		},
	}

	var url = "http://127.0.0.1:8500/index.php"
	body, ctype, _ := CollectionData(f, fSet)

	PostAction(url, body, ctype)
}