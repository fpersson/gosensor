package libsensor

import (
	"testing"
)

func TestReadSensor(t *testing.T) {
	cases := []struct {
		in   string
		want float64
	}{
		{"apa", 0},
		{"../testdata/testfil", 2.562},
	}

	for _, c := range cases {
		got, _ := ReadSensor(c.in)
		if got != c.want {
			t.Errorf("ReadSensor(%q) == %f, want %f", c.in, got, c.want)
		}
	}
}

func TestParseValue(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		hasError bool
	}{
		{"line1\nvalue=12345", 12.345, false},
		{"line1\nvalue=abc", 0, true},
		{"line1", 0, true},
	}

	for _, test := range tests {
		result, err := parseValue(test.input)
		if test.hasError && err == nil {
			t.Errorf("expected error for input %q, got nil", test.input)
		}
		if !test.hasError && result != test.expected {
			t.Errorf("expected %f for input %q, got %f", test.expected, test.input, result)
		}
	}
}
