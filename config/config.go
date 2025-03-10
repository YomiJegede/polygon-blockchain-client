// Application configuration, making it easier to modify settings like the Polygon RPC URL.
package config

import (
	"os"
)

const defaultPolygonRPCURL = "https://polygon-rpc.com/"

// GetPolygonRPCURL returns the Polygon RPC URL from environment variables.
func GetPolygonRPCURL() string {
	url := os.Getenv("POLYGON_RPC_URL")
	if url == "" {
		return defaultPolygonRPCURL
	}
	return url
}
