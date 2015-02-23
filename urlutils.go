package urlutils

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

var (
	ErrURLHostMissing         = errors.New("Missing host in URL structure")
	ErrURLHostMalformed       = errors.New("URL Host is malformed")
	ErrURLPathMissing         = errors.New("Path is missing")
	ErrURLPathPartsEmpty      = errors.New("URL Path parts are empty")
	ErrURLPathDepthOutOfRange = errors.New("Depth out of Path parts range")

	assetRegex = regexp.MustCompile(`(?i)\.png|\.jpe??g|\.gif|\.bmp|\.psd|\.js|\.json|\.css|javascript`)
)

// ResolveURL - resolves relative to absolute URL
func ResolveURL(target *url.URL, relative *url.URL) *url.URL {
	return target.ResolveReference(relative)
}

// IsAsset - true if link is web asset
// e.g.: css or image
func IsAsset(link string) bool {
	return assetRegex.MatchString(link)
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
		u.Scheme = "http"
	}
	return u
}

// NormalizeDomain - parses Host and returns root domain
func NormalizeDomain(u *url.URL) (*url.URL, error) {
	if u.Host == "" {
		return nil, ErrURLHostMissing
	}
	domain := strings.Split(u.Host, ".")
	if len(domain) < 2 {
		return nil, ErrURLHostMalformed
	}
	prefix := ""
	if strings.HasPrefix(strings.ToLower(u.Host), "www.") {
		prefix = "www."
	}
	u.Host = prefix + strings.Join(domain[len(domain)-2:], ".")
	return u, nil
}

// ReverseDomain - reverses given URL Host
// e.g.: www.example.com => com.example.www
func ReverseDomain(u *url.URL) (string, error) {
	if u.Host == "" {
		return "", ErrURLHostMissing
	}
	domain := strings.Split(u.Host, ".")
	if len(domain) < 2 {
		return "", ErrURLHostMalformed
	}
	var reverseDomain []string

	for i := len(domain) - 1; i >= 0; i-- {
		reverseDomain = append(reverseDomain, domain[i])
	}

	return strings.Join(reverseDomain, "."), nil
}

// StripPathQueryFragment - strips URL path, query & fragment #
func StripPathQueryFragment(u *url.URL) *url.URL {
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""
	return u
}

// StripQueryFragment - strips URL query & fragment #
func StripQueryFragment(u *url.URL) *url.URL {
	u.RawQuery = ""
	u.Fragment = ""
	return u
}

// SplitPath - splits URL path into leveled segments
// lenParts  < 2 => 1 level deep
// lenParts == 2 => 2 level deep
// lenParts == 3 => 3 level deep
func SplitPath(u *url.URL, depth int) (string, error) {
	if u.Path == "" {
		return "", ErrURLPathMissing
	}
	parts := strings.Split(u.Path, "/")
	lenParts := (len(parts) - 1) // Golang right most add +1
	if lenParts == 0 {
		return "", ErrURLPathPartsEmpty
	}
	if depth < 0 || depth > lenParts {
		return "", ErrURLPathDepthOutOfRange
	}
	return strings.Join(parts[:depth+1], "/"), nil
}

// NormalizeURL - cleans params, adds www, insecures http scheme
func NormalizeURL(u *url.URL) (*url.URL, error) {
	u = StripPathQueryFragment(u)
	u = AddWWW(u)
	u.Scheme = "http"
	return NormalizeDomain(u)
}
