package main

// Config to hold conf
type Config struct {
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