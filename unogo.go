package unogo

import (
	"encoding/json"
)

type Config struct {
	// URL to Uno API
	Url string
	// AdminKey to be used for accessing Unos admin endpoints
	AdminKey string
}

// Default config to be used in the case of a nil-config in NewClient.
var defaultConfig = Config{
	Url:      "https://uno.echo-webkom.no",
	AdminKey: "",
}

type Client struct {
	config *Config
	client *httpClient
}

// NewClient creates and configures a new Uno client. [config] may be nil, in which case default
// configuration is used. An error is returned if given non-nil configuration is invalid or the
// client failed to connect to Uno.
func NewClient(config *Config) (*Client, error) {
	if config == nil {
		config = &defaultConfig
	}

	client := &Client{
		config: config,
		client: newHttpClient(),
	}

	return client, nil
}

type HealthResponse struct {
	Status string `json:"status"`
}

// Health queries Unos health-check endpoint.
func (c *Client) Health() (HealthResponse, error) {
	req, err := c.client.Get(c.config.Url)
	if err != nil {
		return HealthResponse{}, err
	}

	defer req.Body.Close()

	if req.StatusCode != 200 {
		return HealthResponse{}, unoError(req.StatusCode, "failed to get health status")
	}

	var healthResp HealthResponse
	err = json.NewDecoder(req.Body).Decode(&healthResp)
	return healthResp, err
}
