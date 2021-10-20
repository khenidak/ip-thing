package options

type Options struct {
	Version   string
	BuildTime string
	Runtime   *Runtime
}

func MakeOptions(Version, BuildTime string) *Options {
	return &Options{
		Version:   Version,
		BuildTime: BuildTime,
		Runtime:   &Runtime{},
	}
}
