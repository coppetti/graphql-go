package data

// Album Structure
type Album struct {
	ID     string `json:"id,omitempty"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
	Type   string `json:"type"`
}

// Artist structure
type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Song structure
type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}
