package intercom

type NoteList struct {
	Notes []Note `json:"data,omitempty"`
	Count int64  `json:"total_count,omitempty"`
}

type Note struct {
	Type   string `json:"type,omitempty"`
	ID     string `json:"id,omitempty"`
	Author *Admin `json:"author,omitempty"`
	Body   string `json:"body,omitempty"`
}
