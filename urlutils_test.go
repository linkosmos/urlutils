package urlutils

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var isAssetTests = []struct {
	link     string
	expected bool
}{
	{"https://www.example.com/assets/default-3a228ac6db5e7a521d6442dc37d476f8.css", true},
	{"https://www.example.io/assets/defaultcss", false},
	{"https://www.example/image.jpeg", true},
	{"https://www.example/image.jpg", true},
	{"https://www.example/file.pdf", true},
	{"https://www.example/file.sass", true},
	{"https://www.example/file.js", true},
	{"https://www.example/file.jsx", true},
	{"https://www.example/file.jsz", false},
	{"https://www.example/file.scss", true},
}

func TestIsAsset(t *testing.T) {
	for _, test := range isAssetTests {
		if got := IsAsset(test.link); got != test.expected {
			t.Errorf("Expected - %t for link: %q", test.expected, test.link)
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
			t.Errorf("Expected - %t for link: %q", test.expected, test.link)
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
			t.Errorf("Expected - %t for link: %q", test.expected, test.link)
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

var reverseDomainTests = []struct {
	url      string
	expected string
}{
	{"http://www.example.com", "com.example.www"},
	{"http://example", ""},
	{"example.com", ""},
}

func TestReverseDomain(t *testing.T) {
	for _, test := range reverseDomainTests {
		u, _ := url.Parse(test.url)
		got, _ := ReverseDomain(u)
		if got != test.expected {
			t.Errorf("Expected %q got %q", test.expected, got)
		}
	}

}

var splitPathTests = []struct {
	url      string
	depth    int
	expected string
}{
	{"http://www.example.com/section/here", 1, "/section"},
	{"http://www.example.com/second/here/more?para=22", 3, "/second/here/more"},
	{"http://www.example.com/here/now/params?query=22#fragment", 0, ""},
	{"http://www.example.com/section/here", -13, ""},
	{"http://www.example.com/section/here/where/path/is/long?query=2#ad", 5, "/section/here/where/path/is"},
}

func TestSplitPath(t *testing.T) {
	for _, test := range splitPathTests {
		u, _ := url.Parse(test.url)
		got, _ := SplitPath(u, test.depth)
		if got != test.expected {
			t.Errorf("Expected %q got %q", test.expected, got)
		}
	}
}

var hosttldTests = []struct {
	url      string
	expected string
}{
	{"http://www.example.io/section/here", "io"},
	{"http://www.example.com/second/here/more?para=22", "com"},
	{"http://www.example.dance/here/now/params?query=22#fragment", "dance"},
}

func TestHostTLD(t *testing.T) {
	for _, test := range hosttldTests {
		u, _ := url.Parse(test.url)
		got, _ := HostTLD(u)
		if got != test.expected {
			t.Errorf("Expected %q got %q", test.expected, got)
		}
	}
}

var isRootTests = []struct {
	url      string
	expected bool
}{
	{"http://www.example.io", false},
	{"http://www.example.io/", true},
	{"http://www.example.io#named_link", false},
	{"http://www.example.io?some=parm", false},
	{"http://www.example.io/section/here", false},
	{"http://www.example.com/second/here/more?para=22", false},
	{"http://www.example.dance/here/now/params?query=22#fragment", false},
}

func TestIsHomePage(t *testing.T) {
	for _, test := range isRootTests {
		u, _ := url.Parse(test.url)
		got := IsHomePage(u)
		if got != test.expected {
			t.Errorf("Expected %t got %t", test.expected, got)
		}
	}
}

var isEmptyQueryTests = []struct {
	url      string
	expected bool
}{
	{"http://www.example.io", true},
	{"http://www.example.io/", true},
	{"http://www.example.io#named_link", false},
	{"http://www.example.io?some=parm", false},
	{"http://www.example.io/section/here", true},
	{"http://www.example.com/second/here/more?para=22", false},
	{"http://www.example.dance/here/now/params?query=22#fragment", false},
}

func TestEmptyQuery(t *testing.T) {
	for _, test := range isEmptyQueryTests {
		u, _ := url.Parse(test.url)
		got := IsEmptyQuery(u)

		assert.Equal(t, test.expected, got,
			fmt.Sprintf("Expected %t for %s", test.expected, test.url))
	}
}

var isEmptyPathTests = []struct {
	url      string
	expected bool
}{
	{"http://www.example.io", true},
	{"http://www.example.io/", false},
	{"http://www.example.io#named_link", true},
	{"http://www.example.io?some=parm", true},
	{"http://www.example.io/section/here", false},
	{"http://www.example.com/second/here/more?para=22", false},
	{"http://www.example.dance/here/now/params?query=22#fragment", false},
}

func TestEmptyPath(t *testing.T) {
	for _, test := range isEmptyPathTests {
		u, _ := url.Parse(test.url)
		got := IsEmptyPath(u)

		assert.Equal(t, test.expected, got,
			fmt.Sprintf("Expected %t for %s", test.expected, test.url))
	}
}

var isPlainTests = []struct {
	url      string
	expected bool
}{
	{"http://www.example.io", true},
	{"http://www.example.io/", false},
	{"http://www.example.io#named_link", false},
	{"http://www.example.io?some=parm", false},
	{"http://www.example.io/section/here", false},
	{"http://www.example.com/second/here/more?para=22", false},
	{"http://www.example.dance/here/now/params?query=22#fragment", false},
}

func TestPlain(t *testing.T) {
	for _, test := range isPlainTests {
		u, _ := url.Parse(test.url)
		got := IsPlain(u)

		assert.Equal(t, test.expected, got,
			fmt.Sprintf("Expected %t for %s", test.expected, test.url))
	}
}
func TestIsNotPlain(t *testing.T) {
	for _, test := range isPlainTests {
		u, _ := url.Parse(test.url)
		got := IsNotPlain(u)

		// Expecting reverse
		assert.Equal(t, !test.expected, got,
			fmt.Sprintf("Expected %t for %s", test.expected, test.url))
	}
}
