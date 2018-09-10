package snapas

import (
	"code.as/core/api"
)

const (
	apiURL = "https://snap.as/api"
)

func NewClient() *as.ClientConfig {
	return as.NewClientConfig(apiURL, "go-snapas v1")
}
