package snapas

import "time"

// Photo represents a photo on Snap.as.
type Photo struct {
	ID       string    `json:"id"`
	Created  time.Time `json:"created"`
	Body     *string   `json:"body"`
	Filename string    `json:"filename"`
	Size     int64     `json:"size"`
	URL      string    `json:"url"`
	Album    *Album    `json:"album"`
}
