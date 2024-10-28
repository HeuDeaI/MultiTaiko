package rawencoding_test

import (
	"Taiko/pkg/rawencoding"
	"testing"
)

func TestStringRlpEcnode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0x1231", "821231"},
		{"0xaf23", "82af23"},
		{"0xaf1353126489713abcdefeedde23", "8eaf1353126489713abcdefeedde23"},
		{"0x55", "55"},
		{"", "80"},
		{"0x", "80"},
	}

	for _, test := range tests {
		result := rawencoding.StringRlpEcnode(test.input)
		if result != test.expected {
			t.Errorf("stringRlpEcnode(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestListRlpEcnode(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"0xc4", "0x82", "0x1234", "0x82", "0x5678"}, "cc81c481828212348182825678"},
		{[]string{"0x11", "0xafa41234", "0x7f", "0x5afaaa1241"}, "cd1184afa412347f855afaaa1241"},
		{[]string{}, "c0"},
	}

	for _, test := range tests {
		result := rawencoding.ListRlpEcnode(test.input)
		if result != test.expected {
			t.Errorf("listRlpEcnode(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
