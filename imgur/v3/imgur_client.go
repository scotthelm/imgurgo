package v3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ImgurClient struct {
	ClientId        string
	ClientSecret    string
	AccessToken     string
	ExpiresIn       int64
	TokenType       string
	RefreshToken    string
	AccountUsername string
	AccountId       int64
	http.Client
}

const API_BASE = "https://api.imgur.com/3"
const UPLOAD_IMAGE = API_BASE + "/image"
const AUTH = "https://api.imgur.com/oauth2/authorize"
const TOKEN = "https://api.imgur.com/oauth2/token"

func NewClient(key, secret string) *ImgurClient {
	return &ImgurClient{ClientId: key, ClientSecret: secret}
}

func (cl ImgurClient) AnonymousUpload(path string) (ImgurResponse, error) {
	var err error = nil
	ir := ImgurResponse{}
	auth_header := []string{"Client-ID " + cl.ClientId}
	req, err := cl.newFileUploadRequest(
		UPLOAD_IMAGE,
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

func (cl *ImgurClient) GetAuthorizationUrl(authType string) string {
	return fmt.Sprintf("%s?client_id=%s&response_type=%s", AUTH, cl.ClientId, authType)
}

func (cl *ImgurClient) Authorize(pin, authType string) (ImgurAuthResponse, error) {
	ir := ImgurAuthResponse{}
	v := url.Values{}
	v.Set("client_id", cl.ClientId)
	v.Set("client_secret", cl.ClientSecret)
	v.Set("grant_type", authType)
	v.Set("pin", pin)
	response, err := cl.PostForm(TOKEN, v)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(body, &ir)
	} else {
		err = errors.New(fmt.Sprintf("ImgurClient#Authorize: Status code: %d, authtype: %s", response.StatusCode, authType))
	}
	return ir, err
}

// Creates a new file upload http request with optional extra params
func (cl *ImgurClient) newFileUploadRequest(
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
