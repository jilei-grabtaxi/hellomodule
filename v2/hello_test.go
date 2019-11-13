package helloV2

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world. V2"
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
