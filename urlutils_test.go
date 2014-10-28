package urlutils

import (
	"net/url"
	"testing"
)

var isAssetTests = []struct {
	link     string
	expected bool
}{
	{"https://www.example.com/assets/default-3a228ac6db5e7a521d6442dc37d476f8.css", true},
	{"https://www.example.js/assets/defaultcss", false},
	{"https://www.example/image.GIF", true},
	{"https://www.example/image.jpeg", true},
	{"https://www.example/image.jpg", true},
}

func TestIsAsset(t *testing.T) {
	for _, test := range isAssetTests {
		url, _ := url.Parse(test.link)
		got := IsAsset(url)
		if got != test.expected {
			t.Errorf("Expected - %q for link: %q", test.expected, test.link)
		}
	}
}

var isRelativeTests = []struct {
	link     string
	expected bool
}{
	{"http://www.example.com", false},
	{"//www.example.com", false},
	{"/news/29292-article.html", true},
}

func TestIsRelative(t *testing.T) {
	for _, test := range isRelativeTests {
		url, _ := url.Parse(test.link)
		got := IsRelative(url)
		if got != test.expected {
			t.Errorf("Expected - %q for link: %q", test.expected, test.link)
		}
	}
}

var isAbsoluteTests = []struct {
	link     string
	expected bool
}{
	{"http://www.example.com/some.html", true},
	{"//www.example.com", false},
	{"/news/29292-article.html", false},
}

func TestIsAbsolute(t *testing.T) {
	for _, test := range isAbsoluteTests {
		url, _ := url.Parse(test.link)
		got := IsAbsolute(url)
		if got != test.expected {
			t.Errorf("Expected - %q for link: %q", test.expected, test.link)
		}
	}
}

var addWWWTests = []struct {
	link         string
	expectedLink string
}{
	{"http://example.com", "http://www.example.com"},
	{"www.anohter.com", "http://www.another.com"},
}

func TestAddWWW(t *testing.T) {
	for _, test := range addWWWTests {
		u, _ := url.Parse(test.link)
		got := AddWWW(u)
		if u != got {
			t.Errorf("Expected - %q != %q", u.String(), got.String())
		}
	}
}

func TestAddHTTP(t *testing.T) {
	for _, test := range addWWWTests {
		u, _ := url.Parse(test.link)
		got := AddHTTP(u)
		if u != got {
			t.Errorf("Expeced to add Scheme for %q", u)
		}

	}

}

var normalizeDomainTests = []struct {
	url    string
	domain string
}{
	{"http://www.example.com", "http://www.example.com"},
	{"http://www.sub.sub.example.com", "http://www.example.com"},
	{"http://dom-sub.o.com", "http://o.com"},
}

func TestNormalizeDomain(t *testing.T) {
	for _, test := range normalizeDomainTests {
		u, _ := url.Parse(test.url)
		got, _ := NormalizeDomain(u)
		expected, _ := url.Parse(test.domain)
		if got == expected {
			t.Errorf("Expected %q got %q", test.domain, got)
		}

	}
}
