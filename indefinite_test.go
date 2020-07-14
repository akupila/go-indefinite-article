package indefinite_test

import (
	"fmt"
	"testing"

	"github.com/akupila/indefinite"
)

func TestArticle(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		// Empty
		{"", "a"},

		// https://github.com/Kaivosukeltaja/php-indefinite-article/blob/master/test.php
		{"umbrella", "an"},
		{"hour", "an"},
		{"American", "an"},
		{"German", "a"},
		{"Ukrainian", "an"},
		{"Uzbekistani", "an"},
		{"euphenism", "a"},
		{"Euler", "an"},

		// https://grammar.yourdictionary.com/vs/a-vs-an-basic-rules-and-exceptions.html
		{"Restaurant", "a"},   // The next sound is r-
		{"Man", "a"},          // The next sound is m-
		{"Dog", "a"},          // The next sound is d-
		{"book", "a"},         // The next sound is b-
		{"purple", "a"},       // The next sound is p-
		{"alligator", "an"},   // The next sound is uh-
		{"ice cream", "an"},   // the next sound is eye-
		{"Eager child", "an"}, // the next sound is ee-
		{"orange cup", "an"},  // The next sound is oh-
		{"Hourglass", "an"},   // First sound is ow-
		{"Horse", "a"},        // First sound is h-
		{"Honor", "an"},       // First sound is aw-
		{"Heir", "an"},        // First sound is eh-
		{"Hair", "a"},         // First sound is h-
		// Words that make a You- Sound
		{"European trip", "a"},
		{"ewe", "a"},
		{"union", "a"},
		{"URL", "a"}, // Not working
		{"unicorn", "a"},
		{"eulogy", "a"},
		// Acronyms
		{"GRE", "a"}, // First sound is gee-
		{"DOJ", "a"}, // First sound is dee-

		// https://www.bespeaking.com/a-or-an-the-rules-and-exceptions/
		{"Cat", "a"},
		{"Book", "a"},
		{"Show", "a"},
		{"TV program", "a"},
		{"University", "a"},       // "you"-niversity
		{"unique situation", "a"}, // "you"-nique
		{"united country", "a"},   // "you"-nited
		{"history class", "a"},
		{"hospital", "a"},
		{"helpful hand", "a"},
		{"actor", "an"},
		{"example", "an"},
		{"invitation", "an"},
		{"unbrella", "an"},
		{"hour", "an"}, // silent/soft h
		{"F", "an"},    // efff

		// Numbers
		{"8-ball", "an"},
		{"11th", "an"},

		// Ordinals
		{"t-th", "a"},
		{"n-th", "an"},

		// Misc
		{"T", "a"},

		// Hyphen
		{"x-mas", "an"},
		{"b-tree", "a"},

		// Special vowel form
		{"one", "a"},
		{"unit", "a"},
		{"utter", "an"},
		{"ufo", "a"},

		// Starting with y
		{"ypsilon", "an"},

		// Test cases that ideally would work:
		// {"MBA", "an"}, // First sound is em-
		// {"NFL", "an"}, // First sound is en-
		// {"SOS", "an"}, // First sound is es-
	}

	for _, tc := range tests {
		t.Run(tc.str, func(t *testing.T) {
			got := indefinite.Article(tc.str)
			if got != tc.want {
				t.Errorf(`Got "%s %s", want "%s %s"`, got, tc.str, tc.want, tc.str)
			}
		})
	}
}

func ExampleArticle() {
	words := []string{
		"school",
		"ukelele",
		"ace of spaced",
		"uniform",
		"igloo",
		"dinosaur",
		"air-conditioner",
		"egg",
		"xylophone",
	}

	for _, word := range words {
		article := indefinite.Article(word)
		fmt.Printf("%s %s\n", article, word)
	}
	// Output:
	// a school
	// a ukelele
	// an ace of spaced
	// a uniform
	// an igloo
	// a dinosaur
	// an air-conditioner
	// an egg
	// a xylophone
}
