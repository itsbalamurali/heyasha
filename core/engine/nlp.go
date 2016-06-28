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

/*
// classifier is a set of common english word stems unique among their
// Structured Input Types. This enables extremely fast constant-time O(1)
// lookups of stems to their SITs with high accuracy and no training
// requirements. It consumes just a few MB in memory.
type Classifier map[string]struct{}

// classifyTokens builds a StructuredInput from a tokenized sentence.
func (c Classifier) ClassifyTokens(tokens []string) *models.StructuredInput {
	var s models.StructuredInput
	var sections []string
	for _, t := range tokens {
		var found bool
		lower := strings.ToLower(t)
		_, exists := c["C"+lower]
		if exists {
			s.Commands = append(s.Commands, lower)
			found = true
		}
		_, exists = c["O"+lower]
		if exists {
			s.Objects = append(s.Objects, lower)
			found = true
		}


		// Each time we find an object, add a separator to sections,
		// enabling us to check for times only along continuous
		// stretches of a sentence (i.e. a single time won't appear on
		// either side of the word "Jim" or "Bring")
		if found || len(sections) == 0 {
			sections = append(sections, t)
		} else {
			switch t {
			case ".", ",", ";", "?", "-", "_", "=", "+", "#", "@",
				"!", "$", "%", "^", "&", "*", "(", ")", "'":
				continue
			}
			sections[len(sections)-1] += " " + t
		}
	}
	for _, sec := range sections {
		if len(sec) == 0 {
			continue
		}
		s.Times = append(s.Times, timeparse.Parse(sec)...)
	}
	return &s
}

// BuildClassifier prepares the Named Entity Recognizer (NER) to find Commands
// and Objects using a simple dictionary lookup. This has the benefit of high
// speed--constant time, O(1)--with insignificant memory use and high accuracy
// given false positives (marking something as both a Command and an Object when
// it's really acting as an Object) are OK. Ultimately this should be a first
// pass, and any double-marked words should be passed through something like an
// n-gram Bayesian filter to determine the correct part of speech within its
// context in the sentence.
func BuildClassifier() (Classifier, error) {
	ner := Classifier{}
	p := filepath.Join(os.Getenv("ABOT_PATH"), "data", "ner")
	fi, err := os.Open(filepath.Join(p, "nouns.txt"))
	if err != nil {
		return ner, err
	}
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ner["O"+scanner.Text()] = struct{}{}
	}
	if err = fi.Close(); err != nil {
		return ner, err
	}
	fi, err = os.Open(filepath.Join(p, "verbs.txt"))
	if err != nil {
		return ner, err
	}
	scanner = bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ner["C"+scanner.Text()] = struct{}{}
	}
	if err = fi.Close(); err != nil {
		return ner, err
	}
	fi, err = os.Open(filepath.Join(p, "adjectives.txt"))
	if err != nil {
		return ner, err
	}
	scanner = bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ner["O"+scanner.Text()] = struct{}{}
	}
	if err = fi.Close(); err != nil {
		return ner, err
	}
	fi, err = os.Open(filepath.Join(p, "adverbs.txt"))
	if err != nil {
		return ner, err
	}
	scanner = bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ner["O"+scanner.Text()] = struct{}{}
	}
	if err = fi.Close(); err != nil {
		return ner, err
	}
	fi, err = os.Open(filepath.Join(p, "names_female.txt"))
	if err != nil {
		return ner, err
	}
	scanner = bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ner["PF"+scanner.Text()] = struct{}{}
	}
	if err = fi.Close(); err != nil {
		return ner, err
	}
	fi, err = os.Open(filepath.Join(p, "names_male.txt"))
	if err != nil {
		return ner, err
	}
	scanner = bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ner["PM"+scanner.Text()] = struct{}{}
	}
	if err = fi.Close(); err != nil {
		return ner, err
	}
	return ner, nil
}*/