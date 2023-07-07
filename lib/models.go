package lib

import "time"

type Server struct {
	Hosts                []string `json:"hosts"`
	ValidationPercentage int      `json:"validationPercentage"`
	HealthInterval       int      `json:"healthInterval"`
	HealthTimeout        int      `json:"healthTimeout"`
	HealthRetries        int      `json:"healthRetries"`
	SecureKey            string   `json:"secureKey"`
	// Consider adding a healthcheck key to the server struct
	// Update the healthcheck key every time a server is dropped
}

type Host struct {
	IPAddress   string    `json:"ipAddress"`
	Port        string    `json:"port"`
	Alive       bool      `json:"alive"`
	LastChecked time.Time `json:"lastChecked"`
}
