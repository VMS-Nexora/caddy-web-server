package main

type Version struct {
	Major     int    `json:"major"`
	Minor     int    `json:"minor"`
	Build     int    `json:"build"`
	Value     int    `json:"value"`
	Date      string `json:"date"`
	HashSHA256 string `json:"hash_sha256"`
}

type EndpointConfig struct {
	Base struct {
		URL string `json:"url"`
	} `json:"base"`
}

type S3StorageConfig struct {
	URL string `json:"url"`
}

type PluginConfig struct {
	Enabled bool `json:"enabled"`

	AdditionalFields map[string]struct {
		Type    string      `json:"type"`
		Value   interface{} `json:"value"`
		Comment string      `json:"comment,omitempty"`
	} `json:"-"`
}

type SettingsConfig struct {
	EnableLogging struct {
		Type    string `json:"type"`
		Value   bool   `json:"value"`
		Comment string `json:"comment"`
	} `json:"enable_logging"`
	MaxRetryAttempts struct {
		Type    string `json:"type"`
		Value   int    `json:"value"`
		Comment string `json:"comment"`
	} `json:"max_retry_attempts"`
	TenantCode *struct {
		Type    string `json:"type"`
		Value   int    `json:"value"`
		Comment string `json:"comment"`
	} `json:"tenant_code,omitempty"`
}

type GlobalConfig struct {
	Schema    string         `json:"$schema"`
	Version   Version        `json:"version"`
	Endpoints EndpointConfig `json:"endpoints"`
	Plugins   map[string]PluginConfig `json:"plugins"`
	Settings  SettingsConfig `json:"settings"`
	S3Storage *S3StorageConfig `json:"s3_storage,omitempty"`
}

type TenantConfig struct {
	TenantCode struct {
		Type    string `json:"type"`
		Value   int    `json:"value"`
		Comment string `json:"comment"`
	} `json:"tenant_code"`
}
