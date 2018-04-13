package main

// Configs to store configs from local text file
type Configs struct {
	IP                 string
	Port               string
	StorageDirectory   string
	TokenSigningSecret string
	UploadFormField    string
	// Note: these would normally be strings which would be parsed to int
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	MaxUploadSize       int64
}

// Note: should get configs from local storage but harcoding here for now
func getConfigs() *Configs {
	return &Configs{
		IP:                  "127.0.0.1",
		Port:                "9999",
		StorageDirectory:    "./storage/",
		TokenSigningSecret:  "secret",
		UploadFormField:     "file",
		ReadTimeoutSeconds:  15,
		WriteTimeoutSeconds: 15,
		MaxUploadSize:       10 * 1024,
	}
}
