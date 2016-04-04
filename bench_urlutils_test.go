package urlutils

import "testing"

func BenchmarkIsAsset(b *testing.B) {

	input := []string{
		"https://www.example.com/assets/default-3a228ac6db5e7a521d6442dc37d476f8.css",
		"https://www.example.io/assets/defaultcss",
		"https://www.example/image.GIF",
		"https://www.example/image.jpeg",
		"https://www.example/image.jpg",
		"https://www.example/image.JPG",
		"https://www.example/file.pdf",
		"https://www.example/file.sass",
		"https://www.example/file.js",
		"https://www.example/file.jsx",
		"https://www.example/file.jsz",
		"https://www.example/file.scss",
		"https://www.example/file.SCSS",
		"https://www.example/file.PDF",
		"https://www.example/image.JPG1",
	}

	for n := 0; n < b.N; n++ {
		for _, link := range input {
			IsAsset(link)
		}
	}
}
