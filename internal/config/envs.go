package config

import "os"

// Env is a struct that holds the environment variables.

func GetDataBaseDNS() string {
	dns := os.Getenv("POSTGRES_URL")
	if dns == "" {
		dns = "postgresql://alexandre:lm0NHxneRYtIpxziP7y4Tg@easysearch-9486.7tt.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	}

	return dns
}
