package param

import (
	"../goVirtualHost"
	"../util"
	"crypto/tls"
	"regexp"
	"strings"
)

func LoadCertificate(certFile, keyFile string) (*tls.Certificate, error) {
	return goVirtualHost.LoadCertificate(certFile, keyFile)
}

func EntriesToUsers(entries []string) []*user {
	users := make([]*user, 0, len(entries))
	for _, userEntry := range entries {
		username := userEntry
		password := ""

		colonIndex := strings.IndexByte(userEntry, ':')
		if colonIndex >= 0 {
			username = userEntry[:colonIndex]
			password = userEntry[colonIndex+1:]
		}

		users = append(users, &user{username, password})
	}
	return users
}

func WildcardToRegexp(wildcards []string, found bool) (*regexp.Regexp, error) {
	if !found || len(wildcards) == 0 {
		return nil, nil
	}

	normalizedWildcards := make([]string, 0, len(wildcards))
	for _, wildcard := range wildcards {
		if len(wildcard) == 0 {
			continue
		}
		normalizedWildcards = append(normalizedWildcards, util.WildcardToRegexp(wildcard))
	}

	if len(normalizedWildcards) == 0 {
		return nil, nil
	}

	exp := strings.Join(normalizedWildcards, "|")
	return regexp.Compile(exp)
}
