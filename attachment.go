package intercom

type Attachment struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"filesize"`
	Width       int64  `json:"width"`
	Height      int64  `json:"height"`
}
