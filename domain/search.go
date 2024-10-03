package domain

import (
	"fmt"
	"github.com/likexian/whois"
	"strings"
)

type DomainInfo struct {
	OwnerName      string
	Registrar      string
	CreationDate   string
	ExpirationDate string
	Status         string
}

func Search(domain string) (*DomainInfo, error) {
	// Perform a WHOIS lookup for the domain
	whoisData, err := whois.Whois(domain)
	if err != nil {
		return nil, fmt.Errorf("failed to get WHOIS info: %w", err)
	}

	// Parse the WHOIS data
	ownerName := extractField(whoisData, "Registrant Name")
	registrar := extractField(whoisData, "Registrar")
	creationDate := extractField(whoisData, "Creation Date")
	expirationDate := extractField(whoisData, "Expiration Date")
	status := extractField(whoisData, "Domain Status")

	// Return the extracted information
	return &DomainInfo{
		OwnerName:      ownerName,
		Registrar:      registrar,
		CreationDate:   creationDate,
		ExpirationDate: expirationDate,
		Status:         status,
	}, nil
}

// Helper function to extract fields from WHOIS data
func extractField(data string, fieldName string) string {
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if strings.Contains(line, fieldName) {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}
