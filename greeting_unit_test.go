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

func TestShouldBeAbleToGreetAsHelloWithAnonymousName(t *testing.T) {
	name := ""
	want := "Hello, Anonymous!!"

	got := Greet(name)

	if want != got {
		t.Error("(", want, ") is not equal to (", got, ")")
	}
}
