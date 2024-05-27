package config

type Provider interface {
	provide() interface{}
}
