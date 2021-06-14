package translator

import (
	"errors"
	"strings"
)

type Translator struct {
	Word     string `json:"english-word"`
	Sentence string `json:"english-sentence"`
}

// Check if a character is a vowel or not.
func isVowel(character rune) bool {
	switch character {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

// Translate a single word into gopher.
func (translator *Translator) TranslateWord() (string, error) {
	// Return on an empty word.
	if translator.Word == "" {
		return "", errors.New("word not given")
	}

	// Return an error if a sentence was given.
	if len(strings.Fields(translator.Word)) > 1 {
		return "", errors.New("sentence given for word endpoint")
	}

	// Use the helper method to translate the word.
	translatedWord := TranslateSingleWord(translator.Word)

	return translatedWord, nil
}

// Translate a sentence into gopher.
func (translator *Translator) TranslateSentence() (string, error) {
	// Return on empty sentence.
	if translator.Sentence == "" {
		return "", errors.New("sentence not given")
	}

	// Split the sentence into separate words.
	words := strings.Split(translator.Sentence, " ")

	// Iterate over the words in the sentence and translate them
	// then add them to the result string.
	result := ""
	for _, word := range words {
		result += TranslateSingleWord(word) + " "
	}

	return strings.TrimSpace(result), nil
}

// Translate a single word from english to gopher.
func TranslateSingleWord(word string) string {
	if word[:2] == "xr" {
		return "ge" + word
	}

	// If we have a vowel we only need to add a prefix 'g'.
	if isVowel([]rune(word)[0]) {
		return "g" + word
	}

	// If the first letter is a consonant, iterate to find out how many
	// consonants are there.
	i := 0
	for !isVowel([]rune(word)[i]) {
		i++
	}

	// If we have a 'qu' then we need to a the 'u' to the suffix as well.
	if word[i-1:i+1] == "qu" {
		i++
	}

	// Return all the consonants as a suffix with 'ogo' suffix as well.
	return word[i:] + word[:i] + "ogo"
}
