package greeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeAbleToGreetAsHelloWithName(t *testing.T) {
	name := "Ironman"
	want := "Hello, Ironman!!"

	got := Greet(name)

	assert.Equal(t, want, got)
}

func TestShouldBeAbleToGreetAsHelloWithAnonymousName(t *testing.T) {
	name := ""
	want := "Hello, Anonymous!!"

	got := Greet(name)

	assert.Equal(t, want, got)
}
