package helpers

import (
	"os"
	"strings"
)

// EnforceHTTP adds the "http://" prefix to a URL if it doesn't already have it.
func EnforceHTTP(url string) string {
	// Check if the URL starts with "http" (already has the prefix)
	if url[:4] != "http" {
		// If not, add the "http://" prefix to the URL
		return "http://" + url
	}

	// Return the URL as is since it already has the "http" prefix
	return url
}

// RemoveDomainError checks if a URL is the same as the DOMAIN environment variable and removes domain-related errors.
func RemoveDomainError(url string) bool {
	// Check if the URL is the same as the DOMAIN environment variable
	if url == os.Getenv("DOMAIN") {
		// If they are the same, return false to indicate a domain error
		return false
	}

	// Remove "http://" and "https://" prefixes, remove "www.", and extract the domain part of the URL
	newUrl := strings.Replace(url, "http://", "", 1)
	newUrl = strings.Replace(newUrl, "https://", "", 1)
	newUrl = strings.Replace(newUrl, "www.", "", 1)
	newUrl = strings.Split(newUrl, "/")[0]

	// Check if the extracted domain matches the DOMAIN environment variable
	if newUrl == os.Getenv("DOMAIN") {
		// If they match, return false to indicate a domain error
		return false
	}

	// If no domain error is detected, return true
	return true
}
