package urlutils

import (
	"net/url"
	"regexp"
)

const (
	assetRegex = "\\.png|\\.jpe??g|\\.gif|\\.bmp|\\.psd|\\.js|\\.json|\\.css"
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
	return r.MatchString(link.Path)
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
