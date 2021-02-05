package main

// Config to hold conf
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

type authresp struct {
	StatusCode int `json:"statusCode"`
	Data struct {
		UserID string `json:"id"`
		AccessToken string `json:"access_token"`
	} `json:"data"`
}

type accountresp struct {
	StatusCode int `json:"statusCode"`
	Data struct {
		UserID string `json:"userId"`
		Accounts []struct {
			ID string `json:"custAccountId"`
		} `json:"linkedAccounts"`
	} `json:"data"`
}

type scheduleresp struct {
	ServiceID int `json:"serviceId"`
	NextPickup struct {
		Date string `json:"date"`
		Message string `json:"message"`
	} `json:"pickupDayInfo"`
}

