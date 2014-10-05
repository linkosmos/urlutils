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
