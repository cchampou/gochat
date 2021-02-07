package customIo

import "testing"

func TestToServerAddr(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{6000, ":6000"},
		{0, ":0"},
		{-42, ":-42"},
	}
	for _, c := range cases {
		got := ToServerAddr(c.in)
		if got != c.want {
			t.Errorf("toServerAddress(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
