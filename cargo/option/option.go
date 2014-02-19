package option

type Option struct {
	Image       string `short:"i" long:"image" required:"true" description:"docker image name"`
	Concurrency int    `short:"n" default:"1" default-mask:"-" description:"Number of docker run concurrency"`
	User        string `short:"u" default:"cargo"`
	GroupBy     string `short:"g" default:"file-size"`
	Mount       string `short:"m"`
	Command     string `short:"c"`
	Target      string `short:"t"`
	WorkDir     string `short:"w" default:"/tmp/cargo"`
}
