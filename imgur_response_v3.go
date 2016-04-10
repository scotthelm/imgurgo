package main

type ImgurResponseV3 struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Status  int         `json:"status"`
}
