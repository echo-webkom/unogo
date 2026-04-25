package unogo

import (
	"encoding/json"
	"fmt"
	"net/http"
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

var globalHttpClient = &http.Client{}

type httpClient struct {
	*http.Client
}

func newHttpClient() *httpClient {
	// TODO: apply AdminKey
	return &httpClient{globalHttpClient}
}

type Client struct {
	config *Config
	client *httpClient
}

// NewClient creates and configures a new Uno client. [config] may be nil, in which case default
// configuration is used.
func NewClient(config *Config) *Client {
	if config == nil {
		config = &defaultConfig
	}

	client := &Client{
		config: config,
		client: newHttpClient(),
	}

	return client
}

// Health queries Unos health-check endpoint.
func (c *Client) Health() (HealthResponse, error) {
	res, err := c.client.Get(c.config.Url)
	if err != nil {
		return HealthResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return HealthResponse{}, unoError(res.StatusCode, "failed to get health status")
	}

	var healthResp HealthResponse
	err = json.NewDecoder(res.Body).Decode(&healthResp)
	return healthResp, err
}

// GetUser fetches the user info for the user the given session token belongs to.
func (c *Client) GetUser(sessionToken string) (UserResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/me", c.config.Url), nil)
	if err != nil {
		return UserResponse{}, err
	}

	req.Header.Add("Authorization", "Bearer "+sessionToken)

	res, err := c.client.Do(req)
	if err != nil {
		return UserResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return UserResponse{}, unoError(res.StatusCode, "failed to get user")
	}

	var user UserResponse
	err = json.NewDecoder(res.Body).Decode(&user)
	return user, err
}

func unoError(code int, format string, args ...any) error {
	return fmt.Errorf("uno responded with non-success code (%d): %s", code, fmt.Sprintf(format, args...))
}
