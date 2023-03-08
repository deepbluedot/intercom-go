package intercom

type Attachment struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"filesize"`
	Width       string `json:"width"`
	Height      string `json:"height"`
}
