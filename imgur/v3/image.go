package v3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Image struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Datetime    int64  `json:"datetime"`
	Type        string `json:"type"`
	Animated    bool   `json:"animated"`
	Width       int64  `json:"width"`
	Height      int64  `json:"height"`
	Size        int64  `json:"size"`
	Views       int64  `json:"views"`
	Bandwidth   int64  `json:"bandwidth"`
	Deletehash  string `json:"deletehash"`
	Name        string `json:"name"`
	Section     string `json:"section"`
	Link        string `json:"link"`
	Gifv        string `json:"gifv"`
	Mp4         string `json:"mp4"`
	Webm        string `json:"webm"`
	Mp4Size     int64  `json:"mp4_size"`
	WebmSize    int64  `json:"webm_size"`
	Looping     bool   `json:"looping"`
	Favorite    bool   `json:"favorite"`
	Nsfw        bool   `json:"nsfw"`
	Vote        string `json:"vote"`
	InGallery   bool   `json:"in_gallery"`
}

type ImageResponse struct {
	Data    Image
	Status  int
	Success bool
}

type ImagesResponse struct {
	Data    []Image
	Status  int
	Success bool
}

func (cl *ImgurClient) GetAccountImages(page int) ([]Image, error) {
	isr := ImagesResponse{}
	request, _ := cl.prepareRequest("GET", fmt.Sprintf("account/%s/images/%d", cl.AccountUsername, page))
	fmt.Println(request)
	response, err := cl.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &isr)
	if err != nil {
		return []Image{}, err
	}
	return isr.Data, err
}
