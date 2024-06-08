package usernamecheck

import (
	"os"
	"strings"
)

func CheckForBlacklistedNames() bool {
	bn := []string{"Johnson", "Miller", "malware", "maltest", "CurrentUser", "Sandbox", "virus", "John Doe", "test user", "sand box", "WDAGUtilityAccount"}
	user := strings.ToLower(os.Getenv("USERNAME"))
	for _, bn := range bn {
		if user == strings.ToLower(bn) {
			os.Exit(-1)
		}
	}
	return false
}
