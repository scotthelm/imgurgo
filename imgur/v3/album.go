package v3

type Album struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Datetime    int64   `json:"datetime"`
	Cover       string  `json:"cover"`
	CoverWidth  int64   `json:"cover_width"`
	CoverHeight int64   `json:"cover_height"`
	AccountUrl  string  `json:"account_url"`
	AccountId   int     `json:"account_id"`
	Privacy     string  `json:"privacy"`
	Layout      string  `json:"layout"`
	Views       int64   `json:"views"`
	Link        string  `json:"link"`
	Favorite    bool    `json:"favorite"`
	Nsfw        bool    `json:"nsfw"`
	Section     string  `json:"section"`
	Order       int64   `json:"order"`
	Deletehash  string  `json:"deletehash"`
	ImagesCount int64   `json:"images_count"`
	Images      []Image `json:"images"`
	InGallery   bool    `json:"in_gallery"`
}

type AlbumResponse struct {
	Data    Album
	Status  int
	Success bool
}
