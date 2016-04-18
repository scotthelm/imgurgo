package v3

type GalleryImage struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Datetime     int64  `json:"datetime"`
	Type         string `json:"type"`
	Animated     bool   `json:"animated"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	Size         int64  `json:"size"`
	Views        int64  `json:"views"`
	Bandwidth    int64  `json:"bandwidth"`
	Deletehash   string `json:"deletehash"`
	Link         string `json:"link"`
	Gifv         string `json:"gifv"`
	Mp4          string `json:"mp4"`
	Webm         string `json:"webm"`
	Mp4Size      int64  `json:"mp4_size"`
	WebmSize     int64  `json:"webm_size"`
	Looping      bool   `json:"looping"`
	Vote         string `json:"vote"`
	Favorite     bool   `json:"favorite"`
	Nsfw         bool   `json:"nsfw"`
	CommentCount int    `json:"comment_count"`
	Topic        string `json:"topic"`
	TopicId      int64  `json:"topic_id"`
	Section      string `json:"section"`
	AccountUrl   string `json:"account_url"`
	AccountId    int64  `json:"account_id"`
	Ups          int64  `json:"ups"`
	Downs        int64  `json:"downs"`
	Points       int64  `json:"points"`
	Score        int64  `json:"score"`
	IsAlbum      bool   `json:"is_album"`
}

type GalleryProfile struct {
	TotalGalleryComments    int64    `json:"total_gallery_comments"`
	TotalGalleryFavorites   int64    `json:"total_gallery_favorites"`
	TotalGallerySubmissions int64    `json:"total_gallery_submissions"`
	Trophies                []Trophy `json:"trophies"`
}

type Trophy struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	NameClean   string `json:"name_clean"`
	Description string `json:"description"`
	Data        string `json:"data"`
	DataLink    string `json:"data_link"`
	Datetime    int64  `json:"datetime"`
	Image       string `json:"image"`
}
