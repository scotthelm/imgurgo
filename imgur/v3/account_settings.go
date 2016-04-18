package v3

type AccountSettings struct {
	AccountUrl           string        `json"account_url"`
	Email                string        `json"email"`
	HighQuality          bool          `json"high_quality"`
	PublicImages         bool          `json"public_images"`
	AlbumPrivacy         string        `json"album_privacy"`
	ProExpiration        int           `json"pro_expiration"`
	AcceptedGalleryTerms bool          `json"accepted_gallery_terms"`
	ActiveEmails         []string      `json"active_emails"`
	MessagingEnabled     bool          `json"messaging_enabled"`
	BlockedUsers         []BlockedUser `json"blocked_users"`
}

type AccountSettingsResponse struct {
	Data    AccountSettings
	Status  int
	Success bool
}

type BlockedUser struct {
	BlockedId  int64  `json:"blocked_id"`
	BlockedUrl string `json:"blocked_url"`
}
