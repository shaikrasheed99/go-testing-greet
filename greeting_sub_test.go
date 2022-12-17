package greeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetWithSubTests(main *testing.T) {
	main.Run("ShouldBeAbleToGreetAsHelloWithName", func(sub *testing.T) {
		name := "Ironman"
		want := "Hello, Ironman!!"

		got := Greet(name)

		assert.Equal(sub, want, got)
	})

	main.Run("ShouldBeAbleToGreetAsHelloWithAnonymousName", func(sub *testing.T) {
		name := ""
		want := "Hello, Anonymous!!"

		got := Greet(name)

		assert.Equal(sub, want, got)
	})
}
