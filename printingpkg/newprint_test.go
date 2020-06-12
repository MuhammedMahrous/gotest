package printingpkg
import (
"testing"
"fmt"
)

func TestNewPrintable(t *testing.T){
	fmt.Println(t)
	cases := []struct {
		in, want string
	}{
		{"X", "X"},
		{"Y","Y"},
	}
	for _, c := range cases {
		got := NewPrintable(c.in)
		if got != c.want {
		//	t.Errorf("NewPrintable(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

