package regex

import "testing"

func TestMatchOne(t *testing.T) {
	tests := []struct {
		pattern string
		text    string
		expect  bool
	}{
		{
			pattern: "a",
			text:    "a",
			expect:  true,
		},
		{
			pattern: ".",
			text:    "z",
			expect:  true,
		},
		{
			pattern: "",
			text:    "h",
			expect:  true,
		},
		{
			pattern: "a",
			text:    "b",
			expect:  false,
		},
		{
			pattern: "p",
			text:    "",
			expect:  false,
		},
	}

	for _, test := range tests {
		actual := matchOne(test.pattern, test.text)
		if test.expect != actual {
			t.Errorf("Expect %t but got %t\n", test.expect, actual)
			t.Errorf("pattern %s, text %s\n", test.pattern, test.text)
		}
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		pattern string
		text    string
		expect  bool
	}{
		{
			pattern: "abc",
			text:    "abc",
			expect:  true,
		},
		{
			pattern: "abc",
			text:    "kbc",
			expect:  false,
		},
		{
			pattern: ".bc",
			text:    "abc",
			expect:  true,
		},
		{
			pattern: "abc",
			text:    "abcd",
			expect:  true,
		},
	}

	for _, test := range tests {
		actual := match(test.pattern, test.text)
		if test.expect != actual {
			t.Errorf("Expect %t but got %t\n", test.expect, actual)
			t.Errorf("pattern %s, text %s\n", test.pattern, test.text)
		}
	}
}
