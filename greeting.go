package greeting

var (
	helloPrefix string = "Hello, "
	helloSuffix string = "!!"
	defaultName string = "Anonymous"
)

func Greet(name string) string {
	if name == "" {
		return helloPrefix + defaultName + helloSuffix
	}
	return helloPrefix + name + helloSuffix
}
