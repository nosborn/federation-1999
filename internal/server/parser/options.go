package parser

type Option[T any] func(T)

type DuchyOptioner interface {
	SetDuchy(name string)
}

func WithDuchy[T DuchyOptioner](name string) Option[T] {
	return func(opts T) {
		opts.SetDuchy(name)
	}
}

type PlanetOptioner interface {
	SetPlanet(name string)
}

func WithPlanet[T PlanetOptioner](name string) Option[T] {
	return func(opts T) {
		opts.SetPlanet(name)
	}
}

type MinutesOptioner interface {
	SetMinutes(minutes int32)
}

type MinutesOption func(MinutesOptioner)

func WithMinutes(minutes int32) MinutesOption {
	return func(opts MinutesOptioner) {
		opts.SetMinutes(minutes)
	}
}
