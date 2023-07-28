package umago

import (
	"net/http"
	"strings"
)

// Event represents Umami tracking event
type Event struct {
	Hostname string `json:"hostname"`
	Language string `json:"language"`
	Referer  string `json:"referrer"`
	Screen   string `json:"screen"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Website  string `json:"website"`
	Name     string `json:"name"`
}

// NewEventFromHttpRequest allows to gather event information,
// based on provided request.
func NewEventFromHttpRequest(r *http.Request) Event {
	// Resolve language
	language := r.Header.Get("Accept-Language")
	if language == "*" {
		language = ""
	}
	if language != "" {
		language = strings.Split(language, ",")[0]
	}
	// Compose event
	return Event{
		Hostname: r.Host,
		Language: language,
		Referer:  r.Referer(),
		Url:      r.URL.String(),
	}
}
