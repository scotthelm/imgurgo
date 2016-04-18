package v3

import (
	"fmt"
	"os"
	"testing"
)

func TestAnonymousUpload(t *testing.T) {
	return
	client := NewClient(os.Getenv("IMGUR_CLIENT_ID"), os.Getenv("IMGUR_CLIENT_SECRET"), "", "")
	res, err := client.AnonymousUpload("./test_image.png")
	if err != nil {
		t.Errorf("TestUploadV3: ", err)
	}
	fmt.Println(res)
}

func TestGetAuthorizationUrl(t *testing.T) {
	client := NewClient(os.Getenv("IMGUR_CLIENT_ID"), os.Getenv("IMGUR_CLIENT_SECRET"), "", "")
	url := client.GetAuthorizationUrl("pin")
	expected := fmt.Sprintf("%s?client_id=%s&response_type=pin", AUTH, client.ClientId)
	fmt.Println(url)

	if url != expected {
		t.Error("TestUploadV3#GetAuthorization: unexpected url, got %s", url)
	}
}

func TestAuthorizePin(t *testing.T) {
	return
	client := NewClient(os.Getenv("IMGUR_CLIENT_ID"), os.Getenv("IMGUR_CLIENT_SECRET"), "", "")
	res, err := client.Authorize("df7b715c93", "pin")
	fmt.Println(res)
	if err != nil {
		t.Errorf("TestAuthorize: %s", err)
	}
}

func TestRefresh(t *testing.T) {
	return
	client := NewClient(
		os.Getenv("IMGUR_CLIENT_ID"),
		os.Getenv("IMGUR_CLIENT_SECRET"),
		os.Getenv("IMGUR_ACCESS_TOKEN"),
		os.Getenv("IMGUR_REFRESH_TOKEN"),
	)
	client.Refresh()
}

func TestAccount(t *testing.T) {
	client := client()
	account, err := client.GetAccount(os.Getenv("IMGUR_USERNAME"))
	if err != nil {
		t.Errorf("TestAccount: %s\n", err)
	}
	fmt.Printf("%+v\n", account)
}

func TestImages(t *testing.T) {
	client := client()
	images, err := client.GetAccountImages(0)
	if err != nil {
		t.Errorf("TestImages: %s\n", err)
	}
	fmt.Printf("%+v\n", images)
}

func client() *ImgurClient {
	client := NewClient(
		os.Getenv("IMGUR_CLIENT_ID"),
		os.Getenv("IMGUR_CLIENT_SECRET"),
		os.Getenv("IMGUR_ACCESS_TOKEN"),
		os.Getenv("IMGUR_REFRESH_TOKEN"),
	)
	client.AccountUsername = os.Getenv("IMGUR_USERNAME")
	return client
}
