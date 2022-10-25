package libsettings

import "testing"

func TestParseSettings(t *testing.T) {
	var settings Settings

	settings.Influx.Host = "http://192.168.1.229:8086"
	settings.Influx.Token = "telegraf:apa"
	settings.Influx.Apiorg = "telegraf"
	settings.Influx.Bucket = "telegraf"

	cases := []struct {
		in   string
		want Settings
	}{
		{"../testdata/tempsensor/settings.json", settings},
	}

	for _, c := range cases {
		got, _ := ParseSettings(c.in)
		if got.Influx.Host != c.want.Influx.Host {
			t.Errorf("ParseSettings(%q) == %q, want %q", c.in, got.Influx.Host, c.want.Influx.Host)
		}

		if got.Influx.Token != c.want.Influx.Token {
			t.Errorf("ParseSettings(%q) == %q, want %q", c.in, got.Influx.Token, c.want.Influx.Token)
		}

		if got.Influx.Apiorg != c.want.Influx.Apiorg {
			t.Errorf("ParseSettings(%q) == %q, want %q", c.in, got.Influx.Apiorg, c.want.Influx.Apiorg)
		}

		if got.Influx.Bucket != c.want.Influx.Bucket {
			t.Errorf("ParseSettings(%q) == %q, want %q", c.in, got.Influx.Bucket, c.want.Influx.Bucket)
		}
	}

}

func TestParseSettingsFail(t *testing.T) {
	var settings Settings

	settings.Influx.Host = ""
	settings.Influx.Token = ""
	settings.Influx.Apiorg = ""
	settings.Influx.Bucket = ""

	cases := []struct {
		in   string
		want Settings
	}{
		{"../testdata/apa", settings},
	}

	for _, c := range cases {
		got, _ := ParseSettings(c.in)
		if got.Influx.Host != c.want.Influx.Host {
			t.Errorf("ParseSettings(%q) == %q, want %q", c.in, got.Influx.Host, c.want.Influx.Host)
		}

		if got.Influx.Token != c.want.Influx.Token {
			t.Errorf("ParseSettings(%q) == %q, want %q", c.in, got.Influx.Token, c.want.Influx.Token)
		}

		if got.Influx.Apiorg != c.want.Influx.Apiorg {
			t.Errorf("ParseSettings(%q) == %q, want %q", c.in, got.Influx.Apiorg, c.want.Influx.Apiorg)
		}

		if got.Influx.Bucket != c.want.Influx.Bucket {
			t.Errorf("ParseSettings(%q) == %q, want %q", c.in, got.Influx.Bucket, c.want.Influx.Bucket)
		}
	}

}
