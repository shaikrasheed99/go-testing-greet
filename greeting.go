package greeting

func Greet(name string) string {
	if name == "" {
		return "Hello, Anonymous!!"
	}
	return "Hello, " + name + "!!"
}
