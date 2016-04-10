package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type ImgurClientV3 struct {
	ClientId string
	http.Client
}

const V3_API_BASE = "https://api.imgur.com/3"
const V3_UPLOAD_IMAGE = V3_API_BASE + "/image"

func (cl ImgurClientV3) AnonymousUpload(path string) (ImgurResponseV3, error) {
	var err error = nil
	ir := ImgurResponseV3{}
	auth_header := []string{"Client-ID " + os.Getenv("IMGUR_CLIENT_ID")}
	req, err := cl.newFileUploadRequest(
		V3_UPLOAD_IMAGE,
		nil,
		"image",
		"./test.png",
	)
	req.Header.Add("Authorization", strings.Join(auth_header, " "))
	response, err := cl.Do(req)
	if err != nil {
		return ir, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &ir)
	if err != nil {
		return ir, err
	}
	return ir, err
}

// Creates a new file upload http request with optional extra params
func (cl *ImgurClientV3) newFileUploadRequest(
	uri string,
	params map[string]string,
	fileParam,
	path string,
) (*http.Request, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileParam, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, err

}
