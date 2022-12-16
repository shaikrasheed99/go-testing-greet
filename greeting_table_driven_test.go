package greeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	type TestCase struct {
		description string
		name        string
		want        string
	}

	testCases := []TestCase{
		{
			description: "ShouldBeAbleToGreetAsHelloWithName",
			name:        "Ironman",
			want:        "Hello, Ironman!!",
		},
		{
			description: "ShouldBeAbleToGreetAsHelloWithAnonymousName",
			name:        "",
			want:        "Hello, Anonymous!!",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := Greet(testCase.name)

			assert.Equal(t, testCase.want, got)
		})
	}
}
