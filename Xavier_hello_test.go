package xavierHello
import "testing"

func HelloTest(t *testing.T) {
	want := "Xavier_say_hello_world"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q,want %q", got, want)
	}
}