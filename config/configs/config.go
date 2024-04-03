package configs

import (
	"encoding/json"
	"io"
)

var conf *Config

type Config struct {
	Env            string   `json:"env"`
	Port           string   `json:"port"`
	DatabaseURL    string   `json:"database_url"`
	MaxDBConn      int      `json:"max_db_conn"`
	ItemsPerPage   int      `json:"items_per_page"`
	YoutubeApiKeys []string `json:"youtube_api_keys"`
}

func ParseJSON(r io.Reader, v any) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func Get() *Config {
	return conf
}

func Set(c *Config) {
	conf = c
}
