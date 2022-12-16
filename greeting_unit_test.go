package greeting

import "testing"

func TestShouldBeAbleToGreetAsHelloWithName(t *testing.T) {
	name := "Ironman"
	want := "Hello, Ironman!!"

	got := Greet(name)

	if want != got {
		t.Error()
	}
}
