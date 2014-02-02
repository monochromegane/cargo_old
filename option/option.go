package option

type Option struct {
	Image       string `short:"i" long:"image" required:"true" description:"docker image name"`
	Concurrency int    `short:"n" default:"1" default-mask:"-" description:"Number of docker run concurrency"`
}
