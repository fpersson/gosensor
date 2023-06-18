package syscmd

import "testing"

func TestParseOsRelease(t *testing.T) {
	var result = map[string]string{"NAME": "openSUSE Tumbleweed",
		"VERSION_ID": "20220602"}

	cases := []struct {
		in    string
		wants map[string]string
	}{
		{"../testdata/os-release", result},
	}

	for _, c := range cases {
		got, _ := ParseOsRelease(c.in)
		if got["NAME"] != c.wants["NAME"] {
			t.Errorf("ParseOsRelease(%q) == %q, wants %q", c.in, got["NAME"], c.wants["NAME"])
		}

		if got["VERSION_ID"] != c.wants["VERSION_ID"] {
			t.Errorf("ParseOsRelease(%q) == %q, wants %q", c.in, got["VERSION_ID"], c.wants["VERSION_ID"])
		}
	}

}
