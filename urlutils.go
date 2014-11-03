package urlutils

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

const (
	assetRegex = "\\.png|\\.jpe??g|\\.gif|\\.bmp|\\.psd|\\.js|\\.json|\\.css|javascript"
)

// ResolveURL - resolves relative to absolute URL
func ResolveURL(target *url.URL, relative *url.URL) *url.URL {
	return target.ResolveReference(relative)
}

// IsAsset - true if link is web asset
// e.g.: css or image
func IsAsset(link *url.URL) bool {
	// Case insensitive regexp
	r, _ := regexp.Compile("(?i)" + assetRegex)
	return r.MatchString(link.String())
}

// IsRelative - true if link is relative
// e.g.: /news/article/29191.html
func IsRelative(link *url.URL) bool {
	return link.Host == "" && link.Scheme == ""
}

// IsAbsolute - true is link is absolute
// e.g.: http://www.example.com/news.html
func IsAbsolute(link *url.URL) bool {
	return link.Scheme != "" && link.Host != ""
}

// SameDomain - compares target and link domain host
func SameDomain(target *url.URL, link *url.URL) bool {
	return target.Host == link.Host
}

// AddWWW - adds www in front of given URL
func AddWWW(u *url.URL) *url.URL {
	if len(u.Host) > 0 && !strings.HasPrefix(strings.ToLower(u.Host), "www.") {
		u.Host = "www." + u.Host
	}
	return u
}

// AddHTTP - adds scheme for URL if missing
func AddHTTP(u *url.URL) *url.URL {
	if u.Scheme == "" {
		u.Scheme = "http://"
	}
	return u
}

// NormalizeDomain - parses Host and returns root domain
func NormalizeDomain(u *url.URL) (*url.URL, error) {
	if u.Host == "" {
		return nil, errors.New("Missing host in URL structure")
	}
	domain := strings.Split(u.Host, ".")
	if len(domain) < 2 {
		return nil, errors.New("URL Host is malformed")
	}
	prefix := ""
	if strings.HasPrefix(strings.ToLower(u.Host), "www.") {
		prefix = "www."
	}
	u.Host = prefix + strings.Join(domain[len(domain)-2:], ".")
	return u, nil
}

// StripParams - strip URL path, query & fragment #
func StripParams(u *url.URL) *url.URL {
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""
	return u
}

// ReverseDomain - reverses given URL Host
// e.g.: www.example.com => com.example.www
func ReverseDomain(u *url.URL) (string, error) {
	if u.Host == "" {
		return "", errors.New("Missing host in URL structure")
	}
	domain := strings.Split(u.Host, ".")
	if len(domain) < 2 {
		return "", errors.New("URL Host is malformed")
	}
	var reverseDomain []string

	for i := len(domain) - 1; i >= 0; i-- {
		reverseDomain = append(reverseDomain, domain[i])
	}

	return strings.Join(reverseDomain, "."), nil
}
