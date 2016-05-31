package engine

import (
	"strings"
	"github.com/dchest/stemmer/porter2"
	"log"
)

// TokenizeSentence returns a sentence broken into tokens. Tokens are individual
// words as well as punctuation. For example, "Hi! How are you?" becomes
// []string{"Hi", "!", "How", "are", "you", "?"}.
func TokenizeSentence(sent string) []string {
	tokens := []string{}
	for _, w := range strings.Fields(sent) {
		found := []int{}
		for i, r := range w {
			switch r {
			case '\'', '"', ':', ';', '!', '?':
				found = append(found, i)

			// Handle case of currencies and fractional percents.
			case '.', ',':
				if i+1 < len(w) {
					switch w[i+1] {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						continue
					}
				}
				found = append(found, i)
				i++
			}
		}
		if len(found) == 0 {
			tokens = append(tokens, w)
			continue
		}
		for i, j := range found {
			// If the token marker is not the first character in the
			// sentence, then include all characters leading up to
			// the prior found token.
			if j > 0 {
				if i == 0 {
					tokens = append(tokens, w[:j])
				} else if i-1 < len(found) {
					// Handle case where multiple tokens are
					// found in the same word.
					tokens = append(tokens, w[found[i-1]+1:j])
				}
			}

			// Append the token marker itself
			tokens = append(tokens, string(w[j]))

			// If we're on the last token marker, append all
			// remaining parts of the word.
			if i+1 == len(found) {
				tokens = append(tokens, w[j+1:])
			}
		}
	}
	log.Println("found tokens", tokens)
	return tokens
}

// StemTokens returns the porter2 (snowball) stems for each token passed into
// it.
func StemTokens(tokens []string) []string {
	eng := porter2.Stemmer
	stems := []string{}
	for _, w := range tokens {
		if len(w) == 1 {
			switch w {
			case "'", "\"", ",", ".", ":", ";", "!", "?":
				continue
			}
		}
		w = strings.ToLower(w)
		stems = append(stems, eng.Stem(w))
	}
	return stems
}
