package richcontent

import (
	"bytes"
	"regexp"
)

var (
	urlPattern  = regexp.MustCompile(`(?:^|[^a-zA-Z0-9\-_])(https?://[!-~]+)`)
	httpLiteral = []byte("http")
)

const (
	urlPatternSubmatchLeadIndex    = 0
	urlPatternSubmatchContentIndex = 1
)

func FindAllUrls(input []byte) [][]byte {
	// Fast path.
	if !bytes.Contains(input, httpLiteral) {
		return nil
	}

	matches := urlPattern.FindAllSubmatch(input, -1)
	if len(matches) == 0 {
		return nil
	}

	urls := make([][]byte, len(matches))
	for i := range matches {
		urls[i] = matches[i][urlPatternSubmatchContentIndex]
	}
	return urls
}

// [[start, end], [start, end], ...]
func FindAllUrlsIndex(input []byte) [][]int {
	// Fast path.
	if !bytes.Contains(input, httpLiteral) {
		return nil
	}

	matches := urlPattern.FindAllSubmatchIndex(input, -1)
	for i := range matches {
		matches[i] = matches[i][2*urlPatternSubmatchContentIndex : 2*urlPatternSubmatchContentIndex+2]
	}
	return matches
}
