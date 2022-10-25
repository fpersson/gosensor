package libsensor

import "testing"

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
	cases := []struct {
		in   string
		want float64
	}{
		{"apa\nfoo=1234", 1.234},
		{"apa\nfoo=bar", 0},
	}

	for _, c := range cases {
		got, _ := parse_value(c.in)
		if got != c.want {
			t.Errorf("parse_value(%q) == %f, want %f", c.in, got, c.want)
		}
	}
}
