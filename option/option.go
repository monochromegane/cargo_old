package option

type Option struct {
	Image       string `short:"i" long:"image" required:"true" description:"docker image name"`
	Command     string `short:"c" long:"command" required:"true" description:"docker run command with args"`
	Concurrency int    `short:"n" default:"1" default-mask:"-" description:"Number of docker run concurrency"`
}
