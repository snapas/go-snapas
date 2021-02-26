package snapas

import (
	"html/template"
	"time"
)

// Album represents a Snap.as photo album / gallery.
type Album struct {
	Created  time.Time      `json:"created"`
	Title    string         `json:"title"`
	Body     *string        `json:"body"`
	HTMLBody *template.HTML `json:"html_body"`
	Alias    string         `json:"alias"`
	Views    int64          `json:"views"`
	Photos   []Photo        `json:"photos"`
}
