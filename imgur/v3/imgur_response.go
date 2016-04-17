package v3

type ImgurResponse struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Status  int         `json:"status"`
}
