package main

// Config to hold the full application config
type config struct {
	BaseURL string `yaml:"baseurl"`
	Auth struct {
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"auth"`
	APIKeys struct {
		Auth string `yaml:"auth"`
		Account string `yaml:"account"`
		Service string `yaml:"service"`
	} `yaml:"apikeys"`
}

// Authentication response model holding the Status code.
// AccessToken is a JWT token
type authresp struct {
	StatusCode int `json:"statusCode"`
	Data struct {
		UserID string `json:"id"`
		AccessToken string `json:"access_token"`
	} `json:"data"`
}

// Account response containing account ID, which is used for the Scheduler rest calls
type accountresp struct {
	StatusCode int `json:"statusCode"`
	Data struct {
		UserID string `json:"userId"`
		Accounts []struct {
			ID string `json:"custAccountId"`
		} `json:"linkedAccounts"`
	} `json:"data"`
}

// Scheduler response. The same model represents both Trash and Recycling service schedules
type scheduleresp struct {
	ServiceID int `json:"serviceId"`
	NextPickup struct {
		Date string `json:"date"`
		Message string `json:"message"`
	} `json:"pickupDayInfo"`
}

