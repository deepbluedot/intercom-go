package intercom

type App struct {
	Type      string `json:"type"`
	ID        string `json:"id_code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	TimeZone  string `json:"timezone"`
	Region    string `json:"region"`
}
