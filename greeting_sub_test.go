package greeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetWithSubTests(main *testing.T) {
	main.Run("ShouldBeAbleToGreetAsHelloWithName", func(t *testing.T) {
		name := "Ironman"
		want := "Hello, Ironman!!"

		got := Greet(name)

		assert.Equal(t, want, got)
	})

	main.Run("ShouldBeAbleToGreetAsHelloWithAnonymousName", func(t *testing.T) {
		name := ""
		want := "Hello, Anonymous!!"

		got := Greet(name)

		assert.Equal(t, want, got)
	})
}
