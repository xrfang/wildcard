package wildcard

import (
	"strings"
	"unicode/utf8"
)

type Pattern string

func (p Pattern) Match(subject string) bool {
next:
	for len(p) > 0 {
		var star bool
		var chunk string
		star, chunk, p = scanChunk(p)
		if star && chunk == "" {
			return true // Trailing * matches rest of string.
		}
		// Look for match at current position.
		t, ok := matchChunk(chunk, subject)
		// if we're the last chunk, make sure we've exhausted the subject
		// otherwise we'll give a false result even if we could still match
		// using the star
		if ok && (len(t) == 0 || len(p) > 0) {
			subject = t
			continue
		}
		if star {
			// Look for match skipping i+1 bytes.
			for i := 0; i < len(subject); i++ {
				t, ok := matchChunk(chunk, subject[i+1:])
				if ok {
					// if we're the last chunk, make sure we exhausted the subject
					if len(p) == 0 && len(t) > 0 {
						continue
					}
					subject = t
					continue next
				}
			}
		}
		return false
	}
	return len(subject) == 0
}

func (p Pattern) LowerCaseMatch(subject string) bool {
	return p.Match(strings.ToLower(subject))
}

func (p Pattern) UpperCaseMatch(subject string) bool {
	return p.Match(strings.ToUpper(subject))
}

// scanChunk gets the next segment of pattern, which is a non-star
// string possibly preceded by a star.
func scanChunk(pattern Pattern) (star bool, chunk string, rest Pattern) {
	for len(pattern) > 0 && pattern[0] == '*' {
		pattern = pattern[1:]
		star = true
	}
	var i int
	for i = 0; i < len(pattern); i++ {
		if pattern[i] == '*' {
			break
		}
	}
	return star, string(pattern[0:i]), pattern[i:]
}

// matchChunk checks whether chunk matches the beginning of s.
// If so, it returns the remainder of s (after the match).
// Chunk is all single-character operators: literals, char classes, and ?.
func matchChunk(chunk, s string) (rest string, ok bool) {
	for len(chunk) > 0 {
		if len(s) == 0 {
			return
		}
		switch chunk[0] {
		case '?':
			_, n := utf8.DecodeRuneInString(s)
			s = s[n:]
			chunk = chunk[1:]
		default:
			if chunk[0] != s[0] {
				return
			}
			s = s[1:]
			chunk = chunk[1:]
		}
	}
	return s, true
}
