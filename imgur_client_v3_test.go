package main

import (
	"fmt"
	"os"
	"testing"
)

func TestUploadV3(t *testing.T) {
	client := ImgurClientV3{ClientId: os.Getenv("IMGUR_CLIENT_ID")}
	res, err := client.AnonymousUpload("./test_image.png")
	if err != nil {
		t.Errorf("TestUploadV3: ", err)
	}
	fmt.Println(res)
}